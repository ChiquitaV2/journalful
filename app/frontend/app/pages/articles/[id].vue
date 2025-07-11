<template>
  <div class="container py-8">
    <div v-if="loading" class="flex justify-center py-12">
      <LoadingSpinner :size="40" />
    </div>
    
    <div v-else-if="error" class="text-center py-12">
      <Icon name="heroicons:exclamation-triangle" class="h-16 w-16 text-red-500 mx-auto mb-4" />
      <h2 class="text-2xl font-bold text-gray-900 mb-2">Article Not Found</h2>
      <p class="text-gray-600 mb-4">{{ error }}</p>
      <Button @click="navigateTo('/articles')">
        <Icon name="heroicons:arrow-left" class="h-4 w-4 mr-2" />
        Back to Articles
      </Button>
    </div>
    
    <div v-else-if="article" class="max-w-4xl mx-auto">
      <!-- Article Header -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-8">
        <div class="flex items-start justify-between mb-6">
          <div class="flex-1">
            <h1 class="text-3xl font-bold text-gray-900 mb-4">
              {{ article.title }}
            </h1>
            
            <!-- Authors -->
            <div class="flex flex-wrap items-center gap-2 mb-4">
              <span
                v-for="(author, index) in article.authors"
                :key="author.id"
                class="text-lg text-gray-700"
              >
                <NuxtLink 
                  :to="`/profile/${author.profileId}`" 
                  class="hover:text-primary transition-colors font-medium"
                >
                  {{ author.name }}
                </NuxtLink>
                <span v-if="index < article.authors.length - 1" class="text-gray-500">, </span>
              </span>
            </div>
            
            <!-- Publication Info -->
            <div class="flex flex-wrap items-center gap-6 text-gray-600 mb-6">
              <div v-if="article.journalName" class="flex items-center">
                <Icon name="heroicons:document-text" class="h-5 w-5 mr-2" />
                <span class="font-medium">{{ article.journalName }}</span>
              </div>
              <div v-if="article.publicationYear" class="flex items-center">
                <Icon name="heroicons:calendar" class="h-5 w-5 mr-2" />
                <span>{{ article.publicationYear }}</span>
              </div>
              <div v-if="article.doi" class="flex items-center">
                <Icon name="heroicons:link" class="h-5 w-5 mr-2" />
                <span class="text-sm">DOI: {{ article.doi }}</span>
              </div>
            </div>
          </div>
          
          <!-- Action Buttons -->
          <div class="flex flex-col gap-2 ml-8">
            <Button
              @click="toggleLibrary"
              :loading="isAddingToLibrary"
              :variant="isInLibrary ? 'primary' : 'outline'"
            >
              <Icon 
                :name="isInLibrary ? 'heroicons:bookmark-solid' : 'heroicons:bookmark'" 
                class="h-4 w-4 mr-2"
              />
              {{ isInLibrary ? 'In Library' : 'Add to Library' }}
            </Button>
            
            <Button variant="outline" @click="showShareModal = true">
              <Icon name="heroicons:share" class="h-4 w-4 mr-2" />
              Share
            </Button>
            
            <Button variant="outline" @click="exportCitation">
              <Icon name="heroicons:document-arrow-down" class="h-4 w-4 mr-2" />
              Export Citation
            </Button>
          </div>
        </div>
        
        <!-- Reading Status -->
        <div v-if="readingStatus" class="mb-6">
          <div class="flex items-center gap-4">
            <span class="text-sm font-medium text-gray-700">Reading Status:</span>
            <div class="flex gap-2">
              <Button
                v-for="status in readingStatusOptions"
                :key="status.value"
                size="sm"
                :variant="readingStatus === status.value ? 'primary' : 'outline'"
                @click="updateReadingStatus(status.value)"
              >
                <Icon :name="status.icon" class="h-4 w-4 mr-1" />
                {{ status.label }}
              </Button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Article Content -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">Abstract</h2>
        
        <div v-if="article.abstract" class="prose prose-lg max-w-none">
          <p class="text-gray-700 leading-relaxed">
            {{ article.abstract }}
          </p>
        </div>
        
        <div v-else class="text-gray-500 italic">
          No abstract available for this article.
        </div>
        
        <!-- Metadata -->
        <div class="mt-8 pt-8 border-t border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">Article Details</h3>
          
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <dt class="text-sm font-medium text-gray-500">Publication Year</dt>
              <dd class="text-sm text-gray-900">{{ article.publicationYear || 'Unknown' }}</dd>
            </div>
            
            <div>
              <dt class="text-sm font-medium text-gray-500">Journal</dt>
              <dd class="text-sm text-gray-900">{{ article.journalName || 'Unknown' }}</dd>
            </div>
            
            <div>
              <dt class="text-sm font-medium text-gray-500">DOI</dt>
              <dd class="text-sm text-gray-900">
                <a v-if="article.doi" :href="`https://doi.org/${article.doi}`" target="_blank" class="text-primary hover:underline">
                  {{ article.doi }}
                  <Icon name="heroicons:arrow-top-right-on-square" class="h-4 w-4 inline ml-1" />
                </a>
                <span v-else>Not available</span>
              </dd>
            </div>
            
            <div>
              <dt class="text-sm font-medium text-gray-500">Added to Database</dt>
              <dd class="text-sm text-gray-900">{{ formatDate(article.createdAt) }}</dd>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Related Articles -->
      <div class="mt-8">
        <h2 class="text-2xl font-bold text-gray-900 mb-6">Related Articles</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- TODO: Add related articles -->
          <div class="text-center py-12 text-gray-500">
            <Icon name="heroicons:document-text" class="h-12 w-12 mx-auto mb-4 opacity-50" />
            <p>Related articles will appear here</p>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Share Modal -->
    <Modal v-model:show="showShareModal" title="Share Article">
      <div class="space-y-4">
        <div>
          <label class="form-label">Share Link</label>
          <div class="flex gap-2">
            <Input
              :model-value="shareUrl"
              readonly
              class="flex-1"
            />
            <Button @click="copyShareUrl">
              <Icon name="heroicons:clipboard" class="h-4 w-4" />
            </Button>
          </div>
        </div>
        
        <div>
          <label class="form-label">Citation</label>
          <textarea
            :value="citationText"
            readonly
            class="form-input h-24 resize-none"
          />
        </div>
      </div>
      
      <template #footer>
        <Button variant="outline" @click="showShareModal = false">
          Close
        </Button>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import type { Article } from '#shared/grpc/articles/v1/article'
