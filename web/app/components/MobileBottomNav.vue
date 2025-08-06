<template>
  <nav class="glass-nav fixed bottom-0 left-0 right-0 z-40 md:hidden">
    <div class="grid grid-cols-4 py-2">
      <NuxtLink
        to="/"
        class="flex flex-col items-center justify-center py-2 px-1 transition-colors"
        :class="isActive('/') ? 'text-purple-300' : 'text-gray-400'"
      >
        <Icon :name="isActive('/') ? 'heroicons:home-solid' : 'heroicons:home'" class="h-6 w-6 mb-1" />
        <span class="text-xs font-medium">Home</span>
      </NuxtLink>
      
      <NuxtLink
        to="/libraries"
        class="flex flex-col items-center justify-center py-2 px-1 transition-colors"
        :class="isActive('/libraries') ? 'text-purple-300' : 'text-gray-400'"
      >
        <Icon :name="isActive('/libraries') ? 'heroicons:book-open-solid' : 'heroicons:book-open'" class="h-6 w-6 mb-1" />
        <span class="text-xs font-medium">Libraries</span>
      </NuxtLink>
      
      <NuxtLink
        to="/search"
        class="flex flex-col items-center justify-center py-2 px-1 transition-colors"
        :class="isActive('/search') ? 'text-purple-300' : 'text-gray-400'"
      >
        <Icon :name="isActive('/search') ? 'heroicons:magnifying-glass-solid' : 'heroicons:magnifying-glass'" class="h-6 w-6 mb-1" />
        <span class="text-xs font-medium">Search</span>
      </NuxtLink>
      
      <NuxtLink
        to="/profile"
        class="flex flex-col items-center justify-center py-2 px-1 transition-colors"
        :class="isActive('/profile') ? 'text-purple-300' : 'text-gray-400'"
      >
        <Icon :name="isActive('/profile') ? 'heroicons:user-circle-solid' : 'heroicons:user-circle'" class="h-6 w-6 mb-1" />
        <span class="text-xs font-medium">Profile</span>
      </NuxtLink>
    </div>
  </nav>
</template>

<script setup>
const route = useRoute()

const isActive = (path) => {
  if (path === '/') {
    return route.path === '/'
  }
  return route.path.startsWith(path)
}
</script>
