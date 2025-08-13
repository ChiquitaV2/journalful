<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex items-center justify-center p-4">
    <div class="glass-card max-w-md w-full text-center p-8">
      <!-- Error Icon -->
      <div class="mb-6">
        <div class="p-3 mx-auto rounded-full bg-gradient-to-r from-red-500/20 to-orange-500/20 border border-red-300/30 flex items-center justify-center">
          <Icon
              :name="errorIcon"
              class="h-12 w-12 text-red-400"
          />
        </div>
      </div>

      <!-- Error Content -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-white mb-2">
          {{ errorTitle }}
        </h1>
        <p class="text-gray-400 mb-4">
          {{ errorMessage }}
        </p>

        <!-- Error Details (Development only) -->
        <div v-if="showDetails && error.stack" class="mt-6">
          <details class="text-left">
            <summary class="text-sm text-gray-500 cursor-pointer hover:text-gray-400 mb-2">
              Show technical details
            </summary>
            <div class="glass-card !bg-slate-800/30 p-4 text-xs text-gray-300 font-mono overflow-auto max-h-32">
              <pre>{{ error.stack }}</pre>
            </div>
          </details>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="space-y-3">
        <button
            @click="handleRetry"
            class="w-full glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30"
        >
          <Icon name="heroicons:arrow-path" class="h-4 w-4 mr-2" />
          Try Again
        </button>

        <button
            @click="goHome"
            class="w-full glass-button"
        >
          <Icon name="heroicons:home" class="h-4 w-4 mr-2" />
          Go to Homepage
        </button>

        <button
            v-if="canGoBack"
            @click="goBack"
            class="w-full glass-button"
        >
          <Icon name="heroicons:arrow-left" class="h-4 w-4 mr-2" />
          Go Back
        </button>
      </div>

      <!-- Additional Help -->
      <div class="mt-8 pt-6 border-t border-gray-700/50">
        <p class="text-sm text-gray-500 mb-3">
          Still having trouble?
        </p>
        <div class="flex flex-col sm:flex-row gap-3 text-sm">
          <button
              @click="reportError"
              class="glass-button !py-2 !px-4 text-xs"
          >
            <Icon name="heroicons:bug-ant" class="h-3 w-3 mr-1" />
            Report Issue
          </button>
          <a
              href="mailto:support@example.com"
              class="glass-button !py-2 !px-4 text-xs"
          >
            <Icon name="heroicons:envelope" class="h-3 w-3 mr-1" />
            Contact Support
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Error page props from Nuxt
const props = defineProps({
  error: {
    type: Object,
    required: true
  }
})

// Composables
const { $config } = useNuxtApp()
const router = useRouter()

// State
const canGoBack = ref(false)
const showDetails = computed(() => process.dev)

// Error categorization
const errorInfo = computed(() => {
  const statusCode = props.error.statusCode || 500

  switch (statusCode) {
    case 404:
      return {
        title: 'Page Not Found',
        message: 'The page you\'re looking for doesn\'t exist or has been moved.',
        icon: 'heroicons:document-magnifying-glass'
      }
    case 403:
      return {
        title: 'Access Forbidden',
        message: 'You don\'t have permission to access this resource.',
        icon: 'heroicons:lock-closed'
      }
    case 401:
      return {
        title: 'Authentication Required',
        message: 'Please log in to access this page.',
        icon: 'heroicons:user-circle'
      }
    case 429:
      return {
        title: 'Too Many Requests',
        message: 'Please wait a moment before trying again.',
        icon: 'heroicons:clock'
      }
    case 500:
    case 502:
    case 503:
    case 504:
      return {
        title: 'Server Error',
        message: 'Something went wrong on our end. We\'re working to fix it.',
        icon: 'heroicons:server'
      }
    default:
      return {
        title: 'Something Went Wrong',
        message: props.error.statusMessage || 'An unexpected error occurred.',
        icon: 'heroicons:exclamation-triangle'
      }
  }
})

const errorTitle = computed(() => errorpInfo.value.title)
const errorMessage = computed(() => errorInfo.value.message)
const errorIcon = computed(() => errorInfo.value.icon)

// Methods
const handleRetry = () => {
  // Clear the error and retry
  clearError({ redirect: window.location.pathname })
}

const goHome = () => {
  navigateTo('/')
}

const goBack = () => {
  if (window.history.length > 1) {
    router.back()
  } else {
    goHome()
  }
}

const reportError = () => {
  // In a real app, this would send error details to your error tracking service
  const errorReport = {
    statusCode: props.error.statusCode,
    statusMessage: props.error.statusMessage,
    url: window.location.href,
    userAgent: navigator.userAgent,
    timestamp: new Date().toISOString()
  }

  console.log('Error report:', errorReport)

  // You could integrate with services like Sentry, LogRocket, etc.
  // Example: $sentry.captureException(props.error)

  alert('Error report sent. Thank you for helping us improve!')
}

// Check if user can go back
onMounted(() => {
  canGoBack.value = window.history.length > 1
})

// Set page title based on error
useHead({
  title: computed(() => `${errorInfo.value.title} - Academic Reading List Manager`)
})

// Analytics (optional)
onMounted(() => {
  // Track error for analytics
  if (process.client && window.gtag) {
    window.gtag('event', 'exception', {
      description: `${props.error.statusCode}: ${props.error.statusMessage}`,
      fatal: false
    })
  }
})
</script>

