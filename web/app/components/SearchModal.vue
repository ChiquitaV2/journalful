<template>
  <div class="fixed inset-0 z-50 glass-overlay">
    <div class="flex flex-col h-full">
      <!-- Header -->
      <div class="glass-nav p-4">
        <div class="flex items-center space-x-3">
          <button
            @click="$emit('close')"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:arrow-left" class="h-5 w-5" />
          </button>
          <div class="flex-1">
            <input
              ref="searchInput"
              v-model="searchQuery"
              type="text"
              placeholder="Search by title, author, DOI..."
              class="glass-input !mb-0"
              @input="handleSearch"
              @keyup.enter="performSearch"
            />
          </div>
          <button
            v-if="searchQuery"
            @click="clearSearch"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:x-mark" class="h-5 w-5" />
          </button>
        </div>
      </div>

      <!-- Search Content -->
      <div class="flex-1 overflow-y-auto mobile-padding">
        <!-- Search Suggestions -->
        <div v-if="!searchQuery && !isSearching" class="py-8">
          <div class="text-center mb-6">
            <Icon name="heroicons:magnifying-glass" class="h-16 w-16 text-gray-500 mx-auto mb-4" />
            <h3 class="text-lg font-medium text-white mb-2">Search Your Library</h3>
            <p class="text-gray-400">Find articles by title, author, DOI, or keywords</p>
          </div>
          
          <!-- Quick Search Suggestions -->
          <div class="space-y-2 mb-6">
            <h4 class="text-sm font-medium text-gray-300 mb-3">Quick Searches</h4>
            <button
              v-for="suggestion in searchSuggestions"
              :key="suggestion"
              @click="searchQuery = suggestion; performSearch()"
              class="glass-subtle p-3 w-full text-left rounded-lg hover:bg-white/10 transition-colors"
            >
              <div class="flex items-center space-x-3">
                <Icon name="heroicons:clock" class="h-4 w-4 text-gray-400" />
                <span class="text-white">{{ suggestion }}</span>
              </div>
            </button>
          </div>

          <!-- Recent Searches -->
          <div v-if="recentSearches.length > 0" class="space-y-2">
            <h4 class="text-sm font-medium text-gray-300 mb-3">Recent Searches</h4>
            <button
              v-for="search in recentSearches"
              :key="search"
              @click="searchQuery = search; performSearch()"
              class="glass-subtle p-3 w-full text-left rounded-lg hover:bg-white/10 transition-colors"
            >
              <div class="flex items-center justify-between">
                <div class="flex items-center space-x-3">
                  <Icon name="heroicons:clock" class="h-4 w-4 text-gray-400" />
                  <span class="text-white">{{ search }}</span>
                </div>
                <button
                  @click.stop="removeRecentSearch(search)"
                  class="text-gray-400 hover:text-white"
                >
                  <Icon name="heroicons:x-mark" class="h-4 w-4" />
                </button>
              </div>
            </button>
          </div>
        </div>

        <!-- Loading State -->
        <div v-if="isSearching" class="flex items-center justify-center py-16">
          <div class="text-center">
            <div class="w-6 h-6 border-2 border-purple-500 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
            <p class="text-gray-400">Searching your library...</p>
          </div>
        </div>

        <!-- Search Results -->
        <div v-if="searchResults.length > 0 && !isSearching" class="py-4">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-medium text-white">
              {{ searchResults.length }} {{ searchResults.length === 1 ? 'result' : 'results' }}
            </h3>
            <div class="flex items-center space-x-2">
              <select
                v-model="sortBy"
                @change="sortResults"
                class="glass-input !py-1 !px-2 text-sm"
              >
                <option value="relevance">Sort by Relevance</option>
                <option value="title">Sort by Title</option>
                <option value="year">Sort by Year</option>
                <option value="status">Sort by Status</option>
                <option value="added">Sort by Date Added</option>
              </select>
            </div>
          </div>
          
          <div class="space-y-3">
            <div
              v-for="result in sortedResults"
              :key="result.article.id"
              @click="selectArticle(result.article)"
              class="glass-card p-4 cursor-pointer hover:scale-[1.02] transition-all"
            >
              <div class="flex items-start justify-between">
                <div class="flex-1 min-w-0">
                  <h4 class="font-medium text-white text-sm mb-1 line-clamp-2">
                    <span v-html="highlightMatch(result.article.title)"></span>
                  </h4>
                  <p class="text-gray-400 text-xs mb-2">
                    <span v-html="highlightMatch(formatAuthors(result.article.authors))"></span>
                  </p>
                  <div class="flex items-center space-x-2 text-xs text-gray-500 mb-2">
                    <span v-if="result.article.publicationYear">{{ result.article.publicationYear }}</span>
                    <span v-if="result.article.journalName">• {{ result.article.journalName }}</span>
                  </div>
                  
                  <!-- Library indicators -->
                  <div v-if="result.inLibraries.length > 0" class="flex flex-wrap gap-1">
                    <span
                      v-for="libraryId in result.inLibraries"
                      :key="libraryId"
                      class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-purple-500/20 text-purple-300"
                    >
                      Library {{ libraryId }}
                    </span>
                  </div>
                </div>
                
                <div class="ml-3 flex flex-col items-end space-y-2">
                  <StatusBadge :status="result.article.readingStatus" size="xs" />
                  <Icon name="heroicons:chevron-right" class="h-4 w-4 text-gray-400" />
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- No Results -->
        <div v-if="searchAttempted && searchResults.length === 0 && !isSearching" class="text-center py-16">
          <Icon name="heroicons:document-magnifying-glass" class="h-16 w-16 text-gray-500 mx-auto mb-4" />
          <h3 class="text-lg font-medium text-white mb-2">No articles found</h3>
          <p class="text-gray-400 mb-4">Try searching with different keywords</p>
          <button
            @click="$emit('close')"
            class="glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30"
          >
            <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
            Add New Article
          </button>
        </div>
      </div>

      <!-- Quick Filters (Mobile Bottom) -->
      <div class="glass-nav p-4 border-t border-white/10">
        <div class="flex space-x-2 overflow-x-auto">
          <button
            v-for="filter in quickFilters"
            :key="filter.value"
            @click="applyQuickFilter(filter.value)"
            :class="[
              'flex-shrink-0 glass-button !py-2 !px-3 text-xs',
              activeFilter === filter.value
                ? 'bg-purple-500/30 border-purple-300/50'
                : ''
            ]"
          >
            <Icon :name="filter.icon" class="h-3 w-3 mr-1" />
            {{ filter.label }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ReadingStatus } from '~~/types'

