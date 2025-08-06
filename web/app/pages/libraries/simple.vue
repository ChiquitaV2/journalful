<template>
  <div class="min-h-screen pt-16 pb-20">
    <div class="mobile-padding max-w-7xl mx-auto">
      <!-- Header -->
      <div class="flex items-center justify-between mb-8">
        <h1 class="text-3xl font-bold text-white">My Libraries</h1>
        <button class="glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30">
          <Icon name="heroicons:plus" class="h-5 w-5 mr-2" />
          Create Library
        </button>
      </div>

      <!-- Overview Stats -->
      <section class="mb-8">
        <div class="grid grid-cols-3 gap-4">
          <div class="glass-card text-center p-4">
            <div class="text-2xl font-bold text-green-400">{{ stats.total }}</div>
            <div class="text-sm text-gray-400">Total Libraries</div>
          </div>
          <div class="glass-card text-center p-4">
            <div class="text-2xl font-bold text-blue-300">{{ stats.articles }}</div>
            <div class="text-sm text-gray-400">Total Articles</div>
          </div>
          <div class="glass-card text-center p-4">
            <div class="text-2xl font-bold text-yellow-400">{{ stats.shared }}</div>
            <div class="text-sm text-gray-400">Shared</div>
          </div>
        </div>
      </section>

      <!-- Default Library -->
      <section class="mb-8">
        <h2 class="text-xl font-semibold text-white mb-4">Default Library</h2>
        <div class="glass-card p-6">
          <div class="flex items-center justify-between mb-4">
            <div class="flex-1">
              <div class="flex items-center gap-3 mb-2">
                <Icon name="heroicons:book-open" class="h-6 w-6 text-purple-400" />
                <h3 class="text-lg font-semibold text-white">My Reading List</h3>
              </div>
              <p class="text-gray-400 text-sm mb-3">
                Your default collection for articles
              </p>
              <div class="flex items-center gap-4 text-sm text-gray-500">
                <span>{{ defaultLibrary.articleCount }} articles</span>
                <span>•</span>
                <span>Updated {{ defaultLibrary.lastUpdated }}</span>
              </div>
            </div>
            <button class="glass-button p-2">
              <Icon name="heroicons:arrow-right" class="h-5 w-5" />
            </button>
          </div>
        </div>
      </section>

      <!-- Private Libraries -->
      <section>
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-semibold text-white">Private Libraries</h2>
          <span class="text-gray-400 text-sm">{{ privateLibraries.length }} libraries</span>
        </div>

        <div v-if="privateLibraries.length === 0" class="glass-card text-center p-8">
          <Icon name="heroicons:folder-plus" class="h-16 w-16 text-gray-500 mx-auto mb-4" />
          <h3 class="text-lg font-medium text-white mb-2">No private libraries yet</h3>
          <p class="text-gray-400 mb-4">Create specialized collections to organize your research by topic or project.</p>
          <button class="glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30">
            <Icon name="heroicons:plus" class="h-5 w-5 mr-2" />
            Create Your First Library
          </button>
        </div>

        <div v-else class="grid grid-cols-1 gap-4">
          <div
            v-for="library in privateLibraries"
            :key="library.id"
            class="glass-card p-5 transition-transform hover:scale-105"
          >
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center gap-3">
                <Icon :name="library.icon || 'heroicons:folder'" class="h-6 w-6 text-purple-400" />
                <h3 class="text-lg font-semibold text-white">{{ library.name }}</h3>
              </div>
              <div class="flex items-center gap-2">
                <button class="glass-button p-2">
                  <Icon name="heroicons:ellipsis-horizontal" class="h-4 w-4" />
                </button>
              </div>
            </div>
            
            <p class="text-gray-400 text-sm mb-3">{{ library.description }}</p>
            
            <div class="flex items-center justify-between text-sm">
              <div class="flex items-center gap-4 text-gray-500">
                <span>{{ library.articleCount }} articles</span>
                <span>•</span>
                <span>{{ library.isPublic ? 'Public' : 'Private' }}</span>
              </div>
              <span class="text-gray-500">{{ library.lastUpdated }}</span>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
// Mock data - replace with real API calls
const stats = ref({
  total: 3,
  articles: 42,
  shared: 1
})

const defaultLibrary = ref({
  id: 'default',
  name: 'My Reading List',
  description: 'Your default collection for articles',
  articleCount: 42,
  lastUpdated: '2 hours ago',
  isPublic: false
})

const privateLibraries = ref([
  {
    id: '1',
    name: 'Machine Learning Papers',
    description: 'Latest research in ML and AI',
    articleCount: 18,
    lastUpdated: '1 day ago',
    isPublic: false,
    icon: 'heroicons:cpu-chip'
  },
  {
    id: '2',
    name: 'Climate Research',
    description: 'Environmental science and climate change studies',
    articleCount: 12,
    lastUpdated: '3 days ago',
    isPublic: true,
    icon: 'heroicons:globe-alt'
  }
])

// Set page title
useHead({
  title: 'My Libraries - Academic Reading List Manager'
})
</script>
