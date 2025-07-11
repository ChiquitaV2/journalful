<template>
  <div class="container py-8">
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-4">Research Articles</h1>
      <p class="text-gray-600">
        Discover and explore academic articles from our extensive database.
      </p>
    </div>

    <!-- Search and Filters -->
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 mb-8">
      <div class="flex flex-col lg:flex-row gap-4 items-center">
        <!-- Search Input -->
        <div class="flex-1">
          <Input
            type="search"
            placeholder="Search articles, authors, or keywords..."
            v-model="searchQuery"
            @keyup.enter="handleSearch"
          />
        </div>
        
        <!-- Filter Dropdown -->
        <div class="flex gap-2">
          <Button
            variant="outline"
            @click="showFilters = !showFilters"
            :class="{ 'bg-gray-50': showFilters }"
          >
            <Icon name="heroicons:funnel" class="h-4 w-4 mr-2" />
            Filters
            <Icon name="heroicons:chevron-down" class="h-4 w-4 ml-2" />
          </Button>
          
          <Button
            variant="outline"
            @click="showSort = !showSort"
            :class="{ 'bg-gray-50': showSort }"
          >
            <Icon name="heroicons:bars-arrow-up" class="h-4 w-4 mr-2" />
            Sort
            <Icon name="heroicons:chevron-down" class="h-4 w-4 ml-2" />
          </Button>
        </div>
      </div>
      
      <!-- Expanded Filters -->
      <Transition
        enter-active-class="transition ease-out duration-200"
        enter-from-class="opacity-0 max-h-0"
        enter-to-class="opacity-100 max-h-96"
        leave-active-class="transition ease-in duration-150"
        leave-from-class="opacity-100 max-h-96"
        leave-to-class="opacity-0 max-h-0"
      >
        <div v-if="showFilters" class="mt-6 pt-6 border-t border-gray-200">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label class="form-label">Publication Year</label>
              <div class="flex gap-2">
                <Input
                  type="number"
                  placeholder="From"
                  v-model="filters.yearFrom"
                  class="w-full"
                />
                <Input
                  type="number"
                  placeholder="To"
                  v-model="filters.yearTo"
                  class="w-full"
                />
              </div>
            </div>
            
            <div>
              <label class="form-label">Journal</label>
              <Input
                type="text"
                placeholder="Filter by journal..."
                v-model="filters.journal"
              />
            </div>
            
            <div>
              <label class="form-label">Author</label>
              <Input
                type="text"
                placeholder="Filter by author..."
                v-model="filters.author"
              />
            </div>
          </div>
          
          <div class="flex justify-between items-center mt-4">
            <Button
              variant="outline"
              size="sm"
              @click="clearFilters"
            >
              Clear All
            </Button>
            <Button
              variant="primary"
              size="sm"
              @click="applyFilters"
            >
              Apply Filters
            </Button>
          </div>
        </div>
      </Transition>
      
      <!-- Sort Options -->
      <Transition
        enter-active-class="transition ease-out duration-200"
        enter-from-class="opacity-0 max-h-0"
        enter-to-class="opacity-100 max-h-96"
        leave-active-class="transition ease-in duration-150"
        leave-from-class="opacity-100 max-h-96"
        leave-to-class="opacity-0 max-h-0"
      >
        <div v-if="showSort" class="mt-6 pt-6 border-t border-gray-200">
          <div class="grid grid-cols-2 md:grid-cols-4 gap-2">
            <button
              v-for="option in sortOptions"
              :key="option.value"
              @click="setSortOption(option.value)"
              :class="[
                'text-left p-3 rounded-md border transition-colors',
                currentSort === option.value
                  ? 'border-primary bg-blue-50 text-primary'
                  : 'border-gray-200 hover:bg-gray-50'
              ]"
            >
              <div class="font-medium">{{ option.label }}</div>
              <div class="text-sm text-gray-600">{{ option.description }}</div>
            </button>
          </div>
        </div>
      </Transition>
    </div>

    <!-- Active Filters -->
    <div v-if="hasActiveFilters" class="mb-6">
      <div class="flex flex-wrap gap-2 items-center">
        <span class="text-sm text-gray-600">Active filters:</span>
        <Tag
          v-if="searchQuery"
          variant="primary"
          @click="searchQuery = ''"
          class="cursor-pointer"
        >
          Search: {{ searchQuery }}
          <Icon name="heroicons:x-mark" class="h-3 w-3 ml-1" />
        </Tag>
        <Tag
          v-if="filters.yearFrom || filters.yearTo"
          variant="primary"
          @click="filters.yearFrom = ''; filters.yearTo = ''"
          class="cursor-pointer"
        >
          Year: {{ filters.yearFrom || 'any' }} - {{ filters.yearTo || 'any' }}
          <Icon name="heroicons:x-mark" class="h-3 w-3 ml-1" />
        </Tag>
        <Tag
          v-if="filters.journal"
          variant="primary"
          @click="filters.journal = ''"
          class="cursor-pointer"
        >
          Journal: {{ filters.journal }}
          <Icon name="heroicons:x-mark" class="h-3 w-3 ml-1" />
        </Tag>
        <Tag
          v-if="filters.author"
          variant="primary"
          @click="filters.author = ''"
          class="cursor-pointer"
        >
          Author: {{ filters.author }}
          <Icon name="heroicons:x-mark" class="h-3 w-3 ml-1" />
        </Tag>
      </div>
    </div>

    <!-- Articles Grid -->
    <div v-if="loading" class="flex justify-center py-12">
      <LoadingSpinner :size="40" />
    </div>
    
    <div v-else-if="articles.length === 0" class="text-center py-12">
      <Icon name="heroicons:document-text" class="h-16 w-16 text-gray-400 mx-auto mb-4" />
      <h3 class="text-lg font-semibold text-gray-900 mb-2">No articles found</h3>
      <p class="text-gray-600 mb-4">
        {{ searchQuery ? 'Try adjusting your search terms or filters.' : 'No articles available at the moment.' }}
      </p>
      <Button variant="primary" @click="navigateTo('/articles/search')">
        <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
        Add New Article
      </Button>
    </div>
    
    <div v-else class="space-y-6">
      <!-- Results Count -->
      <div class="flex items-center justify-between">
        <p class="text-sm text-gray-600">
          Showing {{ articles.length }} of {{ totalArticles }} articles
        </p>
        <div class="flex items-center gap-2">
          <span class="text-sm text-gray-600">View:</span>
          <Button
            variant="ghost"
            size="sm"
            @click="viewMode = 'grid'"
            :class="{ 'bg-gray-100': viewMode === 'grid' }"
          >
            <Icon name="heroicons:squares-2x2" class="h-4 w-4" />
          </Button>
          <Button
            variant="ghost"
            size="sm"
            @click="viewMode = 'list'"
            :class="{ 'bg-gray-100': viewMode === 'list' }"
          >
            <Icon name="heroicons:list-bullet" class="h-4 w-4" />
          </Button>
        </div>
      </div>
      
      <!-- Articles -->
      <div
        :class="[
          'grid gap-6',
          viewMode === 'grid' 
            ? 'grid-cols-1 md:grid-cols-2 lg:grid-cols-3' 
            : 'grid-cols-1'
        ]"
      >
        <ArticleCard
          v-for="article in articles"
          :key="article.id"
          :article="article"
          @library-updated="handleLibraryUpdate"
        />
      </div>
      
      <!-- Pagination -->
      <div class="mt-8">
        <Pagination
          :current-page="currentPage"
          :total-pages="totalPages"
          :total-items="totalArticles"
          :items-per-page="itemsPerPage"
          @update:current-page="handlePageChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Article } from '#shared/grpc/articles/v1/article'

