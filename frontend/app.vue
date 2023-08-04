<template>
  <div :class="darkMode ? 'dark' : ''" class="h-full min-h-screen bg-indigo-200 dark:bg-gray-800">
    <Navbar>
      <ModeSwitch />
    </Navbar>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<script setup>
let darkMode = ''
</script>
