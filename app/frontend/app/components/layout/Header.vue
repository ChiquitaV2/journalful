<template>
  <header class="bg-white border-b border-gray-200 sticky top-0 z-40">
    <div class="container">
      <div class="flex items-center justify-between h-16">
        <!-- Logo -->
        <div class="flex items-center">
          <NuxtLink to="/public" class="flex items-center">
            <Icon name="heroicons:book-open" class="h-8 w-8 text-primary mr-3" />
            <span class="text-xl font-bold text-gray-900">Journalful</span>
          </NuxtLink>
        </div>

        <!-- Desktop Navigation -->
        <nav class="hidden md:flex items-center space-x-8">
          <NuxtLink
            to="/app/pages/articles"
            class="text-gray-700 hover:text-primary transition-colors"
            active-class="text-primary font-semibold"
          >
            Articles
          </NuxtLink>
          <NuxtLink
            to="/library"
            class="text-gray-700 hover:text-primary transition-colors"
            active-class="text-primary font-semibold"
          >
            My Library
          </NuxtLink>
          <NuxtLink
            to="/profile/settings"
            class="text-gray-700 hover:text-primary transition-colors"
            active-class="text-primary font-semibold"
          >
            Profile
          </NuxtLink>
        </nav>

        <!-- Search -->
        <div class="hidden lg:flex items-center flex-1 max-w-md mx-8">
          <div class="relative w-full">
            <Input
              type="search"
              placeholder="Search articles, authors, or DOI..."
              v-model="searchQuery"
              @keyup.enter="handleSearch"
              class="w-full"
            />
            <Button
              variant="ghost"
              size="sm"
              class="absolute right-2 top-1/2 transform -translate-y-1/2"
              @click="handleSearch"
            >
              <Icon name="heroicons:magnifying-glass" class="h-4 w-4" />
            </Button>
          </div>
        </div>

        <!-- User Menu -->
        <div class="flex items-center space-x-4">
          <!-- Notifications -->
          <Button variant="ghost" size="sm">
            <Icon name="heroicons:bell" class="h-5 w-5" />
          </Button>

          <!-- User Avatar -->
          <div class="relative" ref="userMenuRef">
            <Button
              variant="ghost"
              size="sm"
              @click="toggleUserMenu"
              class="flex items-center space-x-2"
            >
              <div class="w-8 h-8 bg-primary rounded-full flex items-center justify-center">
                <Icon name="heroicons:user" class="h-5 w-5 text-white" />
              </div>
              <Icon name="heroicons:chevron-down" class="h-4 w-4" />
            </Button>

            <!-- User Dropdown -->
            <Transition
              enter-active-class="transition ease-out duration-100"
              enter-from-class="transform opacity-0 scale-95"
              enter-to-class="transform opacity-100 scale-100"
              leave-active-class="transition ease-in duration-75"
              leave-from-class="transform opacity-100 scale-100"
              leave-to-class="transform opacity-0 scale-95"
            >
              <div
                v-if="showUserMenu"
                class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg border border-gray-200 py-1 z-50"
              >
                <NuxtLink
                  to="/profile/settings"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  @click="showUserMenu = false"
                >
                  <Icon name="heroicons:user" class="h-4 w-4 mr-2 inline" />
                  Profile Settings
                </NuxtLink>
                <NuxtLink
                  to="/library"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  @click="showUserMenu = false"
                >
                  <Icon name="heroicons:book-open" class="h-4 w-4 mr-2 inline" />
                  My Library
                </NuxtLink>
                <div class="border-t border-gray-100 my-1"></div>
                <NuxtLink
                  v-if="isAdmin"
                  to="/admin"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  @click="showUserMenu = false"
                >
                  <Icon name="heroicons:cog-6-tooth" class="h-4 w-4 mr-2 inline" />
                  Admin Panel
                </NuxtLink>
                <button
                  @click="handleLogout"
                  class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                  <Icon name="heroicons:arrow-right-on-rectangle" class="h-4 w-4 mr-2 inline" />
                  Sign Out
                </button>
              </div>
            </Transition>
          </div>

          <!-- Mobile Menu Toggle -->
          <Button
            variant="ghost"
            size="sm"
            class="md:hidden"
            @click="toggleMobileMenu"
          >
            <Icon name="heroicons:bars-3" class="h-6 w-6" />
          </Button>
        </div>
      </div>

      <!-- Mobile Navigation -->
      <Transition
        enter-active-class="transition ease-out duration-200"
        enter-from-class="opacity-0 -translate-y-1"
        enter-to-class="opacity-100 translate-y-0"
        leave-active-class="transition ease-in duration-150"
        leave-from-class="opacity-100 translate-y-0"
        leave-to-class="opacity-0 -translate-y-1"
      >
        <div v-if="showMobileMenu" class="md:hidden border-t border-gray-200 py-4">
          <!-- Mobile Search -->
          <div class="px-4 mb-4">
            <Input
              type="search"
              placeholder="Search articles, authors, or DOI..."
              v-model="searchQuery"
              @keyup.enter="handleSearch"
            />
          </div>

          <!-- Mobile Navigation Links -->
          <nav class="px-4 space-y-2">
            <NuxtLink
              to="/app/pages/articles"
              class="block py-2 text-gray-700 hover:text-primary transition-colors"
              active-class="text-primary font-semibold"
              @click="showMobileMenu = false"
            >
              Articles
            </NuxtLink>
            <NuxtLink
              to="/library"
              class="block py-2 text-gray-700 hover:text-primary transition-colors"
              active-class="text-primary font-semibold"
              @click="showMobileMenu = false"
            >
              My Library
            </NuxtLink>
            <NuxtLink
              to="/profile/settings"
              class="block py-2 text-gray-700 hover:text-primary transition-colors"
              active-class="text-primary font-semibold"
              @click="showMobileMenu = false"
            >
              Profile
            </NuxtLink>
            <NuxtLink
              v-if="isAdmin"
              to="/admin"
              class="block py-2 text-gray-700 hover:text-primary transition-colors"
              active-class="text-primary font-semibold"
              @click="showMobileMenu = false"
            >
              Admin Panel
            </NuxtLink>
          </nav>
        </div>
      </Transition>
    </div>
  </header>
</template>

<script setup lang="ts">
// For now, we'll use placeholder values until auth is implemented
const user = ref(null)
const isAdmin = ref(false)
const logout = () => Promise.resolve()

const searchQuery = ref('')
const showUserMenu = ref(false)
const showMobileMenu = ref(false)
const userMenuRef = ref<HTMLElement>()

const toggleUserMenu = () => {
  showUserMenu.value = !showUserMenu.value
}

const toggleMobileMenu = () => {
  showMobileMenu.value = !showMobileMenu.value
}

const handleSearch = () => {
  if (searchQuery.value.trim()) {
    navigateTo(`/articles/search?q=${encodeURIComponent(searchQuery.value)}`)
    searchQuery.value = ''
    showMobileMenu.value = false
  }
}

const handleLogout = async () => {
  await logout()
  showUserMenu.value = false
  navigateTo('/')
}

// Close menus when clicking outside
onMounted(() => {
  const handleClickOutside = (event: MouseEvent) => {
    if (userMenuRef.value && !userMenuRef.value.contains(event.target as Node)) {
      showUserMenu.value = false
    }
  }
  document.addEventListener('click', handleClickOutside)
  
  onUnmounted(() => {
    document.removeEventListener('click', handleClickOutside)
  })
})

// Close mobile menu on route change
watch(() => useRoute().path, () => {
  showMobileMenu.value = false
})
</script>
