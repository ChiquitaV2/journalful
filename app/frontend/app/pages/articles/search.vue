<template>
  <div class="container py-8">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-4">Search & Add Articles</h1>
      <p class="text-gray-600">
        Search our database or add new articles by DOI.
      </p>
    </div>

    <!-- Search Section -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8">
      <div class="flex flex-col lg:flex-row gap-4">
        <div class="flex-1">
          <Input
            type="search"
            placeholder="Search articles, authors, keywords, or enter DOI..."
            v-model="searchQuery"
            @keyup.enter="handleSearch"
            class="w-full"
          />
        </div>
        <div class="flex gap-2">
          <Button 
            @click="handleSearch"
            :loading="isSearching"
            :disabled="!searchQuery.trim()"
          >
            <Icon name="heroicons:magnifying-glass" class="h-4 w-4 mr-2" />
            Search
          </Button>
          <Button
            variant="outline"
            @click="handleAddByDOI"
            :loading="isAddingByDOI"
            :disabled="!isDOIFormat(searchQuery)"
          >
            <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
            Add by DOI
          </Button>
        </div>
      </div>
      
      <div class="mt-4 text-sm text-gray-600">
        <p>
          <strong>Tip:</strong> Enter a DOI (like "10.1000/182") to add an article directly, 
          or search for keywords to find existing articles.
        </p>
      </div>
    </div>

    <!-- Search Results -->
    <div v-if="hasSearched">
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-2xl font-bold text-gray-900">
          Search Results
          <span v-if="searchResults.length" class="text-gray-500 text-lg font-normal">
            ({{ searchResults.length }} found)
          </span>
        </h2>
        
        <div v-if="searchResults.length > 0" class="flex items-center gap-2">
          <span class="text-sm text-gray-600">Sort by:</span>
          <select 
            v-model="sortBy" 
            @change="sortResults"
            class="form-input py-1 px-2 text-sm"
          >
            <option value="relevance">Relevance</option>
            <option value="year">Publication Year</option>
            <option value="title">Title</option>
            <option value="author">Author</option>
          </select>
        </div>
      </div>
      
      <div v-if="isSearching" class="flex justify-center py-12">
        <LoadingSpinner :size="40" />
      </div>
      
      <div v-else-if="searchResults.length === 0" class="text-center py-12">
        <Icon name="heroicons:document-magnifying-glass" class="h-16 w-16 text-gray-400 mx-auto mb-4" />
        <h3 class="text-lg font-semibold text-gray-900 mb-2">No articles found</h3>
        <p class="text-gray-600 mb-4">
          No articles match your search criteria. Try different keywords or add a new article by DOI.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <Button
            variant="outline"
            @click="clearSearch"
          >
            Clear Search
          </Button>
          <Button
            v-if="isDOIFormat(searchQuery)"
            @click="handleAddByDOI"
            :loading="isAddingByDOI"
          >
            <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
            Add "{{ searchQuery }}" by DOI
          </Button>
        </div>
      </div>
      
      <div v-else class="space-y-6">
        <ArticleCard
          v-for="article in searchResults"
          :key="article.id"
          :article="article"
          @library-updated="handleLibraryUpdate"
        />
      </div>
    </div>

    <!-- Recent Articles -->
    <div v-else class="space-y-8">
      <div>
        <h2 class="text-2xl font-bold text-gray-900 mb-6">Recent Articles</h2>
        <div v-if="isLoadingRecent" class="flex justify-center py-12">
          <LoadingSpinner :size="40" />
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <ArticleCard
            v-for="article in recentArticles"
            :key="article.id"
            :article="article"
            @library-updated="handleLibraryUpdate"
          />
        </div>
      </div>
      
      <!-- Search Tips -->
      <div class="bg-blue-50 rounded-lg p-6">
        <h3 class="text-lg font-semibold text-blue-900 mb-4">Search Tips</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm text-blue-800">
          <div>
            <h4 class="font-medium mb-2">By Keywords</h4>
            <ul class="space-y-1">
              <li>• Search article titles and abstracts</li>
              <li>• Use specific terms for better results</li>
              <li>• Try author names or journal names</li>
            </ul>
          </div>
          <div>
            <h4 class="font-medium mb-2">By DOI</h4>
            <ul class="space-y-1">
              <li>• Enter DOI like "10.1000/182"</li>
              <li>• Automatically adds if not in database</li>
              <li>• Fetches metadata from publishers</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Success Modal -->
    <Modal v-model:show="showSuccessModal" title="Article Added Successfully">
      <div v-if="addedArticle" class="space-y-4">
        <p class="text-gray-600">
          The article has been successfully added to the database.
        </p>
        
        <div class="bg-gray-50 rounded-lg p-4">
          <h4 class="font-semibold text-gray-900 mb-2">{{ addedArticle.title }}</h4>
          <p class="text-sm text-gray-600">
            by {{ addedArticle.authors.map(a => a.name).join(', ') }}
          </p>
          <p class="text-sm text-gray-500 mt-2">
            {{ addedArticle.journalName }} • {{ addedArticle.publicationYear }}
          </p>
        </div>
        
        <div class="flex gap-2">
          <Button 
            size="sm" 
            variant="outline"
            @click="navigateTo(`/articles/${addedArticle.id}`)"
          >
            View Article
          </Button>
          <Button 
            size="sm" 
            variant="primary"
            @click="addToLibrary(addedArticle)"
          >
            Add to Library
          </Button>
        </div>
      </div>
      
      <template #footer>
        <Button variant="outline" @click="showSuccessModal = false">
          Close
        </Button>
      </template>
    </Modal>
  </div>
</template>

