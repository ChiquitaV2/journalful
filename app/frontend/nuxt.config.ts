import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
  devtools: {
    enabled: true
  },
  compatibilityDate: '2025-07-11',
  future: {
    compatibilityVersion: 4
  },
  css: [
    '~/assets/styles/main.css'
  ],
  vite: {
    plugins: [
        tailwindcss()
    ]
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
