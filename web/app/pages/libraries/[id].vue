<template>
  <div class="min-h-screen">
    <!-- Mobile Header -->
    <header class="glass-nav fixed top-0 left-0 right-0 z-40 h-16">
      <div class="flex items-center justify-between h-full mobile-padding">
        <div class="flex items-center space-x-3">
          <NuxtLink to="/libraries" class="glass-button !p-2">
            <Icon name="heroicons:arrow-left" class="h-6 w-6" />
          </NuxtLink>
          <h1 class="text-lg font-bold text-white truncate">Library</h1>
        </div>
        
        <div class="flex items-center space-x-2">
          <button class="glass-button !p-2" @click="addArticle">
            <Icon name="heroicons:plus" class="h-5 w-5" />
          </button>
          <button class="glass-button !p-2" @click="shareLibrary">
            <Icon name="heroicons:share" class="h-5 w-5" />
          </button>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="pt-16">
      <!-- Loading State -->
      <div v-if="pending" class="flex items-center justify-center min-h-[50vh]">
        <div class="glass-card p-8 text-center">
          <div class="animate-pulse">
            <Icon name="heroicons:book-open" class="h-12 w-12 mx-auto mb-4 text-purple-400" />
            <p class="text-gray-300">Loading library...</p>
          </div>
        </div>
      </div>

      <!-- Error State -->
      <div v-else-if="error" class="flex items-center justify-center min-h-[50vh]">
        <div class="glass-card p-8 text-center max-w-md mobile-padding">
          <Icon name="heroicons:exclamation-triangle" class="h-12 w-12 mx-auto mb-4 text-red-400" />
          <h3 class="text-xl font-semibold mb-2">Library Not Found</h3>
          <p class="text-gray-300 mb-4">The library you're looking for doesn't exist.</p>
          <NuxtLink to="/libraries" class="glass-button">
            <Icon name="heroicons:arrow-left" class="h-4 w-4 mr-2" />
            Back to Libraries
          </NuxtLink>
        </div>
      </div>

      <!-- Library Content -->
      <div v-else-if="library" class="mobile-padding max-w-6xl mx-auto pb-20">
        <!-- Library Header -->
        <div class="glass-card p-6 mb-6">
          <div class="flex items-start justify-between mb-4">
            <div class="flex items-center gap-2">
              <button class="glass-button" @click="editLibrary">
                <Icon name="heroicons:pencil" class="h-4 w-4 mr-2" />
                Edit
              </button>
            </div>
          </div>

        <div class="flex items-start gap-4">
          <div 
            class="w-16 h-16 rounded-xl flex items-center justify-center text-2xl"
            :style="{ background: library.color || 'linear-gradient(135deg, #8b5cf6, #3b82f6)' }"
          >
            {{ library.icon || '📚' }}
          </div>
          
          <div class="flex-1">
            <h1 class="text-2xl md:text-3xl font-bold text-white mb-2">
              {{ library.name }}
            </h1>
            <p v-if="library.description" class="text-gray-300 mb-4">
              {{ library.description }}
            </p>
            
            <div class="flex flex-wrap items-center gap-4 text-sm text-gray-400">
              <div class="flex items-center gap-1">
                <Icon name="heroicons:document-text" class="h-4 w-4" />
                <span>{{ library.articleCount || 0 }} articles</span>
              </div>
              <div class="flex items-center gap-1">
                <Icon name="heroicons:calendar" class="h-4 w-4" />
                <span>Created {{ formatDate(library.createdAt) }}</span>
              </div>
              <div class="flex items-center gap-1">
                <Icon name="heroicons:clock" class="h-4 w-4" />
                <span>Updated {{ formatDate(library.updatedAt) }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Library Stats -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
        <div class="glass-card p-4 text-center">
          <div class="text-2xl font-bold text-purple-400 mb-1">{{ stats.total }}</div>
          <div class="text-sm text-gray-400">Total Articles</div>
        </div>
        <div class="glass-card p-4 text-center">
          <div class="text-2xl font-bold text-green-400 mb-1">{{ stats.completed }}</div>
          <div class="text-sm text-gray-400">Completed</div>
        </div>
        <div class="glass-card p-4 text-center">
          <div class="text-2xl font-bold text-yellow-400 mb-1">{{ stats.inProgress }}</div>
          <div class="text-sm text-gray-400">In Progress</div>
        </div>
        <div class="glass-card p-4 text-center">
          <div class="text-2xl font-bold text-gray-400 mb-1">{{ stats.unread }}</div>
          <div class="text-sm text-gray-400">Unread</div>
        </div>
      </div>

      <!-- Controls -->
      <div class="glass-card p-4 mb-6">
        <div class="flex flex-col md:flex-row gap-4 items-start md:items-center justify-between">
          <!-- Search and Filter -->
          <div class="flex-1 flex flex-col sm:flex-row gap-3 w-full md:w-auto">
            <div class="relative flex-1 max-w-md ">
              <Icon name="heroicons:magnifying-glass" class="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-gray-400" />
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Search articles..."
                class="glass-input w-full pl-10 py-2 px-4 "
              />
            </div>
            
            <select v-model="filterStatus" class="glass-select">
              <option value="">All Status</option>
              <option value="unread">Unread</option>
              <option value="in-progress">In Progress</option>
              <option value="completed">Completed</option>
            </select>

            <select v-model="sortBy" class="glass-select">
              <option value="dateAdded">Date Added</option>
              <option value="title">Title</option>
              <option value="publishedDate">Published Date</option>
              <option value="readingProgress">Progress</option>
            </select>
          </div>

          <!-- Action Buttons -->
          <div class="flex gap-3">
            <button class="glass-button" @click="toggleView">
              <Icon :name="viewMode === 'grid' ? 'heroicons:list-bullet' : 'heroicons:squares-2x2'" class="h-4 w-4 mr-2" />
              {{ viewMode === 'grid' ? 'List' : 'Grid' }}
            </button>
            <button class="glass-button-primary" @click="addArticle">
              <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
              Add Article
            </button>
          </div>
        </div>
      </div>

      <!-- Articles List/Grid -->
      <div v-if="filteredArticles.length > 0">
        <!-- Grid View -->
        <div v-if="viewMode === 'grid'" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <ArticleCard
            v-for="article in filteredArticles"
            :key="article.id"
            :article="article"
            @update-status="updateArticleStatus"
            @toggle-favorite="toggleArticleFavorite"
            @remove="removeArticle"
          />
        </div>

        <!-- List View -->
        <div v-else class="space-y-4">
          <div 
            v-for="article in filteredArticles"
            :key="article.id"
            class="glass-card p-4 hover:scale-[1.02] transition-transform cursor-pointer"
            @click="navigateToArticle(article.id)"
          >
            <div class="flex items-start gap-4">
              <div class="flex-1">
                <h3 class="font-semibold text-white mb-2 line-clamp-1">{{ article.title }}</h3>
                <p class="text-sm text-gray-300 mb-2">{{ formatAuthors(article.authors) }}</p>
                
                <div class="flex flex-wrap items-center gap-4 text-xs text-gray-400 mb-3">
                  <span>{{ article.publishedDate }}</span>
                  <span>{{ article.journal }}</span>
                  <span>Added {{ formatDate(article.dateAdded) }}</span>
                </div>

                <div class="flex flex-wrap gap-2">
                  <span 
                    v-for="tag in article.tags?.slice(0, 3)" 
                    :key="tag"
                    class="glass-tag-sm"
                  >
                    {{ tag }}
                  </span>
                  <span v-if="article.tags?.length > 3" class="glass-tag-sm">
                    +{{ article.tags.length - 3 }}
                  </span>
                </div>
              </div>

              <div class="flex flex-col items-end gap-2">
                <StatusBadge :status="article.readingStatus" size="sm" />
                
                <div class="flex items-center gap-1">
                  <button 
                    @click.stop="toggleArticleFavorite(article)"
                    class="p-1 rounded hover:bg-white/10"
                  >
                    <Icon 
                      :name="article.isFavorite ? 'heroicons:heart-solid' : 'heroicons:heart'" 
                      :class="['h-4 w-4', article.isFavorite ? 'text-red-400' : 'text-gray-400']" 
                    />
                  </button>
                  
                  <button 
                    @click.stop="showArticleMenu(article)"
                    class="p-1 rounded hover:bg-white/10"
                  >
                    <Icon name="heroicons:ellipsis-horizontal" class="h-4 w-4 text-gray-400" />
                  </button>
                </div>

                <div v-if="article.readingProgress > 0" class="w-16">
                  <div class="text-xs text-gray-400 mb-1">{{ article.readingProgress }}%</div>
                  <div class="w-full h-1 bg-gray-600 rounded-full overflow-hidden">
                    <div 
                      class="h-full bg-gradient-to-r from-purple-500 to-blue-500 transition-all"
                      :style="{ width: `${article.readingProgress}%` }"
                    ></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-else class="text-center py-16">
        <div class="glass-card p-8 max-w-md mx-auto">
          <Icon name="heroicons:document-plus" class="h-16 w-16 mx-auto mb-4 text-gray-500" />
          <h3 class="text-xl font-semibold mb-2">No Articles Found</h3>
          <p class="text-gray-400 mb-6">
            {{ searchQuery || filterStatus ? 'No articles match your search criteria.' : 'This library is empty. Add your first article to get started.' }}
          </p>
          <button v-if="!searchQuery && !filterStatus" class="glass-button-primary" @click="addArticle">
            <Icon name="heroicons:plus" class="h-4 w-4 mr-2" />
            Add First Article
          </button>
          <button v-else class="glass-button" @click="clearFilters">
            <Icon name="heroicons:x-mark" class="h-4 w-4 mr-2" />
            Clear Filters
          </button>
        </div>
      </div>
    </div>
    </main>

    <!-- Mobile Bottom Navigation -->
    <MobileBottomNav />
    <AddArticleModal
        v-if="showAddModal"
        @close="showAddModal = false"
        @article-added="handleArticleAdded"
    />

    <!-- Floating Action Button -->
    <div class="fixed bottom-20 right-6 z-20 md:bottom-6">
      <button class="glass-fab" @click="addArticle">
        <Icon name="heroicons:plus" class="h-6 w-6" />
      </button>
    </div>
  </div>
</template>

<script setup>
const route = useRoute()
const router = useRouter()
const { success, error: showError, warning } = useToast()

// Reactive state
const searchQuery = ref('')
const filterStatus = ref('')
const sortBy = ref('dateAdded')
const viewMode = ref('grid')

// Mock data - replace with real API call
const { data: library, pending, error } = await useLazyFetch(`/api/libraries/${route.params.id}`, {
  default: () => ({
    id: route.params.id,
    name: "Machine Learning Research",
    description: "A curated collection of papers on machine learning, deep learning, and AI research.",
    icon: "🤖",
    color: "linear-gradient(135deg, #8b5cf6, #3b82f6)",
    createdAt: "2024-01-15",
    updatedAt: "2024-07-28",
    articleCount: 24,
    articles: [
      {
        id: "1",
        title: "Attention Is All You Need",
        authors: [{ name: "Vaswani et al." }],
        abstract: "The dominant sequence transduction models are based on complex recurrent or convolutional neural networks...",
        url: "https://arxiv.org/abs/1706.03762",
        publishedDate: "2017-06-12",
        journal: "NeurIPS",
        tags: ["Transformers", "Attention", "NLP"],
        readingStatus: "completed",
        readingProgress: 100,
        dateAdded: "2024-01-20",
        dateStarted: "2024-01-22",
        dateCompleted: "2024-01-25",
        isFavorite: true
      },
      {
        id: "2",
        title: "BERT: Pre-training of Deep Bidirectional Transformers",
        authors: [{ name: "Devlin et al." }],
        abstract: "We introduce a new language representation model called BERT...",
        url: "https://arxiv.org/abs/1810.04805",
        publishedDate: "2018-10-11",
        journal: "NAACL",
        tags: ["BERT", "Language Models", "NLP"],
        readingStatus: "in-progress",
        readingProgress: 45,
        dateAdded: "2024-02-01",
        dateStarted: "2024-02-05",
        isFavorite: false
      },
      {
        id: "3",
        title: "GPT-3: Language Models are Few-Shot Learners",
        authors: [{ name: "Brown et al." }],
        abstract: "Recent work has demonstrated substantial gains on many NLP tasks...",
        url: "https://arxiv.org/abs/2005.14165",
        publishedDate: "2020-05-28",
        journal: "NeurIPS",
        tags: ["GPT", "Language Models", "Few-Shot Learning"],
        readingStatus: "unread",
        readingProgress: 0,
        dateAdded: "2024-03-10",
        isFavorite: false
      },
      {
        id: "4",
        title: "ResNet: Deep Residual Learning for Image Recognition",
        authors: [{ name: "He et al." }],
        abstract: "Deeper neural networks are more difficult to train...",
        url: "https://arxiv.org/abs/1512.03385",
        publishedDate: "2015-12-10",
        journal: "CVPR",
        tags: ["ResNet", "Computer Vision", "CNN"],
        readingStatus: "completed",
        readingProgress: 100,
        dateAdded: "2024-01-30",
        dateCompleted: "2024-02-15",
        isFavorite: true
      }
    ]
  })
})

// Computed
const stats = computed(() => {
  const articles = library.value?.articles || []
  return {
    total: articles.length,
    completed: articles.filter(a => a.readingStatus === 'completed').length,
    inProgress: articles.filter(a => a.readingStatus === 'in-progress').length,
    unread: articles.filter(a => a.readingStatus === 'unread').length
  }
})

const filteredArticles = computed(() => {
  let articles = library.value?.articles || []

  // Filter by search query
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    articles = articles.filter(article =>
        article.title.toLowerCase().includes(query) ||
        article.authors.some(author => author.name.toLowerCase().includes(query)) ||
        article.tags?.some(tag => tag.toLowerCase().includes(query))
    )
  }

  // Filter by status
  if (filterStatus.value) {
    articles = articles.filter(article => article.readingStatus === filterStatus.value)
  }

  // Sort articles
  articles.sort((a, b) => {
    switch (sortBy.value) {
      case 'title':
        return a.title.localeCompare(b.title)
      case 'publishedDate':
        return new Date(b.publishedDate) - new Date(a.publishedDate)
      case 'readingProgress':
        return b.readingProgress - a.readingProgress
      case 'dateAdded':
      default:
        return new Date(b.dateAdded) - new Date(a.dateAdded)
    }
  })

  return articles
})

