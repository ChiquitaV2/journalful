<template>
  <div class="glass-card p-4 hover:scale-[1.02] transition-all duration-300">
    <!-- Article Header -->
    <div class="flex items-start justify-between mb-3">
      <div class="flex-1 min-w-0">
        <h4 class="font-semibold text-white text-sm md:text-base line-clamp-2 mb-1">
          {{ article.articleTitle }}
        </h4>
        <p class="text-gray-400 text-xs md:text-sm mb-2">
          {{ formatAuthors(article.authors) }}
        </p>
        <div class="flex items-center space-x-2 text-xs text-gray-500">
          <span v-if="article.publicationYear">{{ article.publicationYear }}</span>
          <span v-if="article.publicationYear && article.journalName">•</span>
          <span v-if="article.journalName" class="truncate">{{ article.journalName }}</span>
        </div>
      </div>
      
      <!-- Status Badge -->
      <div class="ml-3 flex-shrink-0">
        <StatusBadge :status="article.readingStatus" size="sm" />
      </div>
    </div>

    <!-- DOI -->
    <div class="mb-3">
      <a
        :href="`https://doi.org/${article.doi}`"
        target="_blank"
        rel="noopener noreferrer"
        class="text-purple-300 hover:text-purple-200 text-xs font-mono bg-white/5 px-2 py-1 rounded"
      >
        {{ article.doi }}
      </a>
    </div>

    <!-- Notes Preview -->
    <div v-if="article.notes" class="mb-3">
      <p class="text-gray-300 text-sm line-clamp-2 bg-white/5 p-2 rounded">
        {{ article.notes }}
      </p>
    </div>

    <!-- Actions -->
    <div class="flex items-center justify-between">
      <div class="flex items-center space-x-2">
        <button
          @click="$emit('view-details', article)"
          class="glass-button !py-1 !px-3 text-xs"
        >
          <Icon name="heroicons:eye" class="h-3 w-3 mr-1" />
          Details
        </button>
        <button
          @click="showStatusSelector = !showStatusSelector"
          class="glass-button !py-1 !px-3 text-xs"
        >
          <Icon name="heroicons:arrow-path" class="h-3 w-3 mr-1" />
          Status
        </button>
      </div>
      
      <div class="text-xs text-gray-500">
        Added {{ formatDate(article.addedAt) }}
      </div>
    </div>

    <!-- Status Selector (Mobile Bottom Sheet) -->
    <div
      v-if="showStatusSelector"
      class="fixed inset-0 z-50 glass-overlay"
      @click="showStatusSelector = false"
    >
      <div
        class="glass-sheet fixed bottom-0 left-0 right-0 p-4 animate-slide-up"
        @click.stop
      >
        <div class="w-12 h-1 bg-white/30 rounded-full mx-auto mb-4"></div>
        <h3 class="text-lg font-semibold text-white mb-4">Update Reading Status</h3>
        
        <div class="space-y-2">
          <button
            v-for="status in statusOptions"
            :key="status.value"
            @click="updateStatus(status.value)"
            :class="[
              'w-full p-3 rounded-lg text-left transition-all',
              article.readingStatus === status.value
                ? 'bg-purple-500/30 border border-purple-300/50'
                : 'glass-subtle hover:bg-white/10'
            ]"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center space-x-3">
                <Icon :name="status.icon" class="h-5 w-5" :class="status.color" />
                <span class="text-white font-medium">{{ status.label }}</span>
              </div>
              <Icon
                v-if="article.readingStatus === status.value"
                name="heroicons:check"
                class="h-5 w-5 text-green-400"
              />
            </div>
          </button>
        </div>
        
        <button
          @click="showStatusSelector = false"
          class="w-full mt-4 glass-button"
        >
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ReadingStatus } from '~~/types'

const props = defineProps({
  article: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['status-change', 'view-details'])

const showStatusSelector = ref(false)
const { showSuccess } = useNotifications()

const statusOptions = [
  {
    value: ReadingStatus.READING_STATUS_TO_READ,
    label: 'To Read',
    icon: 'heroicons:bookmark',
    color: 'text-blue-400'
  },
  {
    value: ReadingStatus.READING_STATUS_READING,
    label: 'Currently Reading',
    icon: 'heroicons:book-open',
    color: 'text-orange-400'
  },
  {
    value: ReadingStatus.READING_STATUS_READ,
    label: 'Completed',
    icon: 'heroicons:check-circle',
    color: 'text-green-400'
  },
  {
    value: ReadingStatus.READING_STATUS_ABANDONED,
    label: 'Abandoned',
    icon: 'heroicons:x-circle',
    color: 'text-gray-400'
  }
]

const formatAuthors = (authors) => {
  if (!authors || authors.length === 0) return 'Unknown Author'
  if (authors.length === 1) return authors[0].name
  if (authors.length === 2) return `${authors[0].name} & ${authors[1].name}`
  return `${authors[0].name} et al.`
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric'
  })
}

const updateStatus = (newStatus) => {
  emit('status-change', props.article.id, newStatus)
  showStatusSelector.value = false
  
  const statusLabel = statusOptions.find(s => s.value === newStatus)?.label
  showSuccess(`Status updated to "${statusLabel}"`)
}
</script>
