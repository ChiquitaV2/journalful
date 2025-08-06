<template>
  <div class="min-h-screen">
    <!-- Mobile Header -->
    <header class="glass-nav fixed top-0 left-0 right-0 z-40 h-16">
      <div class="flex items-center justify-between h-full mobile-padding">
        <div class="flex items-center space-x-3">
          <NuxtLink to="/profile" class="glass-button !p-2">
            <Icon name="heroicons:arrow-left" class="h-6 w-6" />
          </NuxtLink>
          <h1 class="text-lg font-bold text-white">Settings</h1>
        </div>
        
        <div class="flex items-center space-x-2">
          <button @click="saveSettings" class="glass-button !p-2">
            <Icon name="heroicons:check" class="h-5 w-5" />
          </button>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="pt-16">
      <div class="mobile-padding max-w-4xl mx-auto pb-20">
      <!-- Settings Header -->
      <div class="glass-card p-6 mb-6">
        <div class="flex items-center gap-4 mb-4">
          <NuxtLink to="/profile" class="glass-button">
            <Icon name="heroicons:arrow-left" class="h-4 w-4 mr-2" />
            Back
          </NuxtLink>
          <h1 class="text-2xl font-bold text-white flex-1">Settings</h1>
        </div>
        <p class="text-gray-300">Manage your account preferences and app settings.</p>
      </div>

      <!-- Settings Sections -->
      <div class="space-y-6">
        <!-- Profile Settings -->
        <div class="glass-card p-6">
          <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:user-circle" class="h-5 w-5 text-purple-400" />
            Profile Settings
          </h2>
          
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">Display Name</label>
              <input
                v-model="settings.profile.displayName"
                type="text"
                class="w-full glass-input px-4 py-3"
                placeholder="Your display name"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">Email</label>
              <input
                v-model="settings.profile.email"
                type="email"
                class="w-full glass-input px-4 py-3"
                placeholder="your.email@example.com"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">Institution</label>
              <input
                v-model="settings.profile.institution"
                type="text"
                class="w-full glass-input px-4 py-3"
                placeholder="Your university or organization"
              />
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">Research Field</label>
              <input
                v-model="settings.profile.researchField"
                type="text"
                class="w-full glass-input px-4 py-3"
                placeholder="e.g., Computer Science, Biology"
              />
            </div>
          </div>
        </div>

        <!-- Reading Preferences -->
        <div class="glass-card p-6">
          <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:book-open" class="h-5 w-5 text-purple-400" />
            Reading Preferences
          </h2>
          
          <div class="space-y-6">
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Default reading status for new articles</span>
                <select v-model="settings.reading.defaultStatus" class="glass-select">
                  <option value="unread">Unread</option>
                  <option value="in-progress">In Progress</option>
                </select>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Articles per page</span>
                <select v-model="settings.reading.articlesPerPage" class="glass-select">
                  <option value="10">10</option>
                  <option value="20">20</option>
                  <option value="50">50</option>
                </select>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Auto-mark as reading when opened</span>
                <button
                  @click="settings.reading.autoMarkReading = !settings.reading.autoMarkReading"
                  :class="[
                    'glass-toggle',
                    settings.reading.autoMarkReading ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Show reading progress in lists</span>
                <button
                  @click="settings.reading.showProgress = !settings.reading.showProgress"
                  :class="[
                    'glass-toggle',
                    settings.reading.showProgress ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
          </div>
        </div>

        <!-- Notification Settings -->
        <div class="glass-card p-6">
          <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:bell" class="h-5 w-5 text-purple-400" />
            Notifications
          </h2>
          
          <div class="space-y-6">
            <div>
              <label class="flex items-center justify-between">
                <div>
                  <span class="text-gray-300 block">Reading reminders</span>
                  <span class="text-sm text-gray-400">Get reminded to read articles</span>
                </div>
                <button
                  @click="settings.notifications.readingReminders = !settings.notifications.readingReminders"
                  :class="[
                    'glass-toggle',
                    settings.notifications.readingReminders ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <div>
                  <span class="text-gray-300 block">New article suggestions</span>
                  <span class="text-sm text-gray-400">Get notified about relevant papers</span>
                </div>
                <button
                  @click="settings.notifications.suggestions = !settings.notifications.suggestions"
                  :class="[
                    'glass-toggle',
                    settings.notifications.suggestions ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <div>
                  <span class="text-gray-300 block">Weekly reading summary</span>
                  <span class="text-sm text-gray-400">Weekly progress reports</span>
                </div>
                <button
                  @click="settings.notifications.weeklyReport = !settings.notifications.weeklyReport"
                  :class="[
                    'glass-toggle',
                    settings.notifications.weeklyReport ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
            
            <div v-if="settings.notifications.readingReminders">
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Reminder frequency</span>
                <select v-model="settings.notifications.reminderFrequency" class="glass-select">
                  <option value="daily">Daily</option>
                  <option value="weekly">Weekly</option>
                  <option value="biweekly">Bi-weekly</option>
                </select>
              </label>
            </div>
          </div>
        </div>

        <!-- Appearance Settings -->
        <div class="glass-card p-6">
          <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:paint-brush" class="h-5 w-5 text-purple-400" />
            Appearance
          </h2>
          
          <div class="space-y-6">
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Theme</span>
                <select v-model="settings.appearance.theme" class="glass-select">
                  <option value="auto">Auto (System)</option>
                  <option value="dark">Dark</option>
                  <option value="light">Light</option>
                </select>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Default view mode</span>
                <select v-model="settings.appearance.defaultView" class="glass-select">
                  <option value="grid">Grid</option>
                  <option value="list">List</option>
                </select>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Compact mode</span>
                <button
                  @click="settings.appearance.compactMode = !settings.appearance.compactMode"
                  :class="[
                    'glass-toggle',
                    settings.appearance.compactMode ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <span class="text-gray-300">Show article previews</span>
                <button
                  @click="settings.appearance.showPreviews = !settings.appearance.showPreviews"
                  :class="[
                    'glass-toggle',
                    settings.appearance.showPreviews ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
          </div>
        </div>

        <!-- Privacy & Data -->
        <div class="glass-card p-6">
          <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:shield-check" class="h-5 w-5 text-purple-400" />
            Privacy & Data
          </h2>
          
          <div class="space-y-6">
            <div>
              <label class="flex items-center justify-between">
                <div>
                  <span class="text-gray-300 block">Analytics</span>
                  <span class="text-sm text-gray-400">Help improve the app</span>
                </div>
                <button
                  @click="settings.privacy.analytics = !settings.privacy.analytics"
                  :class="[
                    'glass-toggle',
                    settings.privacy.analytics ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
            
            <div>
              <label class="flex items-center justify-between">
                <div>
                  <span class="text-gray-300 block">Auto-sync</span>
                  <span class="text-sm text-gray-400">Sync data across devices</span>
                </div>
                <button
                  @click="settings.privacy.autoSync = !settings.privacy.autoSync"
                  :class="[
                    'glass-toggle',
                    settings.privacy.autoSync ? 'glass-toggle-active' : ''
                  ]"
                >
                  <div class="glass-toggle-thumb"/>
                </button>
              </label>
            </div>
            
            <div class="pt-4 border-t border-white/10">
              <button class="glass-button text-red-400 hover:bg-red-500/20 w-full">
                <Icon name="heroicons:trash" class="h-4 w-4 mr-2" />
                Delete All Data
              </button>
              <p class="text-xs text-gray-500 mt-2 text-center">
                This action cannot be undone
              </p>
            </div>
          </div>
        </div>

        <!-- Data Management -->
        <div class="glass-card p-6">
          <h2 class="text-xl font-semibold text-white mb-4 flex items-center gap-2">
            <Icon name="heroicons:cloud-arrow-down" class="h-5 w-5 text-purple-400" />
            Data Management
          </h2>
          
          <div class="space-y-4">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <button class="glass-button justify-center p-4">
                <Icon name="heroicons:cloud-arrow-down" class="h-5 w-5 mr-2" />
                Export Data
              </button>
              
              <button class="glass-button justify-center p-4">
                <Icon name="heroicons:cloud-arrow-up" class="h-5 w-5 mr-2" />
                Import Data
              </button>
              
              <button class="glass-button justify-center p-4">
                <Icon name="heroicons:document-arrow-down" class="h-5 w-5 mr-2" />
                Backup Libraries
              </button>
              
              <button class="glass-button justify-center p-4">
                <Icon name="heroicons:arrow-path" class="h-5 w-5 mr-2" />
                Sync Now
              </button>
            </div>
            
            <div class="text-sm text-gray-400 text-center pt-4">
              Last sync: {{ formatDate(lastSync) }}
            </div>
          </div>
        </div>
      </div>

      <!-- Save Button -->
      <div class="sticky bottom-6 mt-8">
        <div class="glass-card p-4">
          <div class="flex gap-4">
            <button @click="resetSettings" class="glass-button flex-1">
              <Icon name="heroicons:arrow-path" class="h-4 w-4 mr-2" />
              Reset to Defaults
            </button>
            <button @click="saveSettings" class="glass-button-primary flex-1">
              <Icon name="heroicons:check" class="h-4 w-4 mr-2" />
              Save Changes
            </button>
          </div>
        </div>
      </div>
    </div>
    </main>

    <!-- Mobile Bottom Navigation -->
    <MobileBottomNav />
  </div>
