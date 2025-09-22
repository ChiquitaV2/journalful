<template>
  <div>
    <!-- Profile Header -->
    <section class="mb-6">
      <div class="glass-card p-6">
        <div class="text-center mb-6">
          <!-- Avatar -->
          <div class="relative inline-block mb-4">
            <div class="w-24 h-24 md:w-32 md:h-32 rounded-full bg-gradient-to-r from-purple-500 to-blue-500 flex items-center justify-center text-2xl md:text-3xl font-bold text-white">
              {{ avatarInitials }}
            </div>
            <button
              v-if="isEditing"
              class="absolute bottom-0 right-0 glass-button !p-2 rounded-full"
            >
              <Icon name="heroicons:camera" class="h-4 w-4" />
            </button>
          </div>

          <!-- Name and Bio -->
          <div v-if="!isEditing">
            <h2 class="text-2xl font-bold text-white mb-2">{{ profile.name }}</h2>
            <p class="text-gray-400 mb-4">{{ profile.bio || 'No bio added yet.' }}</p>
            <div class="flex flex-col sm:flex-row items-center justify-center space-y-2 sm:space-y-0 sm:space-x-4 text-sm text-gray-400">
              <div class="flex items-center space-x-1">
                <Icon name="heroicons:building-office" class="h-4 w-4" />
                <span>{{ profile.institution || 'No institution' }}</span>
              </div>
                  <div class="flex items-center space-x-1">
                    <Icon name="heroicons:calendar" class="h-4 w-4" />
                    <span>Joined {{ formatDate(profile.createdAt) }}</span>
                  </div>
                </div>
              </div>

              <!-- Edit Form -->
              <form v-else @submit.prevent="saveProfile" class="text-left max-w-md mx-auto">
                <div class="space-y-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-300 mb-2">
                      Full Name *
                    </label>
                    <input
                      v-model="editForm.name"
                      type="text"
                      required
                      class="glass-input"
                      placeholder="Your full name"
                    />
                  </div>

                  <div>
                    <label class="block text-sm font-medium text-gray-300 mb-2">
                      Bio
                    </label>
                    <textarea
                      v-model="editForm.bio"
                      rows="3"
                      class="glass-input resize-none"
                      placeholder="Tell us about yourself..."
                      maxlength="500"
                    ></textarea>
                    <div class="text-xs text-gray-400 mt-1">
                      {{ editForm.bio.length }}/500 characters
                    </div>
                  </div>

                  <div>
                    <label class="block text-sm font-medium text-gray-300 mb-2">
                      Institution
                    </label>
                    <input
                      v-model="editForm.institution"
                      type="text"
                      class="glass-input"
                      placeholder="Your university or organization"
                    />
                  </div>
                </div>

                <div class="flex space-x-3 mt-6">
                  <button
                    type="button"
                    @click="cancelEdit"
                    class="flex-1 glass-button"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    :disabled="isSaving"
                    class="flex-1 glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30 disabled:opacity-50"
                  >
                    <Icon
                      v-if="isSaving"
                      name="heroicons:arrow-path"
                      class="h-4 w-4 mr-2 animate-spin"
                    />
                    Save Changes
                  </button>
                </div>
              </form>
            </div>
          </div>
        </section>

        <!-- Reading Statistics -->
        <section class="mb-6">
          <h3 class="text-lg font-semibold text-white mb-4">Reading Statistics</h3>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div class="glass-card text-center p-4">
              <div class="text-2xl font-bold text-blue-300">{{ stats.toRead }}</div>
              <div class="text-sm text-gray-400">To Read</div>
            </div>
            <div class="glass-card text-center p-4">
              <div class="text-2xl font-bold text-orange-300">{{ stats.reading }}</div>
              <div class="text-sm text-gray-400">Reading</div>
            </div>
            <div class="glass-card text-center p-4">
              <div class="text-2xl font-bold text-emerald-300">{{ stats.completed }}</div>
              <div class="text-sm text-gray-400">Completed</div>
            </div>
            <div class="glass-card text-center p-4">
              <div class="text-2xl font-bold text-white">{{ stats.total }}</div>
              <div class="text-sm text-gray-400">Total Articles</div>
            </div>
          </div>
        </section>

        <!-- Recent Activity -->
        <section class="mb-6">
          <h3 class="text-lg font-semibold text-white mb-4">Recent Activity</h3>
          <div class="space-y-3">
            <div
              v-for="activity in recentActivity"
              :key="activity.id"
              class="glass-card p-4"
            >
              <div class="flex items-start space-x-3">
                <div class="flex-shrink-0">
                  <Icon
                    :name="activity.icon"
                    :class="activity.iconColor"
                    class="h-5 w-5 mt-0.5"
                  />
                </div>
                <div class="flex-1 min-w-0">
                  <p class="text-white text-sm">{{ activity.description }}</p>
                  <p class="text-gray-400 text-xs mt-1">{{ formatRelativeTime(activity.timestamp) }}</p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Libraries Overview -->
        <section class="mb-6">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold text-white">Your Libraries</h3>
            <NuxtLink to="/libraries" class="text-purple-300 hover:text-purple-200 text-sm">
              View all →
            </NuxtLink>
          </div>
          
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="glass-card p-4">
              <div class="flex items-center justify-between mb-2">
                <h4 class="font-medium text-white">Default Library</h4>
                <div class="flex items-center space-x-1">
                  <Icon name="heroicons:globe-alt" class="h-4 w-4 text-green-400" />
                  <span class="text-xs text-green-400">Public</span>
                </div>
              </div>
              <p class="text-sm text-gray-400">{{ stats.defaultLibraryCount }} articles</p>
            </div>
            
            <div class="glass-card p-4">
              <div class="flex items-center justify-between mb-2">
                <h4 class="font-medium text-white">Private Libraries</h4>
                <div class="flex items-center space-x-1">
                  <Icon name="heroicons:lock-closed" class="h-4 w-4 text-blue-400" />
                  <span class="text-xs text-blue-400">{{ stats.privateLibraryCount }}</span>
                </div>
              </div>
              <p class="text-sm text-gray-400">{{ stats.privateArticleCount }} articles total</p>
            </div>
          </div>
        </section>

        <!-- Settings -->
        <section>
          <h3 class="text-lg font-semibold text-white mb-4">Settings</h3>
          <div class="space-y-3">
            <div class="glass-card p-4">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-white">Export Data</h4>
                  <p class="text-sm text-gray-400">Download your reading list and notes</p>
                </div>
                <button class="glass-button !py-2 !px-4">
                  <Icon name="heroicons:arrow-down-tray" class="h-4 w-4 mr-2" />
                  Export
                </button>
              </div>
            </div>
            
            <div class="glass-card p-4">
              <div class="flex items-center justify-between">
                <div>
                  <h4 class="font-medium text-white">Privacy Settings</h4>
                  <p class="text-sm text-gray-400">Manage who can see your profile</p>
                </div>
                <button class="glass-button !py-2 !px-4">
                  <Icon name="heroicons:cog-6-tooth" class="h-4 w-4 mr-2" />
                  Manage
                </button>
              </div>
            </div>
          </div>
        </section>
      </div>
