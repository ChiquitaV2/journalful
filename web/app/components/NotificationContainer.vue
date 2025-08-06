<template>
  <Teleport to="body">
    <div class="fixed top-4 right-4 z-50 space-y-2">
      <TransitionGroup name="notification" tag="div">
        <div
          v-for="notification in notifications"
          :key="notification.id"
          :class="[
            'glass-notification animate-slide-up',
            {
              'success': notification.type === 'success',
              'error': notification.type === 'error',
              'warning': notification.type === 'warning'
            }
          ]"
          class="max-w-sm w-full"
        >
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <Icon v-if="notification.type === 'success'" name="heroicons:check-circle" class="h-5 w-5 text-emerald-400" />
              <Icon v-else-if="notification.type === 'error'" name="heroicons:x-circle" class="h-5 w-5 text-red-400" />
              <Icon v-else-if="notification.type === 'warning'" name="heroicons:exclamation-triangle" class="h-5 w-5 text-amber-400" />
              <Icon v-else name="heroicons:information-circle" class="h-5 w-5 text-blue-400" />
            </div>
            <div class="ml-3 flex-1">
              <p class="text-sm font-medium text-white">
                {{ notification.message }}
              </p>
            </div>
            <div class="ml-4 flex-shrink-0">
              <button
                @click="removeNotification(notification.id)"
                class="glass-button !p-1 !bg-white/10 hover:!bg-white/20"
              >
                <Icon name="heroicons:x-mark" class="h-4 w-4" />
              </button>
            </div>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup>
const { notifications, removeNotification } = useNotifications()

// Auto-remove notifications after duration
watch(notifications, (newNotifications) => {
  newNotifications.forEach(notification => {
    if (notification.duration && notification.duration > 0) {
      setTimeout(() => {
        removeNotification(notification.id)
      }, notification.duration)
    }
  })
}, { deep: true, immediate: true })
</script>

<style scoped>
.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.notification-move {
  transition: transform 0.3s ease;
}
</style>
