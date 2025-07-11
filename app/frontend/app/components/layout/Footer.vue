<template>
  <footer class="bg-gray-50 border-t border-gray-200 mt-auto">
    <div class="container py-12">
      <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
        <!-- Logo and Description -->
        <div class="col-span-1 md:col-span-2">
          <div class="flex items-center mb-4">
            <Icon name="heroicons:book-open" class="h-8 w-8 text-primary mr-3" />
            <span class="text-xl font-bold text-gray-900">Journalful</span>
          </div>
          <p class="text-gray-600 mb-4 max-w-md">
            Your comprehensive academic journal management system. Organize, track, and discover research articles with ease.
          </p>
          <div class="flex space-x-4">
            <a href="#" class="text-gray-400 hover:text-primary transition-colors">
              <Icon name="heroicons:envelope" class="h-5 w-5" />
            </a>
            <a href="#" class="text-gray-400 hover:text-primary transition-colors">
              <Icon name="simple-icons:twitter" class="h-5 w-5" />
            </a>
            <a href="#" class="text-gray-400 hover:text-primary transition-colors">
              <Icon name="simple-icons:github" class="h-5 w-5" />
            </a>
          </div>
        </div>

        <!-- Quick Links -->
        <div>
          <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wide mb-4">
            Quick Links
          </h3>
          <ul class="space-y-2">
            <li>
              <NuxtLink to="/app/pages/articles" class="text-gray-600 hover:text-primary transition-colors">
                Browse Articles
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/app/pages/articles/search" class="text-gray-600 hover:text-primary transition-colors">
                Search & Add Articles
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/library" class="text-gray-600 hover:text-primary transition-colors">
                My Library
              </NuxtLink>
            </li>
            <li>
              <NuxtLink to="/profile/settings" class="text-gray-600 hover:text-primary transition-colors">
                Profile Settings
              </NuxtLink>
            </li>
          </ul>
        </div>

        <!-- Support -->
        <div>
          <h3 class="text-sm font-semibold text-gray-900 uppercase tracking-wide mb-4">
            Support
          </h3>
          <ul class="space-y-2">
            <li>
              <a href="#" class="text-gray-600 hover:text-primary transition-colors">
                Help Center
              </a>
            </li>
            <li>
              <a href="#" class="text-gray-600 hover:text-primary transition-colors">
                Documentation
              </a>
            </li>
            <li>
              <a href="#" class="text-gray-600 hover:text-primary transition-colors">
                Contact Support
              </a>
            </li>
            <li>
              <a href="#" class="text-gray-600 hover:text-primary transition-colors">
                Privacy Policy
              </a>
            </li>
          </ul>
        </div>
      </div>

      <!-- Bottom Bar -->
      <div class="border-t border-gray-200 mt-8 pt-8 flex flex-col md:flex-row items-center justify-between">
        <p class="text-gray-600 text-sm">
          © {{ currentYear }} Journalful. All rights reserved.
        </p>
        <div class="flex items-center space-x-6 mt-4 md:mt-0">
          <a href="#" class="text-gray-600 hover:text-primary transition-colors text-sm">
            Terms of Service
          </a>
          <a href="#" class="text-gray-600 hover:text-primary transition-colors text-sm">
            Privacy Policy
          </a>
          <a href="#" class="text-gray-600 hover:text-primary transition-colors text-sm">
            Cookie Policy
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
const currentYear = new Date().getFullYear()
</script>
