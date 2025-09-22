<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
    <div class="max-w-md w-full space-y-8">
      <div class="text-center">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Welcome to Journalful!</h1>
        <p class="text-gray-600 mb-8">Let's set up your academic reading workspace</p>
      </div>

      <div class="bg-white rounded-lg shadow-lg p-8">
        <form @submit.prevent="completeSetup" class="space-y-6">
          <!-- Progress indicator -->
          <div class="flex items-center justify-center mb-6">
            <div class="flex items-center space-x-2">
              <div class="w-8 h-8 bg-blue-600 text-white rounded-full flex items-center justify-center text-sm font-semibold">1</div>
              <div class="w-16 h-1 bg-gray-200"></div>
              <div class="w-8 h-8 bg-gray-200 text-gray-500 rounded-full flex items-center justify-center text-sm">2</div>
            </div>
          </div>

          <div v-if="!isCreating">
            <h2 class="text-xl font-semibold text-gray-900 mb-4">Complete Your Profile</h2>
            
            <!-- Name field -->
            <div class="mb-4">
              <label for="name" class="block text-sm font-medium text-gray-700 mb-2">
                Display Name *
              </label>
              <input
                id="name"
                v-model="profileForm.name"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                placeholder="Dr. Jane Smith"
              />
            </div>

            <!-- Institution field -->
            <div class="mb-4">
              <label for="institution" class="block text-sm font-medium text-gray-700 mb-2">
                Institution
              </label>
              <input
                id="institution"
                v-model="profileForm.institution"
                type="text"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors"
                placeholder="Stanford University"
              />
            </div>

            <!-- Bio field -->
            <div class="mb-6">
              <label for="bio" class="block text-sm font-medium text-gray-700 mb-2">
                Bio (Optional)
              </label>
              <textarea
                id="bio"
                v-model="profileForm.bio"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors resize-none"
                placeholder="Tell us about your research interests..."
              ></textarea>
            </div>

            <!-- Submit button -->
            <button
              type="submit"
              :disabled="isCreating || !profileForm.name.trim()"
              class="w-full bg-blue-600 text-white py-3 px-4 rounded-lg font-medium hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              Complete Setup
            </button>
          </div>

          <!-- Loading state -->
          <div v-else class="text-center py-8">
            <div class="inline-flex items-center space-x-3">
              <div class="animate-spin rounded-full h-6 w-6 border-b-2 border-blue-600"></div>
              <span class="text-gray-600">Setting up your workspace...</span>
            </div>
            <div class="mt-4 space-y-2 text-sm text-gray-500">
              <div :class="{ 'text-green-600': setupSteps.profile }">
                <Icon :name="setupSteps.profile ? 'heroicons:check-circle' : 'heroicons:clock'" class="inline w-4 h-4 mr-2" />
                Creating your profile
              </div>
              <div :class="{ 'text-green-600': setupSteps.library }">
                <Icon :name="setupSteps.library ? 'heroicons:check-circle' : 'heroicons:clock'" class="inline w-4 h-4 mr-2" />
                Setting up your default library
              </div>
              <div :class="{ 'text-green-600': setupSteps.complete }">
                <Icon :name="setupSteps.complete ? 'heroicons:check-circle' : 'heroicons:clock'" class="inline w-4 h-4 mr-2" />
                Finalizing setup
              </div>
            </div>
          </div>
        </form>
      </div>

      <!-- Additional info -->
      <div class="text-center text-sm text-gray-500">
        <p>🔒 Your information is secure and can be updated anytime</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  layout: false,
  public: true,
  middleware: [] // Skip middleware for setup page
})

const { user } = useUserSession()
const { addNotification } = useNotifications()

// Convenience wrappers with default durations
const showSuccess = (message: string) => addNotification({ type: 'success', message, duration: 5000 })
const showError = (message: string) => addNotification({ type: 'error', message, duration: 5000 })

// Form state
const isCreating = ref(false)
const profileForm = ref({
  name: '',
  bio: '',
  institution: ''
})

// Setup progress tracking
const setupSteps = ref({
  profile: false,
  library: false,
  complete: false
})

// Initialize form with user data if available
onMounted(() => {
  if (user.value) {
    const userData = user.value as any
    profileForm.value.name = userData.name || userData.given_name || ''
  }
})

const completeSetup = async () => {
  if (!profileForm.value.name.trim()) {
    showError('Please enter your name')
    return
  }

  isCreating.value = true
  
  try {
    // Step 1: Create profile
    const response = await $fetch('/api/profile', {
      method: 'POST',
      body: {
        name: profileForm.value.name.trim(),
        bio: profileForm.value.bio.trim() || null,
        institution: profileForm.value.institution.trim() || null
      }
    })
    
    setupSteps.value.profile = true
    await new Promise(resolve => setTimeout(resolve, 500)) // Show progress
    
    setupSteps.value.library = true
    await new Promise(resolve => setTimeout(resolve, 500)) // Show progress
    
    setupSteps.value.complete = true
    await new Promise(resolve => setTimeout(resolve, 500)) // Show progress
    
    showSuccess('Welcome to Journalful! Your workspace is ready.')
    
    // Redirect to main app
    await navigateTo('/')
    
  } catch (error: any) {
    console.error('Setup failed:', error)
    showError(error.data?.message || 'Setup failed. Please try again.')
    isCreating.value = false
    setupSteps.value = { profile: false, library: false, complete: false }
  }
}

// Set page title
useHead({
  title: 'Setup - Journalful'
})
</script>

<style scoped>
/* Additional styling if needed */
</style>