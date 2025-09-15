import tailwindcss from "@tailwindcss/vite";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  modules: ['@nuxt/fonts', '@nuxt/icon', '@nuxt/image', 'nuxt-auth-utils'],
  css: [
    '~/assets/css/tailwind.css',
      '~/assets/css/glassmorphism.css'
  ],
  app: {
    head: {
      viewport: 'width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no',
      meta: [
        { name: 'theme-color', content: '#1e1b4b' },
        { name: 'apple-mobile-web-app-capable', content: 'yes' },
        { name: 'apple-mobile-web-app-status-bar-style', content: 'black-translucent' }
      ]
    }
  },
  runtimeConfig: {
    backendURL: process.env.BACKEND_URL || 'localhost:50052',
  },
  typescript: {
    typeCheck: true
  },
  vite: {
    plugins: [
      tailwindcss(),
    ],
  }
})