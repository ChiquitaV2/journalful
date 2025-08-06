<template>
  <div class="glass-card p-4 hover:scale-[1.02] transition-all duration-300">
    <!-- Library Header -->
    <div class="flex items-start justify-between mb-3">
      <div class="flex-1 min-w-0">
        <div class="flex items-center space-x-2 mb-1">
          <h3 class="font-semibold text-white text-base truncate">
            {{ library.name || 'My Reading List' }}
          </h3>
          <div v-if="isDefault" class="flex-shrink-0">
            <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-500/20 text-green-300 border border-green-300/30">
              <Icon name="heroicons:globe-alt" class="h-3 w-3 mr-1" />
              Public
            </span>
          </div>
          <div v-else class="flex-shrink-0">
            <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-blue-500/20 text-blue-300 border border-blue-300/30">
              <Icon name="heroicons:lock-closed" class="h-3 w-3 mr-1" />
              Private
            </span>
          </div>
        </div>
        
        <p class="text-gray-400 text-sm">
          {{ library.articles.length }} {{ library.articles.length === 1 ? 'article' : 'articles' }}
        </p>
        
        <div class="text-xs text-gray-500 mt-1">
          Updated {{ formatDate(library.updatedAt) }}
        </div>
      </div>
      
      <!-- Action Menu -->
      <div class="ml-3 flex-shrink-0 relative">
        <button
          @click="showMenu = !showMenu"
          class="glass-button !p-2"
        >
          <Icon name="heroicons:ellipsis-vertical" class="h-4 w-4" />
        </button>
        
        <!-- Dropdown Menu -->
        <div
          v-if="showMenu"
          class="absolute right-0 top-full mt-1 z-10 glass-elevated rounded-lg py-1 w-48"
          @click.stop
        >
          <button
            @click="handleView"
            class="w-full px-4 py-2 text-left text-sm text-white hover:bg-white/10 flex items-center space-x-2"
          >
            <Icon name="heroicons:eye" class="h-4 w-4" />
            <span>View Library</span>
          </button>
          
          <button
            @click="handleEdit"
            class="w-full px-4 py-2 text-left text-sm text-white hover:bg-white/10 flex items-center space-x-2"
          >
            <Icon name="heroicons:pencil" class="h-4 w-4" />
            <span>{{ isDefault ? 'Rename' : 'Edit' }}</span>
          </button>
          
          <div v-if="!isDefault" class="border-t border-white/10 my-1"></div>
          
          <button
            v-if="!isDefault"
            @click="handleDelete"
            class="w-full px-4 py-2 text-left text-sm text-red-400 hover:bg-red-500/10 flex items-center space-x-2"
          >
            <Icon name="heroicons:trash" class="h-4 w-4" />
            <span>Delete</span>
          </button>
          
          <div v-if="!isDefault" class="border-t border-white/10 my-1"></div>
          
          <button
            v-if="!isDefault"
            @click="generateShareLink"
            class="w-full px-4 py-2 text-left text-sm text-white hover:bg-white/10 flex items-center space-x-2"
          >
            <Icon name="heroicons:share" class="h-4 w-4" />
            <span>Share Link</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Reading Status Overview -->
    <div v-if="library.articles.length > 0" class="mb-3">
      <div class="flex items-center space-x-2 mb-2">
        <h4 class="text-sm font-medium text-gray-300">Reading Progress</h4>
        <div class="flex-1 h-1 bg-gray-700 rounded-full overflow-hidden">
          <div 
            class="h-full bg-gradient-to-r from-emerald-500 to-green-500"
            :style="{ width: `${progressPercentage}%` }"
          ></div>
        </div>
        <span class="text-xs text-gray-400">{{ progressPercentage }}%</span>
      </div>
      
      <div class="grid grid-cols-4 gap-2 text-center">
        <div class="text-xs">
          <div class="font-medium text-blue-300">{{ statusCounts.toRead }}</div>
          <div class="text-gray-500">To Read</div>
        </div>
        <div class="text-xs">
          <div class="font-medium text-orange-300">{{ statusCounts.reading }}</div>
          <div class="text-gray-500">Reading</div>
        </div>
        <div class="text-xs">
          <div class="font-medium text-emerald-300">{{ statusCounts.read }}</div>
          <div class="text-gray-500">Read</div>
        </div>
        <div class="text-xs">
          <div class="font-medium text-gray-400">{{ statusCounts.abandoned }}</div>
          <div class="text-gray-500">Abandoned</div>
        </div>
      </div>
    </div>

    <!-- Recent Articles Preview -->
    <div v-if="library.articles.length > 0" class="mb-3">
      <h4 class="text-sm font-medium text-gray-300 mb-2">Recent Articles</h4>
      <div class="space-y-2">
        <div
          v-for="article in recentArticles"
          :key="article.id"
          class="flex items-center space-x-2 p-2 glass-subtle rounded-lg"
        >
          <StatusBadge :status="article.readingStatus" size="xs" />
          <div class="flex-1 min-w-0">
            <p class="text-sm text-white truncate">{{ article.articleTitle }}</p>
            <p class="text-xs text-gray-400">{{ article.publicationYear }}</p>
          </div>
        </div>
      </div>
      
      <div v-if="library.articles.length > 3" class="text-center mt-2">
        <button
          @click="handleView"
          class="text-xs text-purple-300 hover:text-purple-200"
        >
          View all {{ library.articles.length }} articles →
        </button>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="text-center py-4">
      <Icon name="heroicons:book-open" class="h-8 w-8 text-gray-500 mx-auto mb-2" />
      <p class="text-sm text-gray-400 mb-2">No articles yet</p>
      <button
        @click="handleView"
        class="glass-button !py-1 !px-3 text-xs bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30"
      >
        Add Articles
      </button>
    </div>

    <!-- Action Buttons -->
    <div class="flex space-x-2 pt-3 border-t border-white/10">
      <button
        @click="handleView"
        class="flex-1 glass-button !py-2 text-sm"
      >
        <Icon name="heroicons:eye" class="h-4 w-4 mr-2" />
        View
      </button>
      
      <button
        v-if="library.articles.length > 0"
        @click="exportLibrary"
        class="glass-button !py-2 !px-3"
      >
        <Icon name="heroicons:arrow-down-tray" class="h-4 w-4" />
      </button>
    </div>

    <!-- Click outside to close menu -->
    <div
      v-if="showMenu"
      class="fixed inset-0 z-0"
      @click="showMenu = false"
    ></div>
  </div>
