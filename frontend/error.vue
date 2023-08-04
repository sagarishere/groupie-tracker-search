<template>
  <div :class="darkMode ? 'dark' : ''" class="h-full min-h-screen bg-gradient-to-r from-green-400 via-blue-500 to-purple-600 dark:bg-gradient-to-r dark:from-gray-800 dark:via-gray-700 dark:to-gray-600 dark:text-white shadow-2xl transition duration-500">
      <Navbar>
        <ModeSwitch />
      </Navbar>
      <div class="flex justify-center items-center min-h-screen">
          <div class="prose max-w-none text-gray-900 dark:text-gray-250">
              <h2 class="text-4xl font-bold">{{ error.statusCode }}</h2>
              <p class="mt-7 text-2xl">Oops, something went wrong...</p>
              <p class="mt-7 text-xl">{{ error.message }}</p>
              <p class="mt-7 text-xl">Please consider going to Home and starting over again.</p>
          </div>
      </div>
  </div>
</template>

<script setup>
let darkMode = ''
defineProps([ 'error' ])
</script>
