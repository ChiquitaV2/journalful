<template>

  <section class="mb-8">
    <div class="glass-card text-center p-8">
      <h2 class="text-3xl font-bold text-white mb-4">
        Welcome to Your Academic Reading Hub
      </h2>
      <p class="text-gray-300 mb-6">
        Organize, track, and share your research articles with a beautiful, mobile-first interface.
      </p>
      <div class="flex flex-col gap-4 items-center">
        <div
            v-if="!loggedIn"
        >
          <a
              href="/api/auth/login"
              class="glass-button glass-button-primary"
          >
            <Icon name="heroicons:lock" class="h-5 w-5 mr-2" />
            Login
          </a>

        </div>
        <div
          v-else
          >

          <button
              @click="clear"
              class="glass-button glass-button-primary"
          >
            Log out
          </button>
          <button
              @click="test"
              class="glass-button glass-button-primary"
          >
            Test
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
const { loggedIn, user, session, fetch, clear, openInPopup } = useUserSession()

definePageMeta({
  layout: 'simple',
  public: true
})

const test = async () => {
  const {data} = await useFetch('/api/libraries/1')
  console.log("data", data.value)
}
</script>

<style scoped>

</style>