<script setup lang="ts">
import type { Article } from '#shared/grpc/articles/v1/article'

// Meta
useSeoMeta({
  title: 'Search & Add Articles - Journalful',
  description: 'Search our database of academic articles or add new articles by DOI.',
})

// State
const searchQuery = ref('')
const searchResults = ref<Article[]>([])
const recentArticles = ref<Article[]>([])
const isSearching = ref(false)
const isLoadingRecent = ref(false)
const isAddingByDOI = ref(false)
const hasSearched = ref(false)
const sortBy = ref('relevance')
const showSuccessModal = ref(false)
const addedArticle = ref<Article | null>(null)

// Methods
const handleSearch = async () => {
  if (!searchQuery.value.trim()) return
  
  isSearching.value = true
  hasSearched.value = true
  
  try {
    // TODO: Replace with actual search API
    await new Promise(resolve => setTimeout(resolve, 800))
    
    // Mock search results
    searchResults.value = [
      {
        id: 1,
        doi: '10.1000/182',
        title: 'Machine Learning in Healthcare: A Comprehensive Review',
        authors: [
          { 
            id: 1, 
            name: 'Dr. Sarah Johnson', 
            profileId: 1, 
            createdAt: new Date(), 
            updatedAt: new Date() 
          }
        ],
        abstract: 'This paper presents a comprehensive review of machine learning applications in healthcare.',
        publicationYear: 2023,
        journalName: 'Journal of Medical Informatics',
        createdAt: new Date('2023-01-15'),
        updatedAt: new Date('2023-01-15')
      }
    ].filter(article => 
      article.title.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      article.authors.some(author => author.name.toLowerCase().includes(searchQuery.value.toLowerCase()))
    )
    
  } catch (error) {
    console.error('Search error:', error)
  } finally {
    isSearching.value = false
  }
}

const handleAddByDOI = async () => {
  if (!isDOIFormat(searchQuery.value)) return
  
  isAddingByDOI.value = true
  
  try {
    // TODO: Replace with actual DOI lookup API
    await new Promise(resolve => setTimeout(resolve, 1500))
    
    // Mock article creation
    const newArticle: Article = {
      id: Date.now(),
      doi: searchQuery.value,
      title: `Article Added by DOI: ${searchQuery.value}`,
      authors: [
        { 
          id: 999, 
          name: 'Dr. Example Author', 
          profileId: 999, 
          createdAt: new Date(), 
          updatedAt: new Date() 
        }
      ],
      abstract: 'This article was successfully added to the database using DOI lookup.',
      publicationYear: 2024,
      journalName: 'Example Journal',
      createdAt: new Date(),
      updatedAt: new Date()
    }
    
    addedArticle.value = newArticle
    showSuccessModal.value = true
    
    // Add to recent articles
    recentArticles.value.unshift(newArticle)
    
  } catch (error) {
    console.error('DOI lookup error:', error)
  } finally {
    isAddingByDOI.value = false
  }
}

const isDOIFormat = (text: string) => {
  // Simple DOI pattern matching
  const doiPattern = /^10\.\d{4,}\/[^\s]+$/
  return doiPattern.test(text.trim())
}

const clearSearch = () => {
  searchQuery.value = ''
  searchResults.value = []
  hasSearched.value = false
}

const sortResults = () => {
  // TODO: Implement proper sorting
  console.log('Sorting by:', sortBy.value)
}

const handleLibraryUpdate = (articleId: number, isInLibrary: boolean) => {
  console.log('Library updated:', articleId, isInLibrary)
}

const addToLibrary = (article: Article) => {
  // TODO: Implement add to library
  console.log('Adding to library:', article.id)
  showSuccessModal.value = false
}

const loadRecentArticles = async () => {
  isLoadingRecent.value = true
  
  try {
    // TODO: Replace with actual API call
    await new Promise(resolve => setTimeout(resolve, 500))
    
    // Mock recent articles
    recentArticles.value = [
      {
        id: 1,
        doi: '10.1000/182',
        title: 'Machine Learning in Healthcare: A Comprehensive Review',
        authors: [
          { 
            id: 1, 
            name: 'Dr. Sarah Johnson', 
            profileId: 1, 
            createdAt: new Date(), 
            updatedAt: new Date() 
          }
        ],
        abstract: 'This paper presents a comprehensive review of machine learning applications in healthcare.',
        publicationYear: 2023,
        journalName: 'Journal of Medical Informatics',
        createdAt: new Date('2023-01-15'),
        updatedAt: new Date('2023-01-15')
      },
      {
        id: 2,
        doi: '10.1000/183',
        title: 'Climate Change Impact on Agricultural Productivity',
        authors: [
          { 
            id: 2, 
            name: 'Dr. Emily Rodriguez', 
            profileId: 2, 
            createdAt: new Date(), 
            updatedAt: new Date() 
          }
        ],
        abstract: 'We analyze the effects of climate change on global agricultural productivity.',
        publicationYear: 2023,
        journalName: 'Environmental Science & Policy',
        createdAt: new Date('2023-02-20'),
        updatedAt: new Date('2023-02-20')
      }
    ]
    
  } catch (error) {
    console.error('Error loading recent articles:', error)
  } finally {
    isLoadingRecent.value = false
  }
}

// Initialize
onMounted(() => {
  // Check for search query from URL
  const route = useRoute()
  if (route.query.q) {
    searchQuery.value = route.query.q as string
    handleSearch()
  } else {
    loadRecentArticles()
  }
})

// Watch for route changes
watch(() => useRoute().query, (newQuery) => {
  if (newQuery.q && newQuery.q !== searchQuery.value) {
    searchQuery.value = newQuery.q as string
    handleSearch()
  }
})
</script>
