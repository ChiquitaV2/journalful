<template>
  <div>
    <!-- Loading State -->
    <div v-if="pending" class="flex items-center justify-center min-h-[50vh]">
      <div class="glass-card p-8 text-center">
        <div class="animate-pulse">
          <Icon name="heroicons:document-text" class="h-12 w-12 mx-auto mb-4 text-purple-400" />
          <p class="text-gray-300">Loading article...</p>
        </div>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="flex items-center justify-center min-h-[50vh]">
      <div class="glass-card p-8 text-center max-w-md mobile-padding">
        <Icon name="heroicons:exclamation-triangle" class="h-12 w-12 mx-auto mb-4 text-red-400" />
        <h3 class="text-xl font-semibold mb-2">Article Not Found</h3>
        <p class="text-gray-300 mb-4">The article you're looking for doesn't exist.</p>
        <NuxtLink to="/" class="glass-button">
          <Icon name="heroicons:arrow-left" class="h-4 w-4 mr-2" />
          Back to Dashboard
        </NuxtLink>
      </div>
    </div>

    <!-- Article Content -->
    <div v-else-if="article" class="mobile-padding max-w-4xl mx-auto">
      <!-- Article Header -->
      <div class="glass-card p-6 mb-6">
        <div class="flex items-start justify-between mb-4">
          <div class="flex items-center gap-2">
            <StatusBadge :status="article.readingStatus" />
          </div>
        </div>

        <h1 class="text-2xl md:text-3xl font-bold text-white mb-4 leading-tight">
          {{ article.title }}
        </h1>

        <div class="flex flex-wrap items-center gap-4 text-sm text-gray-300 mb-4">
          <div class="flex items-center gap-2">
            <Icon name="heroicons:user-group" class="h-4 w-4" />
            <span>{{ formatArticleAuthors(article.authors) }}</span>
          </div>
          <div class="flex items-center gap-2">
            <Icon name="heroicons:calendar" class="h-4 w-4" />
            <span>{{ formatArticleDate(article.publishedDate) }}</span>
          </div>
          <div class="flex items-center gap-2">
            <Icon name="heroicons:document-text" class="h-4 w-4" />
            <span>{{ article.journal || 'Journal' }}</span>
          </div>
        </div>

        <div class="flex flex-wrap gap-2 mb-4">
          <span 
            v-for="tag in article.tags" 
            :key="tag"
            class="px-3 py-1 bg-purple-500/20 border border-purple-500/30 rounded-full text-xs text-purple-300"
          >
            {{ tag }}
          </span>
        </div>

        <!-- Action Buttons -->
        <div class="flex flex-wrap gap-3">
          <button 
            v-if="article.url"
            class="glass-button-primary"
            @click="openArticle"
          >
            <Icon name="heroicons:arrow-top-right-on-square" class="h-4 w-4 mr-2" />
            Read Original
          </button>
          <button 
            v-if="article.doi"
            class="glass-button"
            @click="copyDOI"
          >
            <Icon name="heroicons:clipboard-document" class="h-4 w-4 mr-2" />
            Copy DOI
          </button>
          <button class="glass-button" @click="exportCitation">
            <Icon name="heroicons:document-arrow-down" class="h-4 w-4 mr-2" />
            Export Citation
          </button>
        </div>
      </div>

      <!-- Abstract -->
      <div v-if="article.abstract" class="glass-card p-6 mb-6">
        <h2 class="text-xl font-semibold text-white mb-4">Abstract</h2>
        <p class="text-gray-300 leading-relaxed">{{ article.abstract }}</p>
      </div>

      <!-- Progress and Notes -->
      <div class="glass-card p-6 mb-6">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-xl font-semibold text-white">Reading Progress</h2>
          <select 
            v-model="article.readingStatus" 
            @change="updateReadingStatus"
            class="glass-select"
          >
            <option value="READING_STATUS_TO_READ">To Read</option>
            <option value="READING_STATUS_READING">Reading</option>
            <option value="READING_STATUS_READ">Completed</option>
          </select>
        </div>

        <!-- Progress Bar -->
        <div class="mb-4">
          <div class="flex items-center justify-between text-sm text-gray-400 mb-2">
            <span>Progress</span>
            <span>{{ Math.round(readingProgress) }}%</span>
          </div>
          <div class="w-full bg-gray-600 rounded-full h-2">
            <div 
              class="bg-gradient-to-r from-purple-500 to-blue-500 h-2 rounded-full transition-all duration-300"
              :style="{ width: `${readingProgress}%` }"
            ></div>
          </div>
          <input 
            v-model="readingProgress" 
            type="range" 
            min="0" 
            max="100" 
            class="w-full mt-2"
            @change="updateProgress"
          />
        </div>

        <!-- Notes -->
        <div>
          <label class="block text-sm font-medium text-gray-300 mb-2">Notes</label>
          <textarea
            v-model="article.notes"
            @blur="updateNotes"
            placeholder="Add your reading notes here..."
            class="glass-input w-full h-32 resize-none"
          ></textarea>
        </div>
      </div>

      <!-- Article Details -->
      <div class="glass-card p-6 mb-6">
        <h2 class="text-xl font-semibold text-white mb-4">Details</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
          <div>
            <span class="text-gray-400">DOI:</span>
            <span class="text-white ml-2">{{ article.doi || 'N/A' }}</span>
          </div>
          <div>
            <span class="text-gray-400">Publication Year:</span>
            <span class="text-white ml-2">{{ article.publicationYear || 'N/A' }}</span>
          </div>
          <div>
            <span class="text-gray-400">Added:</span>
            <span class="text-white ml-2">{{ formatArticleDate(article.addedAt) }}</span>
          </div>
          <div>
            <span class="text-gray-400">Library:</span>
            <span class="text-white ml-2">{{ article.libraryName || 'Default' }}</span>
          </div>
        </div>
      </div>

      <!-- Related Articles (placeholder) -->
      <div class="glass-card p-6">
        <h2 class="text-xl font-semibold text-white mb-4">Related Articles</h2>
        <p class="text-gray-400 text-center py-8">
          Related articles will appear here based on content similarity and tags.
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useFormatters } from '~/composables/useFormatters'
import { useSharing } from '~/composables/useSharing'

