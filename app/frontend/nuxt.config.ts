export default defineNuxtConfig({
  devtools: {
    enabled: true
  },
  css: [
    '~/assets/styles/main.css'
  ],
  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.API_BASE_URL || 'http://localhost:8080',
    },
  },
  app: {
    head: {
      title: 'Journalful',
    },
  },
  typescript: {
    typeCheck: true,
    shim: false,
  },
})
