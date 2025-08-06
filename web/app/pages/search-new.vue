<template>
  <div class="max-w-4xl mx-auto pb-20">
    <!-- Search Header -->
    <div class="glass-card p-6 mb-6">
      <div class="flex items-center gap-4 mb-4">
        <h1 class="text-2xl font-bold text-white flex-1">Search Results</h1>
      </div>

      <!-- Search Input -->
      <div class="relative mb-4">
        <Icon name="heroicons:magnifying-glass" class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search articles, authors, or topics..."
          class="w-full pl-12 text-lg glass-input p-4"
          @keyup.enter="performSearch"
          autofocus
        />
        <button 
          v-if="searchQuery"
          @click="clearSearch"
          class="absolute right-3 top-1/2 transform -translate-y-1/2 p-1 rounded-full hover:bg-white/10"
        >
          <Icon name="heroicons:x-mark" class="h-4 w-4 text-gray-400" />
        </button>
      </div>

      <!-- Search Filters -->
      <div class="flex flex-wrap gap-3">
        <select v-model="searchType" class="glass-select">
          <option value="all">All Content</option>
          <option value="title">Titles Only</option>
          <option value="author">Authors Only</option>
          <option value="abstract">Abstracts</option>
          <option value="tags">Tags</option>
        </select>

        <select v-model="filterLibrary" class="glass-select">
          <option value="">All Libraries</option>
          <option v-for="library in libraries" :key="library.id" :value="library.id">
            {{ library.name }}
          </option>
        </select>

        <select v-model="filterStatus" class="glass-select">
          <option value="">Any Status</option>
          <option value="unread">Unread</option>
          <option value="in-progress">In Progress</option>
          <option value="completed">Completed</option>
        </select>

        <select v-model="sortBy" class="glass-select">
          <option value="relevance">Relevance</option>
          <option value="dateAdded">Date Added</option>
          <option value="publishedDate">Published Date</option>
          <option value="title">Title</option>
        </select>
      </div>
    </div>

    <!-- Search Results -->
    <div v-if="hasSearched">
      <!-- Results Summary -->
      <div class="glass-card p-4 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <span class="text-gray-300">
              {{ filteredResults.length }} result{{ filteredResults.length !== 1 ? 's' : '' }}
              <span v-if="currentQuery">for "<span class="text-white font-medium">{{ currentQuery }}</span>"</span>
            </span>
            <div v-if="isLoading" class="flex items-center gap-2 text-gray-400">
              <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-purple-400"></div>
              <span>Searching...</span>
            </div>
          </div>
          <button
            v-if="currentQuery"
            @click="clearSearch"
            class="glass-button !py-1 !px-3 text-sm"
          >
            Clear
          </button>
        </div>
      </div>

      <!-- Search Results List -->
      <div v-if="filteredResults.length > 0" class="space-y-4">
        <ArticleCard
          v-for="article in filteredResults"
          :key="article.id"
          :article="article"
          :show-library="true"
          @click="viewArticle(article)"
          @status-change="updateArticleStatus"
        />

        <!-- Load More Button (if paginated) -->
        <div v-if="hasMoreResults" class="text-center pt-6">
          <button
            @click="loadMoreResults"
            :disabled="isLoadingMore"
            class="glass-button"
          >
            <span v-if="isLoadingMore" class="flex items-center gap-2">
              <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-purple-400"></div>
              Loading more...
            </span>
            <span v-else>Load More Results</span>
          </button>
        </div>
      </div>

      <!-- No Results -->
      <div v-else-if="!isLoading" class="glass-card text-center p-12">
        <Icon name="heroicons:magnifying-glass-minus" class="h-16 w-16 text-gray-500 mx-auto mb-4" />
        <h3 class="text-xl font-medium text-white mb-2">No articles found</h3>
        <p class="text-gray-400 mb-6">
          We couldn't find any articles matching your search criteria.
        </p>
        <div class="flex flex-col sm:flex-row gap-3 justify-center">
          <button
            @click="clearSearch"
            class="glass-button"
          >
            Clear Search
          </button>
          <button
            @click="showAddModal = true"
            class="glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30"
          >
            <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
            Add Article
          </button>
        </div>
      </div>
    </div>

    <!-- Search Suggestions (before search) -->
    <div v-else class="space-y-6">
      <!-- Recent Searches -->
      <div v-if="recentSearches.length > 0" class="glass-card p-6">
        <h3 class="text-lg font-semibold text-white mb-4">Recent Searches</h3>
        <div class="flex flex-wrap gap-2">
          <button
            v-for="search in recentSearches"
            :key="search"
            @click="searchQuery = search; performSearch()"
            class="glass-button !py-2 !px-3 text-sm"
          >
            {{ search }}
          </button>
        </div>
      </div>

      <!-- Popular Topics -->
      <div class="glass-card p-6">
        <h3 class="text-lg font-semibold text-white mb-4">Popular Topics</h3>
        <div class="grid grid-cols-2 sm:grid-cols-3 gap-3">
          <button
            v-for="topic in popularTopics"
            :key="topic"
            @click="searchQuery = topic; performSearch()"
            class="glass-card p-3 text-left text-sm hover:bg-white/10 transition-colors"
          >
            {{ topic }}
          </button>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="glass-card p-6">
        <h3 class="text-lg font-semibold text-white mb-4">Quick Actions</h3>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <button
            @click="showAddModal = true"
            class="glass-card p-4 text-left transition-transform hover:scale-105"
          >
            <Icon name="heroicons:document-plus" class="h-6 w-6 text-purple-400 mb-2" />
            <h4 class="font-medium text-white mb-1">Add Article</h4>
            <p class="text-gray-400 text-sm">Import from DOI or add manually</p>
          </button>
          <NuxtLink
            to="/libraries"
            class="glass-card p-4 text-left no-underline transition-transform hover:scale-105"
          >
            <Icon name="heroicons:book-open" class="h-6 w-6 text-blue-400 mb-2" />
            <h4 class="font-medium text-white mb-1">Browse Libraries</h4>
            <p class="text-gray-400 text-sm">Explore your organized collections</p>
          </NuxtLink>
        </div>
      </div>
    </div>

    <!-- Add Article Modal -->
    <AddArticleModal
      v-if="showAddModal"
      @close="showAddModal = false"
      @article-added="handleArticleAdded"
    />
  </div>
