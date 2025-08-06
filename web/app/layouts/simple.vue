<template>
  <div class="min-h-screen">
    <!-- Simple Header (no navigation) -->
    <header class="glass-nav fixed top-0 left-0 right-0 z-40 h-16">
      <div class="flex items-center justify-between h-full mobile-padding">
        <div class="flex items-center space-x-3">
          <NuxtLink 
            :to="backRoute" 
            class="glass-button !p-2"
          >
            <Icon name="heroicons:arrow-left" class="h-6 w-6" />
          </NuxtLink>
          <h1 class="text-xl font-bold text-white">{{ pageTitle }}</h1>
        </div>
        
        <div class="flex items-center gap-2">
          <!-- Action Button -->
          <button 
            v-if="showActionButton"
            @click="handleAction"
            class="glass-button !p-2"
          >
            <Icon :name="actionIcon" class="h-5 w-5" />
          </button>
        </div>
      </div>
    </header>

    <!-- Main Content (full viewport minus header) -->
    <main class="pt-16">
      <div class="mobile-padding max-w-7xl mx-auto">
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup>
const props = defineProps({
  pageTitle: {
    type: String,
    default: 'Academic Reader'
  },
  backRoute: {
    type: String,
    default: '/'
  },
  showActionButton: {
    type: Boolean,
    default: false
  },
  actionIcon: {
    type: String,
    default: 'heroicons:ellipsis-horizontal'
  }
})

// Emit events for parent pages to handle
const emit = defineEmits(['action'])

const handleAction = () => {
  emit('action')
}
</script>
