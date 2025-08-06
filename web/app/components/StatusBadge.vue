<template>
  <span
    :class="[
      'inline-flex items-center rounded-full px-2 py-1 text-xs font-medium',
      statusClasses,
      sizeClasses
    ]"
  >
    <Icon :name="statusIcon" :class="iconSizeClasses" class="mr-1" />
    {{ statusLabel }}
  </span>
</template>

<script setup>
import { ReadingStatus } from '~~/types'

const props = defineProps({
  status: {
    type: Number,
    required: true
  },
  size: {
    type: String,
    default: 'md',
    validator: (value) => ['xs', 'sm', 'md', 'lg'].includes(value)
  }
})

const statusConfig = {
  [ReadingStatus.READING_STATUS_TO_READ]: {
    label: 'To Read',
    icon: 'heroicons:bookmark',
    classes: 'bg-blue-500/20 text-blue-300 border border-blue-300/30'
  },
  [ReadingStatus.READING_STATUS_READING]: {
    label: 'Reading',
    icon: 'heroicons:book-open',
    classes: 'bg-orange-500/20 text-orange-300 border border-orange-300/30'
  },
  [ReadingStatus.READING_STATUS_READ]: {
    label: 'Read',
    icon: 'heroicons:check-circle',
    classes: 'bg-green-500/20 text-green-300 border border-green-300/30'
  },
  [ReadingStatus.READING_STATUS_ABANDONED]: {
    label: 'Abandoned',
    icon: 'heroicons:x-circle',
    classes: 'bg-gray-500/20 text-gray-300 border border-gray-300/30'
  },
  [ReadingStatus.READING_STATUS_UNSPECIFIED]: {
    label: 'Unspecified',
    icon: 'heroicons:question-mark-circle',
    classes: 'bg-gray-500/20 text-gray-300 border border-gray-300/30'
  }
}

const config = computed(() => {
  return statusConfig[props.status] || statusConfig[ReadingStatus.READING_STATUS_UNSPECIFIED]
})

const statusLabel = computed(() => config.value.label)
const statusIcon = computed(() => config.value.icon)
const statusClasses = computed(() => config.value.classes)

const sizeClasses = computed(() => {
  const sizes = {
    xs: 'text-xs px-1.5 py-0.5',
    sm: 'text-xs px-2 py-1',
    md: 'text-sm px-2.5 py-1.5',
    lg: 'text-base px-3 py-2'
  }
  return sizes[props.size] || sizes.md
})

const iconSizeClasses = computed(() => {
  const sizes = {
    xs: 'h-3 w-3',
    sm: 'h-3 w-3',
    md: 'h-4 w-4',
    lg: 'h-5 w-5'
  }
  return sizes[props.size] || sizes.md
})
</script>
