<template>
  <div>
    <Cards :items="data" />
  </div>
</template>
  
<script setup>
  import { useBandStore } from '~/stores/band'

  const bandStore = useBandStore()
  onMounted(async () => { await bandStore.fetchBands() })
  const data = computed(() => bandStore.getBands)
</script>
  