const emit = defineEmits(['close', 'article-selected'])

const searchInput = ref(null)
const searchQuery = ref('')
const searchResults = ref([])
const searchAttempted = ref(false)
const isSearching = ref(false)
const sortBy = ref('relevance')
const activeFilter = ref('all')

// Mock data
const searchSuggestions = ref([
  'machine learning',
  'neural networks',
  'deep learning',
  'artificial intelligence'
])

const recentSearches = ref([
  'reinforcement learning',
  'computer vision',
  'natural language processing'
])

const quickFilters = ref([
  { value: 'all', label: 'All', icon: 'heroicons:squares-2x2' },
  { value: 'to-read', label: 'To Read', icon: 'heroicons:bookmark' },
  { value: 'reading', label: 'Reading', icon: 'heroicons:book-open' },
  { value: 'read', label: 'Read', icon: 'heroicons:check-circle' }
])

const handleSearch = debounce(() => {
  if (searchQuery.value.trim().length > 2) {
    performSearch()
  }
}, 300)

const performSearch = async () => {
  if (!searchQuery.value.trim()) return

  isSearching.value = true
  searchAttempted.value = true

  try {
    // Add to recent searches
    if (!recentSearches.value.includes(searchQuery.value)) {
      recentSearches.value.unshift(searchQuery.value)
      if (recentSearches.value.length > 5) {
        recentSearches.value.pop()
      }
    }

    // Mock API call
    await new Promise(resolve => setTimeout(resolve, 800))

    // Mock search results
    searchResults.value = [
      {
        article: {
          id: 1,
          title: 'Deep Learning for Academic Research: A Comprehensive Survey',
          authors: [{ name: 'Dr. Jane Smith' }, { name: 'Prof. John Doe' }],
          publicationYear: 2024,
          journalName: 'Nature Machine Intelligence',
          readingStatus: ReadingStatus.READING_STATUS_TO_READ,
          doi: '10.1038/s42256-024-00001-1'
        },
        inLibraries: [1, 2]
      },
      {
        article: {
          id: 2,
          title: 'Advances in Neural Network Architectures',
          authors: [{ name: 'Dr. Alice Johnson' }],
          publicationYear: 2023,
          journalName: 'Science',
          readingStatus: ReadingStatus.READING_STATUS_READING,
          doi: '10.1126/science.abc123'
        },
        inLibraries: [1]
      }
    ]
  } catch (error) {
    console.error('Search failed:', error)
  } finally {
    isSearching.value = false
  }
}

const sortedResults = computed(() => {
  let results = [...searchResults.value]

  switch (sortBy.value) {
    case 'title':
      results.sort((a, b) => a.article.title.localeCompare(b.article.title))
      break
    case 'year':
      results.sort((a, b) => (b.article.publicationYear || 0) - (a.article.publicationYear || 0))
      break
    case 'status':
      results.sort((a, b) => a.article.readingStatus - b.article.readingStatus)
      break
    case 'added':
      results.sort((a, b) => new Date(b.article.addedAt || 0) - new Date(a.article.addedAt || 0))
      break
  }

  return results
})

const clearSearch = () => {
  searchQuery.value = ''
  searchResults.value = []
  searchAttempted.value = false
}

const selectArticle = (article) => {
  emit('article-selected', article)
}

const removeRecentSearch = (search) => {
  const index = recentSearches.value.indexOf(search)
  if (index > -1) {
    recentSearches.value.splice(index, 1)
  }
}

const applyQuickFilter = (filterValue) => {
  activeFilter.value = filterValue
  // Apply filter logic here
  performSearch()
}

const highlightMatch = (text) => {
  if (!searchQuery.value || !text) return text
  const regex = new RegExp(`(${searchQuery.value})`, 'gi')
  return text.replace(regex, '<mark class="bg-yellow-400/30 text-yellow-200">$1</mark>')
}

const formatAuthors = (authors) => {
  if (!authors || authors.length === 0) return 'Unknown Author'
  if (authors.length === 1) return authors[0].name
  if (authors.length === 2) return `${authors[0].name} & ${authors[1].name}`
  return `${authors[0].name} et al.`
}

const sortResults = () => {
  // Trigger reactivity for sorted results
}

// Utility function for debouncing
function debounce(func, wait) {
  let timeout
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout)
      func(...args)
    }
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
  }
}

// Focus search input when modal opens
onMounted(() => {
  nextTick(() => {
    if (searchInput.value) {
      searchInput.value.focus()
    }
  })
})
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

mark {
  background-color: rgba(251, 191, 36, 0.3);
  color: rgb(251, 191, 36);
  border-radius: 2px;
  padding: 1px 2px;
}
</style>