</template>

<script setup>


// Configure layout
const { setLayoutConfig } = useLayout()

onMounted(() => {
  setLayoutConfig({
    pageTitle: 'Profile',
    showActionButton: true,
    actionIcon: 'heroicons:pencil',
    actionCallback: toggleEditMode
  })
})
const { showSuccess, showError } = useNotifications()

// State
const isEditing = ref(false)
const isSaving = ref(false)
const isLoading = ref(true)

// Profile data from API
const profile = ref(null)
const stats = ref({
  total: 0,
  toRead: 0,
  reading: 0,
  completed: 0,
  abandoned: 0,
  defaultLibraryCount: 0,
  privateLibraryCount: 0,
  privateArticleCount: 0
})

const editForm = ref({
  name: '',
  bio: '',
  institution: ''
})

// Load profile data
const loadProfile = async () => {
  try {
    const profileData = await $fetch('/api/profile/me')
    profile.value = profileData
    
    // Load profile statistics
    const librariesData = await $fetch('/api/libraries')
    const articlesData = await $fetch('/api/articles')
    
    stats.value = {
      total: articlesData.length || 0,
      toRead: articlesData.filter(a => a.readingStatus === 'READING_STATUS_TO_READ').length || 0,
      reading: articlesData.filter(a => a.readingStatus === 'READING_STATUS_READING').length || 0,
      completed: articlesData.filter(a => a.readingStatus === 'READING_STATUS_READ').length || 0,
      abandoned: articlesData.filter(a => a.readingStatus === 'READING_STATUS_ABANDONED').length || 0,
      defaultLibraryCount: librariesData.defaultLibrary?.articles?.length || 0,
      privateLibraryCount: librariesData.privateLibraries?.length || 0,
      privateArticleCount: librariesData.privateLibraries?.reduce((sum, lib) => sum + (lib.articles?.length || 0), 0) || 0
    }
  } catch (error) {
    console.error('Failed to load profile:', error)
    // Handle case where profile doesn't exist yet
    if (error.statusCode === 404) {
      profile.value = null
    }
  } finally {
    isLoading.value = false
  }
}

