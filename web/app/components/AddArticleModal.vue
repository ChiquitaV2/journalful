<template>
  <div class="fixed inset-0 z-50 glass-overlay">
    <div class="flex items-center justify-center min-h-screen p-4">
      <div
        class="glass-elevated w-full max-w-lg rounded-2xl p-6 animate-scale-in"
        @click.stop
      >
        <!-- Header -->
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-xl font-bold text-white">Add New Article</h2>
          <button
            @click="$emit('close')"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:x-mark" class="h-5 w-5" />
          </button>
        </div>

        <!-- Search by DOI Section -->
        <div class="mb-6">
          <h3 class="text-sm font-medium text-gray-300 mb-3">Search by DOI</h3>
          <div class="flex space-x-2">
            <input
              v-model="doiSearch"
              type="text"
              placeholder="Enter DOI (e.g., 10.1000/xyz123)"
              class="glass-input flex-1"
              @keyup.enter="searchByDOI"
            />
            <button
              @click="searchByDOI"
              :disabled="isSearching || !doiSearch.trim()"
              class="glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30 disabled:opacity-50"
            >
              <Icon
                v-if="isSearching"
                name="heroicons:arrow-path"
                class="h-4 w-4 animate-spin"
              />
              <Icon v-else name="heroicons:magnifying-glass" class="h-4 w-4" />
            </button>
          </div>
          
          <!-- Search Results -->
          <div v-if="searchResult" class="mt-4 glass-subtle p-4 rounded-lg">
            <div class="flex items-start justify-between">
              <div class="flex-1">
                <h4 class="font-medium text-white mb-1">{{ searchResult.title }}</h4>
                <p class="text-sm text-gray-400 mb-2">{{ formatAuthors(searchResult.authors) }}</p>
                <div class="flex items-center space-x-2 text-xs text-gray-500">
                  <span v-if="searchResult.publicationYear">{{ searchResult.publicationYear }}</span>
                  <span v-if="searchResult.journalName">• {{ searchResult.journalName }}</span>
                </div>
              </div>
              <button
                @click="addSearchedArticle"
                class="glass-button !py-2 !px-4 bg-green-500/20 border-green-300/30 text-green-300"
              >
                Add to Library
              </button>
            </div>
          </div>
          
          <!-- No Results -->
          <div v-if="searchAttempted && !searchResult && !isSearching" class="mt-4 text-center p-4">
            <Icon name="heroicons:exclamation-triangle" class="h-8 w-8 text-amber-400 mx-auto mb-2" />
            <p class="text-sm text-gray-400">Article not found. You can add it manually below.</p>
          </div>
        </div>

        <!-- Divider -->
        <div class="flex items-center my-6">
          <div class="flex-1 h-px bg-white/20"></div>
          <span class="px-3 text-sm text-gray-400">or</span>
          <div class="flex-1 h-px bg-white/20"></div>
        </div>

        <!-- Manual Entry Form -->
        <form @submit.prevent="addManualArticle">
          <h3 class="text-sm font-medium text-gray-300 mb-4">Add Manually</h3>
          
          <div class="space-y-4">
            <!-- Title -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Title *
              </label>
              <input
                v-model="manualArticle.title"
                type="text"
                required
                class="glass-input"
                placeholder="Enter article title"
              />
            </div>

            <!-- DOI -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                DOI *
              </label>
              <input
                v-model="manualArticle.doi"
                type="text"
                required
                class="glass-input"
                placeholder="10.1000/xyz123"
              />
            </div>

            <!-- Authors -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Authors
              </label>
              <input
                v-model="manualArticle.authors"
                type="text"
                class="glass-input"
                placeholder="John Doe, Jane Smith"
              />
            </div>

            <!-- Publication Year -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Publication Year
              </label>
              <input
                v-model="manualArticle.publicationYear"
                type="number"
                min="1900"
                :max="currentYear"
                class="glass-input"
                placeholder="2024"
              />
            </div>

            <!-- Journal Name -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Journal Name
              </label>
              <input
                v-model="manualArticle.journalName"
                type="text"
                class="glass-input"
                placeholder="Nature, Science, etc."
              />
            </div>

            <!-- Abstract -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Abstract (Optional)
              </label>
              <textarea
                v-model="manualArticle.abstract"
                rows="3"
                class="glass-input resize-none"
                placeholder="Article abstract..."
              ></textarea>
            </div>

            <!-- Reading Status -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Initial Reading Status
              </label>
              <select v-model="manualArticle.readingStatus" class="glass-input">
                <option :value="ReadingStatus.READING_STATUS_TO_READ">To Read</option>
                <option :value="ReadingStatus.READING_STATUS_READING">Currently Reading</option>
                <option :value="ReadingStatus.READING_STATUS_READ">Already Read</option>
              </select>
            </div>

            <!-- Notes -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Notes (Optional)
              </label>
              <textarea
                v-model="manualArticle.notes"
                rows="2"
                class="glass-input resize-none"
                placeholder="Your personal notes about this article..."
              ></textarea>
            </div>
          </div>

          <!-- Form Actions -->
          <div class="flex space-x-3 mt-6">
            <button
              type="button"
              @click="$emit('close')"
              class="flex-1 glass-button"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="isAdding"
              class="flex-1 glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30 disabled:opacity-50"
            >
              <Icon
                v-if="isAdding"
                name="heroicons:arrow-path"
                class="h-4 w-4 mr-2 animate-spin"
              />
              Add Article
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ReadingStatus } from '~~/types'

