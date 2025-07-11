<template>
  <span
    :class="[
      'inline-flex items-center px-2 py-1 rounded-md text-xs font-medium',
      variantClasses,
      sizeClasses
    ]"
  >
    <Icon v-if="icon" :name="icon" :size="iconSize" class="mr-1" />
    <slot />
  </span>
</template>

<script setup lang="ts">
interface Props {
  variant?: 'default' | 'primary' | 'success' | 'warning' | 'error' | 'info'
  size?: 'sm' | 'md' | 'lg'
  icon?: string
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'default',
  size: 'md'
})

const variantClasses = computed(() => {
  switch (props.variant) {
    case 'primary':
      return 'bg-blue-100 text-blue-800 border border-blue-200'
    case 'success':
      return 'bg-green-100 text-green-800 border border-green-200'
    case 'warning':
      return 'bg-yellow-100 text-yellow-800 border border-yellow-200'
    case 'error':
      return 'bg-red-100 text-red-800 border border-red-200'
    case 'info':
      return 'bg-cyan-100 text-cyan-800 border border-cyan-200'
    default:
      return 'bg-gray-100 text-gray-800 border border-gray-200'
  }
})

const sizeClasses = computed(() => {
  switch (props.size) {
    case 'sm':
      return 'px-1.5 py-0.5 text-xs'
    case 'lg':
      return 'px-3 py-1.5 text-sm'
    default:
      return 'px-2 py-1 text-xs'
  }
})

const iconSize = computed(() => {
  switch (props.size) {
    case 'sm':
      return 12
    case 'lg':
      return 16
    default:
      return 14
  }
})
</script>
