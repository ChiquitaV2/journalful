<template>
  <div class="bg-white border border-gray-200 rounded-lg shadow-sm hover:shadow-md transition-shadow p-6">
    <!-- Article Header -->
    <div class="flex items-start justify-between mb-4">
      <div class="flex-1">
        <h3 class="text-lg font-semibold text-gray-900 mb-2 line-clamp-2">
          <NuxtLink 
            :to="`/articles/${article.id}`" 
            class="hover:text-primary transition-colors"
          >
            {{ article.title }}
          </NuxtLink>
        </h3>
        
        <!-- Authors -->
        <div class="flex flex-wrap items-center gap-1 mb-2">
          <span
            v-for="(author, index) in article.authors"
            :key="author.id"
            class="text-sm text-gray-600"
          >
            <NuxtLink 
              :to="`/profile/${author.id}`" 
              class="hover:text-primary transition-colors"
            >
              {{ author.name }}
            </NuxtLink>
            <span v-if="index < article.authors.length - 1">, </span>
          </span>
        </div>
        
        <!-- Publication Info -->
        <div class="flex items-center gap-4 text-sm text-gray-500 mb-3">
          <span v-if="article.journalName">
            <Icon name="heroicons:document-text" class="h-4 w-4 inline mr-1" />
            {{ article.journalName }}
          </span>
          <span v-if="article.publicationYear">
            <Icon name="heroicons:calendar" class="h-4 w-4 inline mr-1" />
            {{ article.publicationYear }}
          </span>
          <span v-if="article.doi">
            <Icon name="heroicons:link" class="h-4 w-4 inline mr-1" />
            DOI: {{ article.doi }}
          </span>
        </div>
      </div>
      
      <!-- Actions -->
      <div class="ml-4 flex flex-col gap-2">
        <Button
          size="sm"
          variant="outline"
          @click="toggleLibrary"
          :loading="isAddingToLibrary"
        >
          <Icon 
            :name="isInLibrary ? 'heroicons:bookmark-solid' : 'heroicons:bookmark'" 
            class="h-4 w-4"
          />
        </Button>
        
        <div class="relative" ref="menuRef">
          <Button
            size="sm"
            variant="ghost"
            @click="showMenu = !showMenu"
          >
            <Icon name="heroicons:ellipsis-vertical" class="h-4 w-4" />
          </Button>
          
          <!-- Dropdown Menu -->
          <Transition
            enter-active-class="transition ease-out duration-100"
            enter-from-class="transform opacity-0 scale-95"
            enter-to-class="transform opacity-100 scale-100"
            leave-active-class="transition ease-in duration-75"
            leave-from-class="transform opacity-100 scale-100"
            leave-to-class="transform opacity-0 scale-95"
          >
            <div
              v-if="showMenu"
              class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg border border-gray-200 py-1 z-10"
            >
              <button
                @click="copyDOI"
                class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              >
                <Icon name="heroicons:clipboard" class="h-4 w-4 mr-2 inline" />
                Copy DOI
              </button>
              <button
                @click="copyLink"
                class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              >
                <Icon name="heroicons:link" class="h-4 w-4 mr-2 inline" />
                Copy Link
              </button>
              <button
                @click="exportCitation"
                class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
              >
                <Icon name="heroicons:document-arrow-down" class="h-4 w-4 mr-2 inline" />
                Export Citation
              </button>
            </div>
          </Transition>
        </div>
      </div>
    </div>
    
    <!-- Abstract Preview -->
    <div v-if="article.abstract" class="mb-4">
      <p class="text-sm text-gray-600 line-clamp-3">
        {{ article.abstract }}
      </p>
    </div>
    
    <!-- Reading Status -->
    <div v-if="readingStatus" class="mb-4">
      <Tag
        :variant="getReadingStatusVariant(readingStatus)"
        :icon="getReadingStatusIcon(readingStatus)"
        size="sm"
      >
        {{ formatReadingStatus(readingStatus) }}
      </Tag>
    </div>
    
    <!-- Footer -->
    <div class="flex items-center justify-between pt-4 border-t border-gray-100">
      <div class="flex items-center gap-2 text-xs text-gray-500">
        <span>Added {{ formatDate(article.createdAt) }}</span>
        <span v-if="article.updatedAt !== article.createdAt">
          • Updated {{ formatDate(article.updatedAt) }}
        </span>
      </div>
      
      <Button
        size="sm"
        variant="ghost"
        @click="navigateTo(`/articles/${article.id}`)"
      >
        Read More
        <Icon name="heroicons:arrow-right" class="h-4 w-4 ml-1" />
      </Button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Article } from '~~/shared/grpc/articles/v1/article'