const emit = defineEmits(['close', 'article-added'])
const { showSuccess, showError } = useNotifications()

// DOI Search
const doiSearch = ref('')
const searchResult = ref(null)
const searchAttempted = ref(false)
const isSearching = ref(false)

// Manual Entry
const currentYear = new Date().getFullYear()
const isAdding = ref(false)
const manualArticle = ref({
  title: '',
  doi: '',
  authors: '',
  publicationYear: null,
  journalName: '',
  abstract: '',
  readingStatus: ReadingStatus.READING_STATUS_TO_READ,
  notes: ''
})

const searchByDOI = async () => {
  if (!doiSearch.value.trim()) return
  
  isSearching.value = true
  searchAttempted.value = true
  searchResult.value = null
  
  try {
    // Mock API call - replace with actual API
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Mock successful search result
    if (doiSearch.value.includes('10.')) {
      searchResult.value = {
        id: Date.now(),
        title: 'Advances in Machine Learning for Academic Research',
        doi: doiSearch.value,
        authors: [
          { id: 1, name: 'Dr. Jane Smith', profileId: 1 },
          { id: 2, name: 'Prof. John Doe', profileId: 2 }
        ],
        publicationYear: 2024,
        journalName: 'Nature Machine Intelligence',
        abstract: 'This paper presents novel approaches to machine learning applications in academic research...'
      }
    }
  } catch (error) {
    showError('Failed to search for article. Please try again.')
  } finally {
    isSearching.value = false
  }
}

const addSearchedArticle = async () => {
  try {
    // Mock API call to add article
    await new Promise(resolve => setTimeout(resolve, 500))
    
    emit('article-added', searchResult.value)
    showSuccess('Article added successfully!')
  } catch (error) {
    showError('Failed to add article. Please try again.')
  }
}

const addManualArticle = async () => {
  isAdding.value = true
  
  try {
    // Validate required fields
    if (!manualArticle.value.title.trim() || !manualArticle.value.doi.trim()) {
      showError('Title and DOI are required fields.')
      return
    }
    
    // Mock API call to add article
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Transform authors string to array
    const authorsArray = manualArticle.value.authors
      ? manualArticle.value.authors.split(',').map((name, index) => ({
          id: Date.now() + index,
          name: name.trim(),
          profileId: 0
        }))
      : []
    
    const newArticle = {
      id: Date.now(),
      ...manualArticle.value,
      authors: authorsArray,
      createdAt: new Date(),
      updatedAt: new Date()
    }
    
    emit('article-added', newArticle)
    showSuccess('Article added successfully!')
  } catch (error) {
    showError('Failed to add article. Please try again.')
  } finally {
    isAdding.value = false
  }
}

const formatAuthors = (authors) => {
  if (!authors || authors.length === 0) return 'Unknown Author'
  if (authors.length === 1) return authors[0].name
  if (authors.length === 2) return `${authors[0].name} & ${authors[1].name}`
  return `${authors[0].name} et al.`
}

// Reset form when modal opens
onMounted(() => {
  doiSearch.value = ''
  searchResult.value = null
  searchAttempted.value = false
  manualArticle.value = {
    title: '',
    doi: '',
    authors: '',
    publicationYear: null,
    journalName: '',
    abstract: '',
    readingStatus: ReadingStatus.READING_STATUS_TO_READ,
    notes: ''
  }
})
</script>
