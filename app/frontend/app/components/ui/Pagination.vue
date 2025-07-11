<template>
  <nav v-if="totalPages > 1" class="flex items-center justify-between">
    <div class="flex items-center gap-2">
      <Button
        variant="outline"
        size="sm"
        :disabled="currentPage === 1"
        @click="goToPage(currentPage - 1)"
      >
        <Icon name="heroicons:chevron-left" class="h-4 w-4" />
        Previous
      </Button>
      
      <Button
        variant="outline"
        size="sm"
        :disabled="currentPage === totalPages"
        @click="goToPage(currentPage + 1)"
      >
        Next
        <Icon name="heroicons:chevron-right" class="h-4 w-4" />
      </Button>
    </div>
    
    <div class="hidden sm:flex items-center gap-1">
      <!-- First page -->
      <Button
        v-if="showFirstPage"
        variant="outline"
        size="sm"
        :class="{ 'bg-primary text-white': currentPage === 1 }"
        @click="goToPage(1)"
      >
        1
      </Button>
      
      <!-- First ellipsis -->
      <span v-if="showFirstEllipsis" class="px-2 py-1 text-gray-500">...</span>
      
      <!-- Page numbers -->
      <Button
        v-for="page in visiblePages"
        :key="page"
        variant="outline"
        size="sm"
        :class="{ 'bg-primary text-white': page === currentPage }"
        @click="goToPage(page)"
      >
        {{ page }}
      </Button>
      
      <!-- Last ellipsis -->
      <span v-if="showLastEllipsis" class="px-2 py-1 text-gray-500">...</span>
      
      <!-- Last page -->
      <Button
        v-if="showLastPage"
        variant="outline"
        size="sm"
        :class="{ 'bg-primary text-white': currentPage === totalPages }"
        @click="goToPage(totalPages)"
      >
        {{ totalPages }}
      </Button>
    </div>
    
    <div class="text-sm text-gray-700">
      Showing {{ startItem }} to {{ endItem }} of {{ totalItems }} results
    </div>
  </nav>
</template>

<script setup lang="ts">
interface Props {
  currentPage: number
  totalPages: number
  totalItems: number
  itemsPerPage: number
  maxVisible?: number
}

const props = withDefaults(defineProps<Props>(), {
  maxVisible: 5
})

const emit = defineEmits<{
  'update:currentPage': [page: number]
}>()

const goToPage = (page: number) => {
  if (page >= 1 && page <= props.totalPages) {
    emit('update:currentPage', page)
  }
}

const startItem = computed(() => {
  return (props.currentPage - 1) * props.itemsPerPage + 1
})

const endItem = computed(() => {
  return Math.min(props.currentPage * props.itemsPerPage, props.totalItems)
})

const visiblePages = computed(() => {
  const delta = Math.floor(props.maxVisible / 2)
  const start = Math.max(1, props.currentPage - delta)
  const end = Math.min(props.totalPages, props.currentPage + delta)
  
  const pages = []
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

const showFirstPage = computed(() => {
  return !visiblePages.value.includes(1)
})

const showLastPage = computed(() => {
  return !visiblePages.value.includes(props.totalPages)
})

const showFirstEllipsis = computed(() => {
  return showFirstPage.value && visiblePages.value[0] && visiblePages.value[0] > 2
})

const showLastEllipsis = computed(() => {
  return showLastPage.value && (visiblePages.value[visiblePages.value.length - 1] ?? visiblePages.value.length) < props.totalPages - 1
})
</script>
