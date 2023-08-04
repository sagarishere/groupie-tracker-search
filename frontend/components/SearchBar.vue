<template>
  <div class="search-bar">
    <input
      v-model="searchInput"
      @input="updateResults"
      type="text"
      placeholder="Search bands..."
      class="block w-full p-4 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
    />
    <div
      class="search-results"
      v-if="searchResults && (
        searchResults?.matchedCountry?.length
      || searchResults?.selectedBands?.length 
      || searchResults?.bandMembers?.length 
      || searchResults?.bandCreationDate?.length
      || searchResults?.bandFirstAlbumDate?.length
      || searchResults?.bandLocations?.length)"
    >

    <div v-if="searchResults?.matchedCountry?.length" class="search-result">
      💁 Searching for ?<br />
      </div>
    <div class="search-result nolink-no-go" v-for="result in searchResults?.matchedCountry">
      
      <p class="mt-2">{{ result }}</p>     
    </div>

    <NuxtLink
    class="search-result bandhere"
    v-for="result in searchResults.selectedBands"
    :key="result.id"
    :to="`/band/${result.id}`"
    @click="showModal = false"
    >
    🎸{{ result.name }}<br/>
    </NuxtLink>

    <NuxtLink
        class="search-result"
        v-for="result in searchResults.bandCreationDate"
        :key="result.id"
        :to="`/band/${result.id}`"
        @click="showModal = false"
    >
    🎂⏳️{{ result.creationDate }}: {{ result.name }}<br/>
    </NuxtLink>

    <NuxtLink
        class="search-result"
        v-for="result in searchResults.bandFirstAlbumDate"
        :key="result.id"
        :to="`/band/${result.id}`"
        @click="showModal = false"
    >
    🍾⌛{{ result.firstAlbum }}: {{ result.name }}<br/>
    </NuxtLink>

    <div
        class="search-result nolink-no-go"
        v-for="result in searchResults.bandMembers"
        :key="result.id"
        :to="`/band/${result.id}`"
    >
    🕺<span class="golden">{{ result }}</span>><br/>
      <span v-for="band in artist2Band[result]" :key="band.bandid" @click="showModal = false">
          <NuxtLink :to="`/band/${band.bandid}`" class="member-of-band">
            🎸🕺 {{ band.bandname }}
          </NuxtLink><br/>
      </span>
    </div>

    <div
        class="search-result nolink-no-go-loc"
        v-for="result in searchResults.bandLocations"
        :key="result.id"
        :to="`/band/${result.id}`"
    >
    &nbsp{{ formatLocation(result) }} ✈️ :<br/>
      <span v-for="band in loc2Band[result]" :key="band.bandid" @click="showModal = false">
          <NuxtLink :to="`/band/${band.bandid}`" class="member-of-band">
            &nbsp🎸{{ band.bandname }}
          </NuxtLink><br/>
      </span>
    </div>

    </div>
  </div>
</template>

<script setup>
import { useBandStore } from '~/stores/band'

const showModal = useState('showModal')

const formatLocation = (location) => {
    let [city, country] = location.split('-').map(str => str.split('_').map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' '));
    return `${city}, ${country.length <= 3 ? country.toUpperCase() : country}`;
  }

const bandStore = useBandStore()
const allBands = bandStore.getBands

// artist2Band = {artistname1: [{ bandid: 1, bandname: 'The Beatles' }]}
const artist2Band = {}

for (let i = 0; i < allBands.length; i++) {
  for (let j = 0; j < allBands[i].members.length; j++) {
    if (!artist2Band[allBands[i].members[j]]) {
      artist2Band[allBands[i].members[j]] = []
    }
    artist2Band[allBands[i].members[j]].push({bandid: allBands[i].id, bandname: allBands[i].name})
  }
}

// loc2Band = {locationname1: [{ bandid: 1, bandname: 'The Beatles' }]}
const loc2Band = {}

for (let i = 0; i < allBands.length; i++) {
  const locationsArrayHere = Object.keys(allBands[i].concertsLocDates)
  for (let j = 0; j < locationsArrayHere.length; j++) {
    if (!loc2Band[locationsArrayHere[j]]) {
      loc2Band[locationsArrayHere[j]] = []
    }
    loc2Band[locationsArrayHere[j]].push({bandid: allBands[i].id, bandname: allBands[i].name})
  }
}

const sanitizedBands = allBands.map(({ id, name, members, concertsLocDates, creationDate, firstAlbum }) => 
({ id, name, 
  members: [...members],
  concertsLocs: Object.keys(concertsLocDates),
  creationDate: creationDate.toString(), 
  firstAlbum: firstAlbum.toString() 
}));

// go through each object in santizedBands, and members of each object to allArtistsSet
const allAritistSet = new Set(sanitizedBands.flatMap(({ members }) => members));
const allArtists = [...allAritistSet];

const allLocationsSet = new Set(sanitizedBands.flatMap(({ concertsLocs }) => concertsLocs));
const allLocations = [...allLocationsSet];

const countryHasLocs = {}

for (let i = 0; i < allLocations.length; i++) {
  const country = allLocations[i].split('-')[1].toLocaleLowerCase()
  const specificLoc = allLocations[i].toLocaleLowerCase()
  if (!countryHasLocs[country]) {
    countryHasLocs[country] = []
  }
  countryHasLocs[country].push(specificLoc)
}

const searchInput = ref('')
const searchResults = computed(() => {
  if (!searchInput.value) return []
  return updateResults()
})

const updateResults = () => {
  if (!searchInput.value) return []

  // Update the search results based on the search input

  // Step 1: search in the names
  const selectedBands = sanitizedBands.filter(band => band.name.toLocaleLowerCase().startsWith(searchInput.value.toLocaleLowerCase()))

  // Step 2: search in the members
  const bandMembers = allArtists.filter(artist => artist.toLocaleLowerCase().startsWith(searchInput.value.toLocaleLowerCase()))

  // Step 3: search in the creation date
  const bandCreationDate = sanitizedBands.filter(band => band.creationDate.toLocaleLowerCase().startsWith(searchInput.value.toLocaleLowerCase()))

  // Step 4: search in the first album
  const bandFirstAlbumDate = sanitizedBands.filter(band => band.firstAlbum.toLocaleLowerCase().startsWith(searchInput.value.toLocaleLowerCase()))

  // Step 5: search in the locations
  const bandLocations = allLocations.filter(loc => loc.toLocaleLowerCase().startsWith(searchInput.value.toLocaleLowerCase()))

  // Step 0: search if exact match for country
  const matchedCountry = countryHasLocs[searchInput.value.toLocaleLowerCase()]

  return {
    selectedBands,
    bandMembers,
    bandCreationDate,
    bandFirstAlbumDate,
    bandLocations,
    matchedCountry
  }

}

</script>

<style scoped>
.search-bar {
  justify-content: center;
  align-items: center;
}

.search-results {
  justify-content: left;
  font-family: monospace;
  font-size: 1em;
}

.golden {
  color: gold;
}

.search-result:hover {
  color: red;
}


.member-of-band {
  color: lightsalmon;
}
.member-of-band:hover {
  color: green;
}

.nolink-no-go {
  color: grey;
  border-color: grey;
  border-width: 0.5px;
  border-style: solid;
  border-radius: 5px;
}
.nolink-no-go:hover {
  color: pink;
}


.nolink-no-go-loc {
  color: grey;
  border-color: lightblue;
  border-width: 0.5px;
  border-style: solid;
  border-radius: 5px;
}
.nolink-no-go-loc:hover {
  color: black;
}
</style>