// Mock recent activity
const recentActivity = ref([
  {
    id: 1,
    description: 'Completed reading "Deep Learning for Computer Vision"',
    timestamp: new Date(Date.now() - 2 * 60 * 60 * 1000), // 2 hours ago
    icon: 'heroicons:check-circle',
    iconColor: 'text-green-400'
  },
  {
    id: 2,
    description: 'Added "Neural Networks in NLP" to Machine Learning library',
    timestamp: new Date(Date.now() - 6 * 60 * 60 * 1000), // 6 hours ago
    icon: 'heroicons:plus-circle',
    iconColor: 'text-blue-400'
  },
  {
    id: 3,
    description: 'Created new library "Research Methods"',
    timestamp: new Date(Date.now() - 24 * 60 * 60 * 1000), // 1 day ago
    icon: 'heroicons:folder-plus',
    iconColor: 'text-purple-400'
  },
  {
    id: 4,
    description: 'Started reading "Attention Is All You Need"',
    timestamp: new Date(Date.now() - 2 * 24 * 60 * 60 * 1000), // 2 days ago
    icon: 'heroicons:book-open',
    iconColor: 'text-orange-400'
  }
])

// Computed
const avatarInitials = computed(() => {
  const names = profile.value.name.split(' ')
  return names.map(name => name.charAt(0)).join('').toUpperCase()
})

// Methods
const toggleEditMode = () => {
  if (isEditing.value) {
    cancelEdit()
  } else {
    startEdit()
  }
}

const startEdit = () => {
  editForm.value = {
    name: profile.value.name,
    bio: profile.value.bio || '',
    institution: profile.value.institution || ''
  }
  isEditing.value = true
}

const cancelEdit = () => {
  isEditing.value = false
  editForm.value = {
    name: '',
    bio: '',
    institution: ''
  }
}

const saveProfile = async () => {
  isSaving.value = true
  
  try {
    if (profile.value) {
      // Update existing profile
      await $fetch(`/api/profile/${profile.value.id}`, {
        method: 'PUT',
        body: {
          name: editForm.value.name,
          bio: editForm.value.bio,
          institution: editForm.value.institution
        }
      })
    } else {
      // Create new profile
      const response = await $fetch('/api/profile', {
        method: 'POST',
        body: {
          name: editForm.value.name,
          bio: editForm.value.bio,
          institution: editForm.value.institution
        }
      })
      profile.value = { id: response.id, ...editForm.value }
    }
    
    // Reload profile to get updated data
    await loadProfile()
    isEditing.value = false
    showSuccess('Profile updated successfully!')
  } catch (error) {
    console.error('Failed to save profile:', error)
    showError('Failed to save profile. Please try again.')
  } finally {
    isSaving.value = false
  }
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('en-US', {
    month: 'long',
    year: 'numeric'
  })
}

const formatRelativeTime = (date) => {
  if (!date) return ''
  
  const now = new Date()
  const diff = now - new Date(date)
  const hours = Math.floor(diff / (1000 * 60 * 60))
  const days = Math.floor(hours / 24)
  
  if (hours < 1) return 'Just now'
  if (hours < 24) return `${hours} hour${hours > 1 ? 's' : ''} ago`
  if (days < 7) return `${days} day${days > 1 ? 's' : ''} ago`
  
  return new Date(date).toLocaleDateString()
}

// Set page title
useHead({
  title: 'Profile - Academic Reading List Manager'
})

// Load profile on mount
onMounted(() => {
  loadProfile()
})
</script>