import { ReadingStatus } from '#shared/grpc/library/v1/library'

const route = useRoute()
const articleId = Number(route.params.id)

const article = ref<Article | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)
const isAddingToLibrary = ref(false)
const isInLibrary = ref(false)
const readingStatus = ref<ReadingStatus | null>(null)
const showShareModal = ref(false)

const readingStatusOptions = [
  {
    value: ReadingStatus.READING_STATUS_TO_READ,
    label: 'To Read',
    icon: 'heroicons:clock'
  },
  {
    value: ReadingStatus.READING_STATUS_READING,
    label: 'Reading',
    icon: 'heroicons:eye'
  },
  {
    value: ReadingStatus.READING_STATUS_READ,
    label: 'Read',
    icon: 'heroicons:check-circle'
  },
  {
    value: ReadingStatus.READING_STATUS_ABANDONED,
    label: 'Abandoned',
    icon: 'heroicons:x-circle'
  }
]

const shareUrl = computed(() => {
  if (process.client) {
    return window.location.href
  }
  return ''
})

const citationText = computed(() => {
  if (!article.value) return ''
  
  const authors = article.value.authors.map(a => a.name).join(', ')
  const year = article.value.publicationYear || 'n.d.'
  const title = article.value.title
  const journal = article.value.journalName || 'Unknown Journal'
  
  return `${authors} (${year}). ${title}. ${journal}.`
})

const fetchArticle = async () => {
  loading.value = true
  error.value = null
  
  try {
    // TODO: Replace with actual API call
    await new Promise(resolve => setTimeout(resolve, 500))
    
    // Mock article data
    const mockArticle: Article = {
      id: articleId,
      doi: '10.1000/182',
      title: 'Machine Learning in Healthcare: A Comprehensive Review',
      authors: [
        { 
          id: 1, 
          name: 'Dr. Sarah Johnson', 
          profileId: 1, 
          createdAt: new Date('2023-01-01'), 
          updatedAt: new Date('2023-01-01') 
        },
        { 
          id: 2, 
          name: 'Dr. Michael Chen', 
          profileId: 2, 
          createdAt: new Date('2023-01-01'), 
          updatedAt: new Date('2023-01-01') 
        }
      ],
      abstract: 'This paper presents a comprehensive review of machine learning applications in healthcare, covering recent advances in medical imaging, drug discovery, and patient diagnosis. We examine the current state of the art and discuss future directions for the field. The integration of artificial intelligence in healthcare has shown promising results across various domains, from diagnostic imaging to personalized treatment plans. Our analysis covers over 200 research papers published in the last five years, highlighting key trends and breakthrough technologies.',
      publicationYear: 2023,
      journalName: 'Journal of Medical Informatics',
      createdAt: new Date('2023-01-15'),
      updatedAt: new Date('2023-01-15')
    }
    
    article.value = mockArticle
    
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Article not found'
  } finally {
    loading.value = false
  }
}

const toggleLibrary = async () => {
  isAddingToLibrary.value = true
  try {
    // TODO: Implement actual library toggle
    await new Promise(resolve => setTimeout(resolve, 500))
    isInLibrary.value = !isInLibrary.value
    
    if (isInLibrary.value && !readingStatus.value) {
      readingStatus.value = ReadingStatus.READING_STATUS_TO_READ
    }
  } catch (err) {
    console.error('Error toggling library:', err)
  } finally {
    isAddingToLibrary.value = false
  }
}

const updateReadingStatus = async (status: ReadingStatus) => {
  try {
    // TODO: Implement actual status update
    await new Promise(resolve => setTimeout(resolve, 200))
    readingStatus.value = status
  } catch (err) {
    console.error('Error updating reading status:', err)
  }
}

const exportCitation = () => {
  // TODO: Implement citation export
  console.log('Exporting citation:', citationText.value)
}

const copyShareUrl = async () => {
  if (process.client) {
    await navigator.clipboard.writeText(shareUrl.value)
    // TODO: Show success toast
  }
}

const formatDate = (date: Date | undefined) => {
  if (!date) return 'Unknown'
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  }).format(new Date(date))
}

// SEO
useSeoMeta({
  title: computed(() => article.value ? `${article.value.title} - Journalful` : 'Article - Journalful'),
  description: computed(() => article.value?.abstract || 'Read this academic article on Journalful'),
})

// Initialize
onMounted(() => {
  fetchArticle()
})
</script>