// Meta tags
useSeoMeta({
  title: 'Research Articles - Journalful',
  description: 'Browse and discover academic articles from our extensive research database.',
})

// Reactive state
const loading = ref(false)
const articles = ref<Article[]>([])
const totalArticles = ref(0)
const currentPage = ref(1)
const itemsPerPage = 20
const searchQuery = ref('')
const showFilters = ref(false)
const showSort = ref(false)
const viewMode = ref<'grid' | 'list'>('grid')
const currentSort = ref('newest')

// Filters
const filters = reactive({
  yearFrom: '',
  yearTo: '',
  journal: '',
  author: ''
})

// Sort options
const sortOptions = [
  {
    value: 'newest',
    label: 'Newest First',
    description: 'Most recently published'
  },
  {
    value: 'oldest',
    label: 'Oldest First',
    description: 'Earliest published'
  },
  {
    value: 'title',
    label: 'Title A-Z',
    description: 'Alphabetical by title'
  },
  {
    value: 'relevance',
    label: 'Most Relevant',
    description: 'Best match for search'
  }
]

// Computed properties
const totalPages = computed(() => Math.ceil(totalArticles.value / itemsPerPage))

const hasActiveFilters = computed(() => {
  return searchQuery.value || 
         filters.yearFrom || 
         filters.yearTo || 
         filters.journal || 
         filters.author
})

// Methods
const handleSearch = () => {
  currentPage.value = 1
  fetchArticles()
}

const applyFilters = () => {
  currentPage.value = 1
  showFilters.value = false
  fetchArticles()
}

const clearFilters = () => {
  searchQuery.value = ''
  filters.yearFrom = ''
  filters.yearTo = ''
  filters.journal = ''
  filters.author = ''
  currentPage.value = 1
  fetchArticles()
}

const setSortOption = (value: string) => {
  currentSort.value = value
  showSort.value = false
  fetchArticles()
}

const handlePageChange = (page: number) => {
  currentPage.value = page
  fetchArticles()
}

const handleLibraryUpdate = (articleId: number, isInLibrary: boolean) => {
  // Handle library update logic
  console.log('Article library updated:', articleId, isInLibrary)
}

const fetchArticles = async () => {
  loading.value = true
  try {
    // TODO: Use actual composable when ready
    // const { fetchArticles: fetch } = useArticles()
    // const result = await fetch({
    //   page: currentPage.value,
    //   pageSize: itemsPerPage
    // })
    
    // For now, simulate API call
    await new Promise(resolve => setTimeout(resolve, 500))
    
    // Mock data
    articles.value = [
      {
        id: 1,
        doi: '10.1000/182',
        title: 'Machine Learning in Healthcare: A Comprehensive Review',
        authors: [
          { id: 1, name: 'Dr. Sarah Johnson', profileId: 1, createdAt: new Date(), updatedAt: new Date() }
        ],
        abstract: 'This paper presents a comprehensive review of machine learning applications in healthcare.',
        publicationYear: 2023,
        journalName: 'Journal of Medical Informatics',
        createdAt: new Date('2023-01-15'),
        updatedAt: new Date('2023-01-15')
      }
    ]
    totalArticles.value = 1
    
  } catch (error) {
    console.error('Error fetching articles:', error)
  } finally {
    loading.value = false
  }
}

// Initialize
onMounted(() => {
  // Get query parameters
  const route = useRoute()
  if (route.query.q) {
    searchQuery.value = route.query.q as string
  }
  
  fetchArticles()
})

// Watch for route changes
watch(() => useRoute().query, (newQuery) => {
  if (newQuery.q !== searchQuery.value) {
    searchQuery.value = (newQuery.q as string) || ''
    fetchArticles()
  }
})
</script>
