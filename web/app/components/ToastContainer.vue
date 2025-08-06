<template>
  <Teleport to="body">
    <div class="fixed top-4 right-4 z-50 space-y-3 max-w-sm w-full">
      <TransitionGroup
        enter-active-class="transform transition duration-300 ease-out"
        enter-from-class="translate-x-full opacity-0"
        enter-to-class="translate-x-0 opacity-100"
        leave-active-class="transform transition duration-200 ease-in"
        leave-from-class="translate-x-0 opacity-100"
        leave-to-class="translate-x-full opacity-0"
      >
        <div
          v-for="toast in toasts"
          :key="toast.id"
          :class="[
            'p-4 shadow-lg border-l-4 cursor-pointer group relative overflow-hidden backdrop-blur-lg bg-black/60 border-white/10 rounded-2xl',
            getToastColors(toast.type)
          ]"
          @click="removeToast(toast.id)"
        >
          <!-- Progress bar for timed toasts -->
          <div
            v-if="!toast.persistent && toast.duration"
            class="absolute bottom-0 left-0 h-1 bg-current opacity-30 animate-progress"
            :style="{ animationDuration: `${toast.duration}ms` }"
          ></div>

          <div class="flex items-start gap-3">
            <!-- Icon -->
            <div class="flex-shrink-0 mt-0.5">
              <Icon 
                :name="getToastIcon(toast.type)" 
                class="h-5 w-5"
                :class="getIconColor(toast.type)"
              />
            </div>

            <!-- Content -->
            <div class="flex-1 min-w-0">
              <h4 class="text-sm font-medium text-white mb-1">
                {{ toast.title }}
              </h4>
              <p 
                v-if="toast.message" 
                class="text-sm text-gray-300 leading-relaxed"
              >
                {{ toast.message }}
              </p>
            </div>

            <!-- Close button -->
            <button
              @click.stop="removeToast(toast.id)"
              class="flex-shrink-0 p-1 rounded-full hover:bg-white/10 opacity-60 hover:opacity-100 transition-opacity"
            >
              <Icon name="heroicons:x-mark" class="h-4 w-4 text-white" />
            </button>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup>
const { toasts, removeToast } = useToast()

const getToastColors = (type) => {
  switch (type) {
    case 'success':
      return 'border-green-400 bg-green-900/20'
    case 'error':
      return 'border-red-400 bg-red-900/20'
    case 'warning':
      return 'border-yellow-400 bg-yellow-900/20'
    case 'info':
    default:
      return 'border-blue-400 bg-blue-900/20'
  }
}

const getToastIcon = (type) => {
  switch (type) {
    case 'success':
      return 'heroicons:check-circle'
    case 'error':
      return 'heroicons:exclamation-circle'
    case 'warning':
      return 'heroicons:exclamation-triangle'
    case 'info':
    default:
      return 'heroicons:information-circle'
  }
}

const getIconColor = (type) => {
  switch (type) {
    case 'success':
      return 'text-green-400'
    case 'error':
      return 'text-red-400'
    case 'warning':
      return 'text-yellow-400'
    case 'info':
    default:
      return 'text-blue-400'
  }
}
</script>

<style scoped>
@keyframes progress {
  from {
    width: 100%;
  }
  to {
    width: 0%;
  }
}

.animate-progress {
  animation: progress linear forwards;
}


</style>
