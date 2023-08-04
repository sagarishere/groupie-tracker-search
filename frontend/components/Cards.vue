<template>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      <Card v-for="item in items" :key="item.id" :item="item" />
    </div>
  </template>
  
  <script setup>
  const props = defineProps({
    items: {
      type: Array,
      default: () => []
    }
  })
  </script>
  