// Set layout
definePageMeta({
  layout: 'reader'
})

// Get route params
const route = useRoute()
const articleId = route.params.id

// Composables
const { formatArticleAuthors, formatArticleDate } = useFormatters()
const { shareArticle: shareArticleAction, copyDOI: copyDOIAction } = useSharing()

// Fetch article data
const { data: article, pending, error } = await useFetch(`/api/articles/${articleId}`)

// Local state
const readingProgress = ref(article.value?.progress || 0)

// Configure reader layout
const { setReaderConfig } = useReaderLayout()

onMounted(() => {
  if (article.value) {
    setReaderConfig({
      showBookmark: true,
      isBookmarked: article.value.isFavorite || false,
      showShare: true,
      showSettings: false,
      showProgress: true,
      progress: readingProgress.value,
      progressText: `${Math.round(readingProgress.value)}% complete`
    })
  }
})

// Methods
const openArticle = () => {
  if (article.value.url) {
    window.open(article.value.url, '_blank')
  }
}

const copyDOI = async () => {
  if (article.value.doi) {
    await copyDOIAction(article.value.doi)
  }
}

const exportCitation = () => {
  // Mock citation export
  const citation = `${formatArticleAuthors(article.value.authors)}. "${article.value.title}." ${article.value.journal || 'Journal'} (${article.value.publicationYear}).`
  navigator.clipboard.writeText(citation)
  // Show toast notification
}

const updateReadingStatus = async () => {
  try {
    // Mock API call to update status
    await new Promise(resolve => setTimeout(resolve, 500))
    // Show success toast
  } catch (error) {
    console.error('Failed to update reading status:', error)
  }
}

const updateProgress = async () => {
  try {
    // Update progress indicator in layout
    const { setProgress } = useReaderLayout()
    setProgress(readingProgress.value, `${Math.round(readingProgress.value)}% complete`)
    
    // Mock API call to save progress
    await new Promise(resolve => setTimeout(resolve, 300))
  } catch (error) {
    console.error('Failed to update progress:', error)
  }
}

const updateNotes = async () => {
  try {
    // Mock API call to save notes
    await new Promise(resolve => setTimeout(resolve, 500))
  } catch (error) {
    console.error('Failed to update notes:', error)
  }
}

// Handle layout events
const handleBookmark = () => {
  article.value.isFavorite = !article.value.isFavorite
  // Update API
}

const handleShare = () => {
  shareArticleAction(article.value)
}

// Set page title
useHead({
  title: computed(() => article.value ? `${article.value.title} - Academic Reading List Manager` : 'Article - Academic Reading List Manager')
})
</script>