</template>

<script setup>
import { useSearch } from '~/composables/useSearch'

// Set layout
definePageMeta({
  layout: 'simple'
})

// Composables
const { searchArticles, clearSearch: clearSearchData } = useSearch()

// State
const searchQuery = ref('')
const currentQuery = ref('')
const searchType = ref('all')
const filterLibrary = ref('')
const filterStatus = ref('')
const sortBy = ref('relevance')
const isLoading = ref(false)
const isLoadingMore = ref(false)
const hasSearched = ref(false)
const hasMoreResults = ref(false)
const showAddModal = ref(false)

// Mock data
const libraries = ref([
  { id: 1, name: 'My Reading List' },
  { id: 2, name: 'Machine Learning' },
  { id: 3, name: 'Natural Language Processing' }
])

const recentSearches = ref([
  'machine learning',
  'neural networks',
  'deep learning',
  'computer vision'
])

const popularTopics = ref([
  'Machine Learning',
  'Neural Networks',
  'Computer Vision',
  'Natural Language Processing',
  'Deep Learning',
  'Artificial Intelligence'
])

const searchResults = ref([])
const filteredResults = computed(() => {
  // Apply filters to search results
  let results = [...searchResults.value]

  if (filterLibrary.value) {
    results = results.filter(article => article.libraryId === parseInt(filterLibrary.value))
  }

  if (filterStatus.value) {
    const statusMap = {
      'unread': 'READING_STATUS_TO_READ',
      'in-progress': 'READING_STATUS_READING',
      'completed': 'READING_STATUS_READ'
    }
    results = results.filter(article => article.readingStatus === statusMap[filterStatus.value])
  }

  // Apply sorting
  if (sortBy.value === 'dateAdded') {
    results.sort((a, b) => new Date(b.addedAt) - new Date(a.addedAt))
  } else if (sortBy.value === 'publishedDate') {
    results.sort((a, b) => b.publicationYear - a.publicationYear)
  } else if (sortBy.value === 'title') {
    results.sort((a, b) => a.articleTitle.localeCompare(b.articleTitle))
  }

  return results
})

// Methods
const performSearch = async () => {
  if (!searchQuery.value.trim()) return

  isLoading.value = true
  hasSearched.value = true
  currentQuery.value = searchQuery.value

  try {
    // Add to recent searches
    if (!recentSearches.value.includes(searchQuery.value)) {
      recentSearches.value.unshift(searchQuery.value)
      recentSearches.value = recentSearches.value.slice(0, 5)
    }

    // Mock search API call
    const results = await searchArticles(searchQuery.value, {
      type: searchType.value,
      library: filterLibrary.value,
      status: filterStatus.value
    })

    searchResults.value = results
    hasMoreResults.value = results.length >= 20 // Mock pagination
  } catch (error) {
    console.error('Search failed:', error)
    searchResults.value = []
  } finally {
    isLoading.value = false
  }
}

const clearSearch = () => {
  searchQuery.value = ''
  currentQuery.value = ''
  searchResults.value = []
  hasSearched.value = false
  clearSearchData()
}

const loadMoreResults = async () => {
  isLoadingMore.value = true
  try {
    // Mock API call for more results
    await new Promise(resolve => setTimeout(resolve, 1000))
    // Append more results...
    hasMoreResults.value = false // Mock: no more results
  } finally {
    isLoadingMore.value = false
  }
}

const viewArticle = (article) => {
  navigateTo(`/articles/${article.id}`)
}

const updateArticleStatus = (articleId, newStatus) => {
  const article = searchResults.value.find(a => a.id === articleId)
  if (article) {
    article.readingStatus = newStatus
  }
}

const handleArticleAdded = (newArticle) => {
  showAddModal.value = false
  // Optionally refresh search results or add to current results
}

// Watch for filter changes to trigger new search
watch([searchType, filterLibrary, filterStatus], () => {
  if (hasSearched.value && currentQuery.value) {
    performSearch()
  }
})

// Set page title
useHead({
  title: 'Search - Academic Reading List Manager'
})
</script>
