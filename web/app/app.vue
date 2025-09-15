<template>
  <NuxtLayout id="app">
    <NuxtRouteAnnouncer />
    <NuxtPage />
    <NotificationContainer />
    <ToastContainer />
  </NuxtLayout>
</template>

<script setup>

// Set up the head for SEO and mobile optimization
useHead({
  title: 'Academic Reading List Manager',
  meta: [
    { name: 'description', content: 'Organize, track, and share your academic article reading progress with a modern glassmorphic interface.' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no' },
    { name: 'theme-color', content: '#1e1b4b' },
    { name: 'apple-mobile-web-app-capable', content: 'yes' },
    { name: 'apple-mobile-web-app-status-bar-style', content: 'black-translucent' }
  ],
  link: [
    { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
  ]
})

// Initialize global state
const isLoading = ref(false)
provide('isLoading', isLoading)
</script>

<style>
#app {
  min-height: 100vh;
}
</style>
