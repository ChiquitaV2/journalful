<template>
  <div>
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
</template>

<script setup>
import { ReadingStatus } from '~~/types'

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

// Mock data - replace with real API calls
const defaultLibrary = ref({
  id: 1,
  name: 'My Reading List',
  userId: 1,
  articles: [
    {
      id: 1,
      articleTitle: 'Deep Learning for Academic Research',
      doi: '10.1038/s42256-024-00001-1',
      readingStatus: ReadingStatus.READING_STATUS_READING,
      addedAt: new Date('2024-01-15'),
      publicationYear: 2024
    }
  ],
  createdAt: new Date('2024-01-01'),
  updatedAt: new Date('2024-01-15')
})

const privateLibraries = ref([
  {
    id: 2,
    name: 'Machine Learning',
    userId: 1,
    articles: [
      {
        id: 2,
        articleTitle: 'Neural Networks in Computer Vision',
        doi: '10.1109/CVPR.2024.12345',
        readingStatus: ReadingStatus.READING_STATUS_TO_READ,
        addedAt: new Date('2024-01-20'),
        publicationYear: 2024
      }
    ],
    createdAt: new Date('2024-01-10'),
    updatedAt: new Date('2024-01-20')
  },
  {
    id: 3,
    name: 'Natural Language Processing',
    userId: 1,
    articles: [],
    createdAt: new Date('2024-01-25'),
    updatedAt: new Date('2024-01-25')
  }
])

// Computed properties
const totalLibraries = computed(() => 1 + privateLibraries.value.length)

const totalArticles = computed(() => {
  const defaultCount = defaultLibrary.value.articles.length
  const privateCount = privateLibraries.value.reduce((sum, lib) => sum + lib.articles.length, 0)
  return defaultCount + privateCount
})

const completedArticles = computed(() => {
  const allArticles = [
    ...defaultLibrary.value.articles,
    ...privateLibraries.value.flatMap(lib => lib.articles)
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

const confirmDelete = async () => {
  try {
    // Mock API call
    await new Promise(resolve => setTimeout(resolve, 500))
    
    const index = privateLibraries.value.findIndex(lib => lib.id === deletingLibrary.value.id)
    if (index > -1) {
      privateLibraries.value.splice(index, 1)
    }
    
    showSuccess(`Library "${deletingLibrary.value.name}" deleted successfully`)
  } catch (error) {
    showError('Failed to delete library. Please try again.')
  } finally {
    showDeleteConfirm.value = false
    deletingLibrary.value = null
  }
}

const handleLibraryCreated = (newLibrary) => {
  privateLibraries.value.push(newLibrary)
  showCreateModal.value = false
  showSuccess('Library created successfully!')
}

const handleLibraryUpdated = (updatedLibrary) => {
  if (updatedLibrary.id === defaultLibrary.value.id) {
    defaultLibrary.value = { ...defaultLibrary.value, ...updatedLibrary }
  } else {
    const index = privateLibraries.value.findIndex(lib => lib.id === updatedLibrary.id)
    if (index > -1) {
      privateLibraries.value[index] = { ...privateLibraries.value[index], ...updatedLibrary }
    }
  }
  showEditModal.value = false
  showSuccess('Library updated successfully!')
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
