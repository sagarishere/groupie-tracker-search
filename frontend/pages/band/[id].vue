<template>
    <div v-if="loading" class="prose dark:prose-dark text-gray-900 dark:text-gray-50">
      Loading...
    </div>
    <div v-else 
         class="prose dark:prose-dark max-w-fit items-center mx-auto rounded-lg shadow-2xl border-2 p-6 my-4 transition duration-500 bg-[conic-gradient(at_left,_var(--tw-gradient-stops))] from-yellow-500 via-purple-500 to-blue-500 dark:bg-gradient-to-r dark:from-gray-800 dark:via-gray-700 dark:to-gray-600 dark:text-white">
         <h2 class="text-2xl font-semibold text-gray-900 dark:text-stone-400">{{ band.name }}</h2>
      <div class="flex justify-center items-center max-w-md">
        <img :src="band.image" alt="" class="w-full h-64 object-cover object-center rounded-lg shadow-md" style="mix-blend-mode: multiply;"  />
      </div>
      <h3 class="text-xl font-bold text-gray-900 dark:text-stone-400 mt-4">Members:</h3>
      <p class="text-gray-700 dark:text-gray-400">{{ band.members.join(', ') }}</p>
        <h3 class="text-xl font-bold text-gray-900 dark:text-stone-400 mt-4">Creation Date:</h3>
      <p class="text-gray-700 dark:text-gray-400">{{ band.creationDate }}</p>
        <h3 class="text-xl font-bold text-gray-900 dark:text-stone-400 mt-4">First Album:</h3>
      <p class="text-gray-700 dark:text-gray-400"> {{ band.firstAlbum }}</p>
      <hr class="my-2 border-t-2 border-gray-200 dark:border-gray-700" />
        <h3 class="text-xl font-bold text-gray-900 dark:text-stone-400 mt-4">CONCERTS:</h3>
      <div v-for="(dates, location) in band.concertsLocDates" :key="location" class="mt-2">
        <p class="text-gray-700 dark:text-neutral-400">{{ formatLocation(location) }}:</p>
        <p class="text-gray-600 dark:text-slate-400">{{ dates.join(', ') }}</p>
      </div>
    </div>
  </template>

  
<script setup>
import { useBandStore } from '~/stores/band'
import { useRoute } from 'vue-router'

const route = useRoute()

const bandId = route.params.id

const bandStore = useBandStore()
const band = ref(null);
const loading = ref(true);

onMounted(async () => {
    try { band.value = await bandStore.fetchBandById(bandId);
        if (!band.value) {
            bandStore.loading = false;
            navigateTo('/error');
        }
    }
    catch (error) {
        console.error(error);
    }
    finally {
        loading.value = false;
    } 
    loading.value = bandStore.loading;
});

const formatLocation = (location) => {
    let [city, country] = location.split('-')
        .map(str => str.split('_')
        .map(word => word.charAt(0).toUpperCase() + word.slice(1)).join(' '));
    return `${city}, ${country.length <= 3 ? country.toUpperCase() : country}`
}
</script>
