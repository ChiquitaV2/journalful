<template>
  <div class="min-h-screen">
    <!-- Minimal header for focused reading -->
    <header class="glass-nav fixed top-0 left-0 right-0 z-40 h-14">
      <div class="flex items-center justify-between h-full mobile-padding">
        <button 
          @click="goBack"
          class="glass-button !p-2"
        >
          <Icon name="heroicons:arrow-left" class="h-5 w-5" />
        </button>
        
        <div class="flex items-center gap-2">
          <!-- Reading Progress Indicator -->
          <div 
            v-if="showProgress"
            class="flex items-center gap-2 text-sm text-gray-300"
          >
            <span>{{ progressText }}</span>
            <div class="w-16 h-1 bg-gray-600 rounded-full overflow-hidden">
              <div 
                class="h-full bg-gradient-to-r from-purple-500 to-blue-500 transition-all duration-300"
                :style="{ width: `${progress}%` }"
              ></div>
            </div>
          </div>
          
          <!-- Action Buttons -->
          <button 
            v-if="showBookmark"
            @click="toggleBookmark"
            class="glass-button !p-2"
            :class="isBookmarked ? 'text-yellow-400' : 'text-gray-400'"
          >
            <Icon :name="isBookmarked ? 'heroicons:bookmark-solid' : 'heroicons:bookmark'" class="h-5 w-5" />
          </button>
          
          <button 
            v-if="showShare"
            @click="handleShare"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:share" class="h-5 w-5" />
          </button>
          
          <button 
            v-if="showSettings"
            @click="handleSettings"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:cog-6-tooth" class="h-5 w-5" />
          </button>
        </div>
      </div>
    </header>

    <!-- Main Content (minimal padding for reading) -->
    <main class="pt-14">
      <slot />
    </main>
  </div>
</template>

<script setup>
const router = useRouter()

const props = defineProps({
  backRoute: {
    type: String,
    default: null
  },
  showProgress: {
    type: Boolean,
    default: false
  },
  progress: {
    type: Number,
    default: 0
  },
  progressText: {
    type: String,
    default: ''
  },
  showBookmark: {
    type: Boolean,
    default: false
  },
  isBookmarked: {
    type: Boolean,
    default: false
  },
  showShare: {
    type: Boolean,
    default: false
  },
  showSettings: {
    type: Boolean,
    default: false
  }
})

// Emit events for parent pages to handle
const emit = defineEmits(['bookmark', 'share', 'settings'])

const goBack = () => {
  if (props.backRoute) {
    router.push(props.backRoute)
  } else {
    router.back()
  }
}

const toggleBookmark = () => {
  emit('bookmark')
}

const handleShare = () => {
  emit('share')
}

const handleSettings = () => {
  emit('settings')
}
</script>
