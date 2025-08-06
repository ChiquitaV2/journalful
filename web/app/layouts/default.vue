<template>
  <div class="min-h-screen">
    <!-- Mobile Header -->
    <header class="glass-nav fixed top-0 left-0 right-0 z-40 h-16">
      <div class="flex items-center justify-between h-full mobile-padding">
        <div class="flex items-center space-x-3">
          <!-- Back Button for non-home pages -->
          <NuxtLink 
            v-if="showBackButton" 
            :to="backRoute" 
            class="glass-button !p-2"
          >
            <Icon name="heroicons:arrow-left" class="h-6 w-6" />
          </NuxtLink>
          
          <!-- Mobile Menu Button for home page -->
          <button
            v-if="isHomePage"
            @click="toggleMobileMenu"
            class="glass-button p-2 block"
          >
            <Icon name="heroicons:bars-3" class="h-6 w-6" />
          </button>
          
          <h1 class="text-xl font-bold text-white">{{ pageTitle }}</h1>
        </div>
        
        <div class="flex items-center gap-2">
          <!-- Search Button -->
          <button 
            v-if="showSearchButton"
            @click="toggleSearch"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:magnifying-glass" class="h-5 w-5" />
          </button>
          
          <!-- Add Button -->
          <button 
            v-if="showAddButton"
            @click="handleAddAction"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:plus" class="h-5 w-5" />
          </button>
          
          <!-- Additional Action Button -->
          <button 
            v-if="showActionButton"
            @click="handleAction"
            class="glass-button !p-2"
          >
            <Icon :name="actionIcon" class="h-5 w-5" />
          </button>
          
          <!-- Profile Avatar -->
          <div 
            v-if="showProfile"
            class="w-8 h-8 rounded-full bg-gradient-to-r from-purple-500 to-blue-500 flex items-center justify-center cursor-pointer"
            @click="navigateToProfile"
          >
            <Icon name="heroicons:user" class="h-5 w-5 text-white" />
          </div>
        </div>
      </div>
    </header>

    <!-- Mobile Menu Overlay (only on home page) -->
    <div
      v-if="showMobileMenu && isHomePage"
      class="fixed inset-0 z-30 bg-black/50 backdrop-blur-sm"
      @click="toggleMobileMenu"
    >
      <div class="glass-elevated animate-slide-up w-80 h-full p-6">
        <nav class="mt-16">
          <div class="flex flex-col gap-4">
            <NuxtLink 
              to="/" 
              class="glass-card flex items-center gap-3 p-4 text-white no-underline"
              @click="closeMobileMenu"
            >
              <Icon name="heroicons:home" class="h-5 w-5" />
              <span>Dashboard</span>
            </NuxtLink>
            <NuxtLink 
              to="/libraries" 
              class="glass-card flex items-center gap-3 p-4 text-white no-underline"
              @click="closeMobileMenu"
            >
              <Icon name="heroicons:book-open" class="h-5 w-5" />
              <span>My Libraries</span>
            </NuxtLink>
            <NuxtLink 
              to="/search" 
              class="glass-card flex items-center gap-3 p-4 text-white no-underline"
              @click="closeMobileMenu"
            >
              <Icon name="heroicons:magnifying-glass" class="h-5 w-5" />
              <span>Search</span>
            </NuxtLink>
            <NuxtLink 
              to="/profile" 
              class="glass-card flex items-center gap-3 p-4 text-white no-underline"
              @click="closeMobileMenu"
            >
              <Icon name="heroicons:user-circle" class="h-5 w-5" />
              <span>Profile</span>
            </NuxtLink>
            <NuxtLink 
              to="/stats" 
              class="glass-card flex items-center gap-3 p-4 text-white no-underline"
              @click="closeMobileMenu"
            >
              <Icon name="heroicons:chart-bar" class="h-5 w-5" />
              <span>Statistics</span>
            </NuxtLink>
            <NuxtLink 
              to="/settings" 
              class="glass-card flex items-center gap-3 p-4 text-white no-underline"
              @click="closeMobileMenu"
            >
              <Icon name="heroicons:cog-6-tooth" class="h-5 w-5" />
              <span>Settings</span>
            </NuxtLink>
          </div>
        </nav>
      </div>
    </div>

    <!-- Main Content -->
    <main class="pt-16 pb-20">
      <div class="mobile-padding max-w-7xl mx-auto">
        <slot />
      </div>
    </main>

    <SearchModal
        v-if="showSearch"
        @close="showSearch = false"
        @article-selected="handleArticleSelected"
    />
    <!-- Mobile Bottom Navigation -->
    <MobileBottomNav />
  </div>
</template>

<script setup>
const route = useRoute()
const router = useRouter()

// Layout configuration props (provided by pages)
const props = defineProps({
  pageTitle: {
    type: String,
    default: 'Academic Reader'
  },
  backRoute: {
    type: String,
    default: '/'
  },
  showSearchButton: {
    type: Boolean,
    default: true
  },
  showAddButton: {
    type: Boolean,
    default: false
  },
  showActionButton: {
    type: Boolean,
    default: false
  },
  actionIcon: {
    type: String,
    default: 'heroicons:ellipsis-horizontal'
  },
  showProfile: {
    type: Boolean,
    default: true
  }
})

// Mobile menu state
const showMobileMenu = ref(false)

// Computed properties
const isHomePage = computed(() => route.path === '/')
const showBackButton = computed(() => !isHomePage.value)

// Mobile menu methods
const toggleMobileMenu = () => {
  showMobileMenu.value = !showMobileMenu.value
}

const closeMobileMenu = () => {
  showMobileMenu.value = false
}
const showSearch = ref(false)
const toggleSearch = () => {
  showSearch.value = !showSearch.value
}

const handleArticleSelected = (article) => {
  showSearch.value = false
  // Handle article selection
  navigateTo(`/articles/${article.id}`)
}

// Navigation methods
const navigateToProfile = () => {
  navigateTo('/profile')
  // router.push('/profile')
}

// Handle events from pages
const emit = defineEmits(['add-action', 'action'])

const handleAddAction = () => {
  // Try to emit to parent page first, fallback to default behavior
  emit('add-action')
}

const handleAction = () => {
  // Try to emit to parent page first, fallback to default behavior
  emit('action')
}

// Close mobile menu on route change
watch(() => route.path, () => {
  closeMobileMenu()
})

// Close mobile menu on escape key
onMounted(() => {
  const handleEscape = (e) => {
    if (e.key === 'Escape') {
      closeMobileMenu()
    }
  }
  document.addEventListener('keydown', handleEscape)
  
  onUnmounted(() => {
    document.removeEventListener('keydown', handleEscape)
  })
})
</script>

<style scoped>
.animate-slide-up {
  animation: slideUp 0.3s ease-out;
}

@keyframes slideUp {
  from {
    transform: translateX(-100%);
  }
  to {
    transform: translateX(0);
  }
}
</style>