import { ReadingStatus } from '~~/shared/grpc/library/v1/library'

interface Props {
  article: Article
  readingStatus?: ReadingStatus
  showLibraryActions?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  showLibraryActions: true
})

const emit = defineEmits<{
  'library-updated': [articleId: number, isInLibrary: boolean]
}>()

const showMenu = ref(false)
const menuRef = ref<HTMLElement>()
const isAddingToLibrary = ref(false)
const isInLibrary = ref(false)

const toggleLibrary = async () => {
  isAddingToLibrary.value = true
  try {
    // TODO: Implement library toggle logic
    isInLibrary.value = !isInLibrary.value
    emit('library-updated', props.article.id, isInLibrary.value)
  } catch (error) {
    console.error('Error toggling library:', error)
  } finally {
    isAddingToLibrary.value = false
  }
}

const copyDOI = async () => {
  if (props.article.doi) {
    await navigator.clipboard.writeText(props.article.doi)
    // TODO: Show toast notification
  }
  showMenu.value = false
}

const copyLink = async () => {
  const url = `${window.location.origin}/articles/${props.article.id}`
  await navigator.clipboard.writeText(url)
  // TODO: Show toast notification
  showMenu.value = false
}

const exportCitation = () => {
  // TODO: Implement citation export
  showMenu.value = false
}

const getReadingStatusVariant = (status: ReadingStatus) => {
  switch (status) {
    case ReadingStatus.READING_STATUS_READ:
      return 'success'
    case ReadingStatus.READING_STATUS_READING:
      return 'primary'
    case ReadingStatus.READING_STATUS_TO_READ:
      return 'warning'
    case ReadingStatus.READING_STATUS_ABANDONED:
      return 'error'
    default:
      return 'default'
  }
}

const getReadingStatusIcon = (status: ReadingStatus) => {
  switch (status) {
    case ReadingStatus.READING_STATUS_READ:
      return 'heroicons:check-circle'
    case ReadingStatus.READING_STATUS_READING:
      return 'heroicons:eye'
    case ReadingStatus.READING_STATUS_TO_READ:
      return 'heroicons:clock'
    case ReadingStatus.READING_STATUS_ABANDONED:
      return 'heroicons:x-circle'
    default:
      return 'heroicons:document'
  }
}

const formatReadingStatus = (status: ReadingStatus) => {
  switch (status) {
    case ReadingStatus.READING_STATUS_READ:
      return 'Read'
    case ReadingStatus.READING_STATUS_READING:
      return 'Reading'
    case ReadingStatus.READING_STATUS_TO_READ:
      return 'To Read'
    case ReadingStatus.READING_STATUS_ABANDONED:
      return 'Abandoned'
    default:
      return 'Unknown'
  }
}

const formatDate = (date: Date | undefined) => {
  if (!date) return 'Unknown'
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(new Date(date))
}

// Close menu when clicking outside
onMounted(() => {
  const handleClickOutside = (event: MouseEvent) => {
    if (menuRef.value && !menuRef.value.contains(event.target as Node)) {
      showMenu.value = false
    }
  }
  document.addEventListener('click', handleClickOutside)
  
  onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside)
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

.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
