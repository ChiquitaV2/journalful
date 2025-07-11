<template>
  <button
    :class="[
      'inline-flex items-center justify-center font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed',
      sizeClasses,
      variantClasses,
      { 'opacity-50 cursor-not-allowed': loading }
    ]"
    :disabled="disabled || loading"
    @click="$emit('click', $event)"
  >
    <LoadingSpinner v-if="loading" :size="iconSize" class="mr-2" />
    <Icon v-else-if="icon" :name="icon" :size="iconSize" class="mr-2" />
    <slot />
  </button>
</template>

<script setup lang="ts">
interface Props {
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost' | 'danger'
  size?: 'sm' | 'md' | 'lg'
  disabled?: boolean
  loading?: boolean
  icon?: string
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  disabled: false,
  loading: false
})

defineEmits<{
  click: [event: MouseEvent]
}>()

const sizeClasses = computed(() => {
  switch (props.size) {
    case 'sm':
      return 'px-3 py-1.5 text-sm rounded-md'
    case 'lg':
      return 'px-6 py-3 text-lg rounded-lg'
    default:
      return 'px-4 py-2 text-base rounded-md'
  }
})

const variantClasses = computed(() => {
  switch (props.variant) {
    case 'secondary':
      return 'bg-gray-100 text-gray-900 hover:bg-gray-200 focus:ring-gray-500'
    case 'outline':
      return 'border border-gray-300 bg-white text-gray-700 hover:bg-gray-50 focus:ring-primary'
    case 'ghost':
      return 'text-gray-700 hover:bg-gray-100 focus:ring-gray-500'
    case 'danger':
      return 'bg-red-600 text-white hover:bg-red-700 focus:ring-red-500'
    default:
      return 'bg-primary text-white hover:bg-blue-700 focus:ring-primary'
  }
})

const iconSize = computed(() => {
  switch (props.size) {
    case 'sm':
      return 16
    case 'lg':
      return 20
    default:
      return 18
  }
})
</script>