</template>

<script setup>
// Reactive settings
const { success, error: showError, warning } = useToast()
const settings = ref({
  profile: {
    displayName: 'Dr. Sarah Chen',
    email: 'sarah.chen@university.edu',
    institution: 'MIT',
    researchField: 'Computer Science'
  },
  reading: {
    defaultStatus: 'unread',
    articlesPerPage: 20,
    autoMarkReading: true,
    showProgress: true
  },
  notifications: {
    readingReminders: true,
    suggestions: true,
    weeklyReport: false,
    reminderFrequency: 'weekly'
  },
  appearance: {
    theme: 'auto',
    defaultView: 'grid',
    compactMode: false,
    showPreviews: true
  },
  privacy: {
    analytics: true,
    autoSync: true
  }
})

const lastSync = ref('2024-07-28T10:30:00Z')
const originalSettings = ref(JSON.parse(JSON.stringify(settings.value)))

// Methods
const saveSettings = async () => {
  try {
    // Simulate API call
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    // Simulate potential errors for demo
    if (Math.random() < 0.1) { // 10% chance of error for demo
      throw new Error('Settings service temporarily unavailable')
    }
    
    // Update original settings
    originalSettings.value = JSON.parse(JSON.stringify(settings.value))
    
    success('Settings Saved', 'Your preferences have been saved successfully.')
  } catch (err) {
    showError('Save Failed', 'Unable to save settings. Please check your connection and try again.')
  }
}

const resetSettings = async () => {
  try {
    if (confirm('Are you sure you want to reset all settings to their defaults?')) {
      // Reset to original values
      settings.value = JSON.parse(JSON.stringify(originalSettings.value))
      warning('Settings Reset', 'All settings have been restored to their previous saved state.')
    }
  } catch (err) {
    showError('Reset Failed', 'Unable to reset settings.')
  }
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

// Watch for changes to show save indicator
const hasChanges = computed(() => {
  return JSON.stringify(settings.value) !== JSON.stringify(originalSettings.value)
})

// Meta
useHead({
  title: 'Settings - Academic Reader',
  meta: [
    { name: 'description', content: 'Manage your account preferences and app settings' }
  ]
})
</script>
