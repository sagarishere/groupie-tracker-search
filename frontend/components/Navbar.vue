<template>
  <nav
      class="p-4 flex flex-col md:flex-row md:justify-around items-center bg-indigo-100 dark:bg-gray-800 border-b border-gray-300 dark:border-gray-700 w-full sticky top-0 z-50"
      style="margin-top: 0px; margin-bottom: 0px;"
    >
      <NuxtLink to="/" class="text-xl font-bold text-black dark:text-white md:block md:mb-4 lg:m-8" style="margin-top: 0px; margin-bottom: 0px;">Home</NuxtLink>
  <!-- Other code remains the same -->
    <button class="text-2xl font-semibold text-black dark:text-white mx-auto sm:mx-0 md:mb-4 lg:m-8" 
      style="margin-top: 0px; margin-bottom: 0px;" @click="showModal = true">
      Search
    </button>
    <slot class="lg:m-8" style="margin-top: 0px; margin-bottom: 0px;"/> <!-- This slot is for child components to be inserted. -->

    <div v-if="showModal" class="fixed z-10 inset-0 overflow-y-auto" aria-labelledby="modal-title" role="dialog" aria-modal="true">
      <div class="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- The following div is the backdrop of the modal -->
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>
        <div class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full dark:bg-gray-800">
          <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4 dark:bg-gray-800 dark:text-white">
            <SearchBar class="text-2xl font-semibold text-black dark:text-white mx-auto sm:mx-0 md:mb-4 lg:m-8" style="margin-top: 0px; margin-bottom: 0px;"/>
          </div>
          <div class="bg-gray-50 px-4 py-3 sm:px-6 sm:flex sm:flex-row-reverse dark:bg-gray-800">
            <button type="button" class="mt-3 w-full inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 bg-indigo-600 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm dark:bg-gray-700 dark:text-white" 
            @click="showModal = false">
              Close
            </button>
          </div>
        </div>
      </div>
    </div>

  </nav>


</template>
  
<script setup>
const showModal = useState('showModal', () => false)
  // let showModal = ref(false)
  // const closeModal = () => showModal.value = false
</script>
