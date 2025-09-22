<template>
  <div>
    <!-- Loading State -->
    <div v-if="loading" class="flex items-center justify-center py-16">
      <div class="glass-card p-8 text-center">
        <Icon name="heroicons:arrow-path" class="h-8 w-8 text-white animate-spin mx-auto mb-4" />
        <p class="text-white">Loading libraries...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="glass-card p-8 text-center">
      <Icon name="heroicons:exclamation-triangle" class="h-16 w-16 text-red-400 mx-auto mb-4" />
      <h3 class="text-xl font-semibold text-white mb-2">Error Loading Libraries</h3>
      <p class="text-gray-400 mb-4">{{ error }}</p>
      <button @click="loadLibraries" class="glass-button">
        <Icon name="heroicons:arrow-path" class="h-4 w-4 mr-2" />
        Try Again
      </button>
    </div>

    <!-- Main Content -->
    <div v-else>
      <!-- Overview Stats -->
      <section class="mb-6">
        <div class="glass-card p-6">
          <h2 class="text-lg font-semibold text-white mb-4">Overview</h2>
          <div class="grid grid-cols-3 gap-4">
            <div class="text-center">
              <div class="text-2xl font-bold text-white">{{ totalLibraries }}</div>
              <div class="text-sm text-gray-400">Libraries</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-purple-300">{{ totalArticles }}</div>
              <div class="text-sm text-gray-400">Articles</div>
            </div>
            <div class="text-center">
              <div class="text-2xl font-bold text-green-300">{{ completedArticles }}</div>
              <div class="text-sm text-gray-400">Completed</div>
            </div>
          </div>
        </div>
      </section>

    <!-- Default Library -->
    <section class="mb-6">
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-white">Default Library</h3>
        <div class="text-sm text-gray-400">Public</div>
      </div>
      
      <LibraryCard
        :library="defaultLibrary"
        :is-default="true"
        @view="viewLibrary"
        @edit="editLibrary"
      />
    </section>

    <!-- Private Libraries -->
    <section>
      <div class="flex items-center justify-between mb-4">
        <h3 class="text-lg font-semibold text-white">Private Libraries</h3>
        <button
          @click="showCreateModal = true"
          class="glass-button !py-2 !px-4 bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30"
        >
          <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
          New Library
        </button>
      </div>

      <div v-if="privateLibraries.length === 0" class="glass-card text-center p-8">
        <Icon name="heroicons:folder-plus" class="h-16 w-16 text-gray-500 mx-auto mb-4" />
        <h4 class="text-lg font-medium text-white mb-2">No private libraries yet</h4>
        <p class="text-gray-400 mb-4">Create private libraries to organize your articles by topic, project, or course.</p>
        <button
          @click="showCreateModal = true"
          class="glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30"
        >
          <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
          Create Your First Library
        </button>
      </div>

      <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <LibraryCard
          v-for="library in privateLibraries"
          :key="library.id"
          :library="library"
          @view="viewLibrary"
          @edit="editLibrary"
          @delete="deleteLibrary"
        />
      </div>
    </section>

    <!-- Floating Action Button -->
    <button
      @click="showCreateModal = true"
      class="glass-fab fixed bottom-24 right-4 md:bottom-8"
    >
      <Icon name="heroicons:plus" class="h-6 w-6 text-white" />
    </button>

    <!-- Modals -->
    <CreateLibraryModal
      v-if="showCreateModal"
      @close="showCreateModal = false"
      @library-created="handleLibraryCreated"
    />
    
    <EditLibraryModal
      v-if="showEditModal"
      :library="editingLibrary"
      @close="showEditModal = false"
      @library-updated="handleLibraryUpdated"
    />
    
    <SearchModal
      v-if="showSearchModal"
      @close="showSearchModal = false"
      @article-selected="handleArticleSelected"
    />

      <!-- Delete Confirmation -->
      <ConfirmDialog
        v-if="showDeleteConfirm"
        title="Delete Library"
        :message="`Are you sure you want to delete '${deletingLibrary?.name}'? This action cannot be undone.`"
        confirm-text="Delete"
        cancel-text="Cancel"
        variant="danger"
        @confirm="confirmDelete"
        @cancel="showDeleteConfirm = false"
      />
    </div>
  </div>