</template>

<script setup>
import { ReadingStatus } from '~~/types'

const props = defineProps({
  library: {
    type: Object,
    required: true
  },
  isDefault: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['view', 'edit', 'delete'])

const showMenu = ref(false)
const { success, info } = useToast()

// Computed properties
const statusCounts = computed(() => {
  const counts = {
    toRead: 0,
    reading: 0,
    read: 0,
    abandoned: 0
  }

  props.library.articles.forEach(article => {
    switch (article.readingStatus) {
      case ReadingStatus.READING_STATUS_TO_READ:
        counts.toRead++
        break
      case ReadingStatus.READING_STATUS_READING:
        counts.reading++
        break
      case ReadingStatus.READING_STATUS_READ:
        counts.read++
        break
      case ReadingStatus.READING_STATUS_ABANDONED:
        counts.abandoned++
        break
    }
  })

  return counts
})

const progressPercentage = computed(() => {
  if (props.library.articles.length === 0) return 0
  return Math.round((statusCounts.value.read / props.library.articles.length) * 100)
})

const recentArticles = computed(() => {
  return props.library.articles
    .sort((a, b) => new Date(b.addedAt || 0) - new Date(a.addedAt || 0))
    .slice(0, 3)
})

// Methods
const handleView = () => {
  showMenu.value = false
  emit('view', props.library)
}

const handleEdit = () => {
  showMenu.value = false
  emit('edit', props.library)
}

const handleDelete = () => {
  showMenu.value = false
  emit('delete', props.library)
}

const generateShareLink = async () => {
  showMenu.value = false

  try {
    // Mock share link generation
    const shareLink = `${window.location.origin}/libraries/shared/${props.library.id}`

    if (navigator.share) {
      await navigator.share({
        title: props.library.name,
        text: `Check out my reading library: ${props.library.name}`,
        url: shareLink
      })
    } else if (navigator.clipboard) {
      await navigator.clipboard.writeText(shareLink)
      success('Share Link Copied', 'The share link has been copied to your clipboard.')
    } else {
      info('Share Link', `Share link: ${shareLink}`)
    }
  } catch (error) {
    console.error('Failed to share:', error)
  }
}

const exportLibrary = async () => {
  try {
    // Mock export functionality
    info('Coming Soon', 'Export feature will be available in the next update!')
  } catch (error) {
    console.error('Failed to export:', error)
  }
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric'
  })
}

// Close menu when clicking outside
onMounted(() => {
  document.addEventListener('click', () => {
    showMenu.value = false
  })
})

onUnmounted(() => {
  document.removeEventListener('click', () => {
    showMenu.value = false
  })
})
</script>
