import packageJSON from './package.json';

export default defineNuxtConfig({
  app: {
    head: {
      title: packageJSON.name || '',
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      meta: [
        { name: 'description', content: packageJSON.description || '' }
      ],
      noscript: [
        // <noscript>JavaScript is required</noscript>
        { children: 'JavaScript is required' }
      ]
    }
  },
  routeRules : {
    '/**': { ssr: false }
  },
  devtools: { enabled: true },
  imports: {
    dirs: ['stores']
  },
  build: {
    transpile: ["@heroicons/vue", "@headlessui/vue"],
  },
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
  ],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  pinia: {
    autoImports: [
      'defineStore', // import { defineStore } from 'pinia'
    ],
  },
})