</template>

<script setup>
import { ReadingStatus } from '#shared/types'

// Set layout and configure header
definePageMeta({
  layout: 'default'
})

// Use layout composable to configure header
const { setLayoutConfig } = useLayout()

onMounted(() => {
  setLayoutConfig({
    pageTitle: 'My Libraries',
    showAddButton: true,
    showActionButton: false
  })
})

const { showSuccess, showError } = useNotifications()

// State
const showCreateModal = ref(false)
const showEditModal = ref(false)
const showSearchModal = ref(false)
const showDeleteConfirm = ref(false)
const editingLibrary = ref(null)
const deletingLibrary = ref(null)

// Data from API
const defaultLibrary = ref(null)
const privateLibraries = ref([])
const loading = ref(true)
const error = ref(null)

// Load libraries data
const loadLibraries = async () => {
  try {
    loading.value = true
    error.value = null
    
    console.log('Loading libraries...')
    const response = await $fetch('/api/libraries')
    console.log('API response:', response)
    
    // Check if response is an error object
    if (response.error) {
      throw new Error(response.message || 'Failed to load libraries')
    }
    
    defaultLibrary.value = response.defaultLibrary
    privateLibraries.value = response.privateLibraries || []
  } catch (err) {
    console.error('Failed to load libraries:', err)
    error.value = err.message || 'Failed to load libraries'
  } finally {
    loading.value = false
  }
}

// Load data on mount
onMounted(() => {
  loadLibraries()
})

// Computed properties
const totalLibraries = computed(() => {
  if (!defaultLibrary.value) return 0
  return 1 + (privateLibraries.value?.length || 0)
})

const totalArticles = computed(() => {
  if (!defaultLibrary.value) return 0
  const defaultCount = defaultLibrary.value.articles?.length || 0
  const privateCount = privateLibraries.value?.reduce((sum, lib) => sum + (lib.articles?.length || 0), 0) || 0
  return defaultCount + privateCount
})

const completedArticles = computed(() => {
  if (!defaultLibrary.value) return 0
  const allArticles = [
    ...(defaultLibrary.value.articles || []),
    ...(privateLibraries.value?.flatMap(lib => lib.articles || []) || [])
  ]
  return allArticles.filter(article =>
    article.readingStatus === ReadingStatus.READING_STATUS_READ
  ).length
})

// Methods
const viewLibrary = (library) => {
  navigateTo(`/libraries/${library.id}`)
}

const editLibrary = (library) => {
  editingLibrary.value = library
  showEditModal.value = true
}

const deleteLibrary = (library) => {
  deletingLibrary.value = library
  showDeleteConfirm.value = true
}

const handleLibraryCreated = async (libraryData) => {
  showCreateModal.value = false
  showSuccess('Library created successfully!')
  await loadLibraries() // Reload the libraries to get fresh data
}

const handleLibraryUpdated = async (libraryData) => {
  showEditModal.value = false
  showSuccess('Library updated successfully!')
  await loadLibraries() // Reload the libraries to get fresh data
}

const confirmDelete = async () => {
  try {
    const response = await $fetch(`/api/libraries/${deletingLibrary.value.id}`, {
      method: 'DELETE'
    })
    
    if (response.success) {
      showSuccess('Library deleted successfully!')
      await loadLibraries() // Reload the libraries
    }
  } catch (err) {
    console.error('Failed to delete library:', err)
    showError(err.message || 'Failed to delete library')
  } finally {
    showDeleteConfirm.value = false
    deletingLibrary.value = null
  }
}

const handleArticleSelected = (article) => {
  showSearchModal.value = false
  // Handle article selection
}

// Handle add button click from layout
const handleAddAction = () => {
  showCreateModal.value = true
}

// Set page title
useHead({
  title: 'My Libraries - Academic Reading List Manager'
})
</script>