// Methods
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

const formatAuthors = (authors) => {
  if (!authors || authors.length === 0) return 'Unknown'
  if (authors.length === 1) return authors[0].name
  if (authors.length === 2) return `${authors[0].name} & ${authors[1].name}`
  return `${authors[0].name} et al.`
}

const toggleView = () => {
  viewMode.value = viewMode.value === 'grid' ? 'list' : 'grid'
}

const clearFilters = () => {
  searchQuery.value = ''
  filterStatus.value = ''
}

const navigateToArticle = (id) => {
  navigateTo(`/articles/${id}`)
}

const updateArticleStatus = async (article, newStatus) => {
  try {
    const index = library.value.articles.findIndex(a => a.id === article.id)
    if (index !== -1) {
      const previousStatus = library.value.articles[index].readingStatus

      library.value.articles[index].readingStatus = newStatus
      if (newStatus === 'completed') {
        library.value.articles[index].readingProgress = 100
        library.value.articles[index].dateCompleted = new Date().toISOString().split('T')[0]
      }

      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 300))

      success('Status Updated', `Article marked as ${newStatus.replace('-', ' ')}.`)
    }
  } catch (err) {
    // Revert state on error
    const index = library.value.articles.findIndex(a => a.id === article.id)
    if (index !== -1) {
      library.value.articles[index].readingStatus = previousStatus
    }
    showError('Update Failed', 'Unable to update article status. Please try again.')
  }
}

