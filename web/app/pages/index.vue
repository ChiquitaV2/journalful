<template>
  <div>
    <!-- Welcome Section -->
    <section class="mb-8">
      <div class="glass-card text-center p-8">
        <h2 class="text-3xl font-bold text-white mb-4">
          Welcome to Your Academic Reading Hub
        </h2>
        <p class="text-gray-300 mb-6">
          Organize, track, and share your research articles with a beautiful, mobile-first interface.
        </p>
        <div class="flex flex-col gap-4 items-center">
          <button
            @click="showAddArticleModal = true"
            class="glass-button glass-button-primary"
          >
            <Icon name="heroicons:plus" class="h-5 w-5 mr-2" />
            Add Article
          </button>
          <NuxtLink to="/libraries" class="glass-button no-underline">
            <Icon name="heroicons:book-open" class="h-5 w-5 mr-2" />
            View Libraries
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- Stats Overview -->
    <section class="mb-8">
      <div class="grid grid-cols-2 gap-4">
        <div class="glass-card text-center p-4">
          <div class="text-2xl font-bold text-blue-300">{{ stats.toRead }}</div>
          <div class="text-sm text-gray-400">To Read</div>
        </div>
        <div class="glass-card text-center p-4">
          <div class="text-2xl font-bold text-yellow-400">{{ stats.reading }}</div>
          <div class="text-sm text-gray-400">Reading</div>
        </div>
        <div class="glass-card text-center p-4">
          <div class="text-2xl font-bold text-green-400">{{ stats.read }}</div>
          <div class="text-sm text-gray-400">Completed</div>
        </div>
        <div class="glass-card text-center p-4">
          <div class="text-2xl font-bold text-white">{{ stats.total }}</div>
          <div class="text-sm text-gray-400">Total</div>
        </div>
      </div>
    </section>

    <!-- Recent Activity -->
    <section class="mb-8">
      <div class="glass-card p-6">
        <h3 class="text-lg font-semibold text-white mb-4">Recent Activity</h3>
        <div v-if="recentArticles.length === 0" class="text-center py-8">
          <Icon name="heroicons:document-text" class="h-16 w-16 text-gray-500 mx-auto mb-4" />
          <p class="text-gray-400 mb-4">No articles yet</p>
          <button
            @click="showAddArticleModal = true"
            class="glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30"
          >
            <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
            Add Your First Article
          </button>
        </div>
      </div>
    </section>

    <!-- Quick Actions -->
    <section class="mb-8">
      <h3 class="text-lg font-semibold text-white mb-4">Quick Actions</h3>
      <div class="grid grid-cols-2 gap-4">
        <button
          @click="showAddArticleModal = true"
          class="glass-card p-6 text-left transition-transform hover:scale-105"
        >
          <Icon name="heroicons:document-plus" class="h-8 w-8 text-purple-400 mb-3" />
          <h4 class="font-semibold text-white mb-2">Add Article</h4>
          <p class="text-gray-400 text-sm">Import from DOI or add manually</p>
        </button>
        
        <NuxtLink
          to="/libraries/new"
          class="glass-card p-6 text-left no-underline transition-transform hover:scale-105 block"
        >
          <Icon name="heroicons:folder-plus" class="h-8 w-8 text-blue-400 mb-3" />
          <h4 class="font-semibold text-white mb-2">Create Library</h4>
          <p class="text-gray-400 text-sm">Organize articles into collections</p>
        </NuxtLink>
      </div>
    </section>

    <!-- Modals -->
    <AddArticleModal
        v-if="showAddArticleModal"
        @close="showAddArticleModal = false"
    />
  </div>
</template>

<script setup>
import { useFormatters } from '~/composables/useFormatters'

// Set layout
definePageMeta({
  layout: 'default'
})
useLayout().setLayoutConfig()
// Composables
const { formatCount } = useFormatters()
const { showWelcomeGuide, isNewUser, checkIfNewUser } = useOnboarding()

// Mock data - replace with real API calls
const stats = ref({
  total: 42,
  toRead: 15,
  reading: 8,
  read: 17
})

const recentArticles = ref([])

const showAddArticleModal = ref(false)
const doiSearch = ref('')

// Check for new user and show welcome if needed
onMounted(async () => {
  await checkIfNewUser()
  
  if (isNewUser.value) {
    // Show welcome message after a short delay
    setTimeout(() => {
      showWelcomeGuide()
    }, 1000)
  }
})

// Set page title
useHead({
  title: 'Dashboard - Academic Reading List Manager'
})

</script>
