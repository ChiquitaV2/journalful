<template>
  <div class="fixed inset-0 z-50 glass-overlay">
    <div class="flex items-center justify-center min-h-screen p-4">
      <div
        class="glass-elevated w-full max-w-sm rounded-2xl p-6 animate-scale-in"
        @click.stop
      >
        <!-- Icon -->
        <div class="text-center mb-4">
          <div
            :class="[
              'w-16 h-16 mx-auto rounded-full flex items-center justify-center mb-4',
              variant === 'danger' ? 'bg-red-500/20' : 'bg-yellow-500/20'
            ]"
          >
            <Icon
              :name="variant === 'danger' ? 'heroicons:exclamation-triangle' : 'heroicons:question-mark-circle'"
              :class="[
                'h-8 w-8',
                variant === 'danger' ? 'text-red-400' : 'text-yellow-400'
              ]"
            />
          </div>
        </div>

        <!-- Content -->
        <div class="text-center mb-6">
          <h3 class="text-lg font-semibold text-white mb-2">
            {{ title }}
          </h3>
          <p class="text-gray-400 text-sm">
            {{ message }}
          </p>
        </div>

        <!-- Actions -->
        <div class="flex space-x-3">
          <button
            @click="$emit('cancel')"
            class="flex-1 glass-button"
          >
            {{ cancelText }}
          </button>
          <button
            @click="$emit('confirm')"
            :class="[
              'flex-1 glass-button',
              variant === 'danger'
                ? 'bg-red-500/20 border-red-300/30 text-red-300 hover:bg-red-500/30'
                : 'bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30'
            ]"
          >
            {{ confirmText }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    required: true
  },
  message: {
    type: String,
    required: true
  },
  confirmText: {
    type: String,
    default: 'Confirm'
  },
  cancelText: {
    type: String,
    default: 'Cancel'
  },
  variant: {
    type: String,
    default: 'default',
    validator: (value) => ['default', 'danger'].includes(value)
  }
})

defineEmits(['confirm', 'cancel'])
</script>