const toggleArticleFavorite = async (article) => {
  try {
    const index = library.value.articles.findIndex(a => a.id === article.id)
    if (index !== -1) {
      const previousState = library.value.articles[index].isFavorite
      library.value.articles[index].isFavorite = !library.value.articles[index].isFavorite

      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 300))

      success(
          library.value.articles[index].isFavorite ? 'Added to Favorites' : 'Removed from Favorites'
      )
    }
  } catch (err) {
    // Revert state on error
    const index = library.value.articles.findIndex(a => a.id === article.id)
    if (index !== -1) {
      library.value.articles[index].isFavorite = !library.value.articles[index].isFavorite
    }
    showError('Update Failed', 'Unable to update favorite status. Please try again.')
  }
}

const removeArticle = async (article) => {
  try {
    if (confirm('Are you sure you want to remove this article from the library?')) {
      const index = library.value.articles.findIndex(a => a.id === article.id)
      if (index !== -1) {
        // Store for potential undo
        const removedArticle = library.value.articles[index]

        library.value.articles.splice(index, 1)
        library.value.articleCount = library.value.articles.length

        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 300))

        success('Article Removed', `"${removedArticle.title}" has been removed from the library.`)
      }
    }
  } catch (err) {
    showError('Remove Failed', 'Unable to remove article. Please try again.')
  }
}

const showAddModal = ref(false)
const addArticle = () => {
  // Open add article modal or navigate to add page
  // router.push(`/libraries/${library.value.id}/add-article`)
  showAddModal.value = true
}
const handleArticleAdded = async (article) => {
  try {
    // Add article to library via API
    const response = await $fetch(`/api/libraries/${library.value.id}/articles`, {
      method: 'POST',
      body: {
        articleId: article.id,
        readingStatus: article.readingStatus || 'unread',
        notes: article.notes || null
      }
    })
    
    if (response.success) {
      // Update local state
      library.value.articles.push(article)
      library.value.articleCount = library.value.articles.length
      
      success('Article Added', 'Article has been added to your library.')
      showAddModal.value = false
    }
  } catch (err) {
    console.error('Failed to add article to library:', err)
    showError('Add Failed', 'Unable to add article to library. Please try again.')
  }
}
const editLibrary = () => {
  router.push(`/libraries/${library.value.id}/edit`)
}

const shareLibrary = async () => {
  try {
    if (navigator.share) {
      await navigator.share({
        title: library.value.name,
        text: library.value.description,
        url: window.location.href
      })
      success('Library Shared', 'Library has been shared successfully.')
    } else {
      await navigator.clipboard.writeText(window.location.href)
      success('Link Copied', 'Library link has been copied to clipboard.')
    }
  } catch (err) {
    if (err.name === 'AbortError') {
      // User cancelled share - don't show error
      return
    }
    showError('Share Failed', 'Unable to share library. Please try copying the link manually.')
  }
}

const showArticleMenu = (article) => {
  // Show context menu or navigate to article options
  console.log('Show menu for article:', article.id)
}

// Meta
useHead({
  title: () => library.value ? `${library.value.name} - Academic Reader` : 'Loading Library...',
  meta: [
    { name: 'description', content: () => library.value?.description || 'Academic library details' }
  ]
})
</script>