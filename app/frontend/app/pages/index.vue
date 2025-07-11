<template>
  <div>
    <!-- Hero Section -->
    <section class="bg-gradient-to-r from-blue-600 to-blue-700 text-white py-20">
      <div class="container text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-6">
          Welcome to Journalful
        </h1>
        <p class="text-xl md:text-2xl mb-8 max-w-3xl mx-auto opacity-90">
          Your comprehensive academic journal management system. Organize, track, and discover research articles with ease.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <Button size="lg" variant="secondary" @click="navigateTo('/articles')">
            <Icon name="heroicons:book-open" class="mr-2" />
            Browse Articles
          </Button>
          <Button size="lg" variant="outline" @click="navigateTo('/articles/search')">
            <Icon name="heroicons:plus" class="mr-2" />
            Add Article
          </Button>
        </div>
      </div>
    </section>

    <!-- Features Section -->
    <section class="py-20 bg-white">
      <div class="container">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
            Everything you need to manage your research
          </h2>
          <p class="text-xl text-gray-600 max-w-2xl mx-auto">
            Powerful features designed to help researchers organize, track, and discover academic articles.
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div class="text-center p-6">
            <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <Icon name="heroicons:book-open" class="h-8 w-8 text-primary" />
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-2">
              Article Management
            </h3>
            <p class="text-gray-600">
              Browse, search, and organize academic articles with detailed metadata and full-text search capabilities.
            </p>
          </div>

          <div class="text-center p-6">
            <div class="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <Icon name="heroicons:bookmark" class="h-8 w-8 text-green-600" />
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-2">
              Personal Libraries
            </h3>
            <p class="text-gray-600">
              Create custom reading lists and libraries to organize your research by topic, project, or reading status.
            </p>
          </div>

          <div class="text-center p-6">
            <div class="w-16 h-16 bg-purple-100 rounded-full flex items-center justify-center mx-auto mb-4">
              <Icon name="heroicons:users" class="h-8 w-8 text-purple-600" />
            </div>
            <h3 class="text-xl font-semibold text-gray-900 mb-2">
              Research Collaboration
            </h3>
            <p class="text-gray-600">
              Share your libraries with colleagues and discover research from other users in your field.
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- Stats Section -->
    <section class="py-20 bg-gray-50">
      <div class="container">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 text-center">
          <div>
            <div class="text-4xl font-bold text-primary mb-2">10,000+</div>
            <div class="text-gray-600">Articles in Database</div>
          </div>
          <div>
            <div class="text-4xl font-bold text-primary mb-2">500+</div>
            <div class="text-gray-600">Active Researchers</div>
          </div>
          <div>
            <div class="text-4xl font-bold text-primary mb-2">1,000+</div>
            <div class="text-gray-600">Research Libraries</div>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="py-20 bg-primary text-white">
      <div class="container text-center">
        <h2 class="text-3xl md:text-4xl font-bold mb-4">
          Ready to organize your research?
        </h2>
        <p class="text-xl mb-8 opacity-90">
          Join thousands of researchers who trust Journalful to manage their academic libraries.
        </p>
        <Button size="lg" variant="secondary" @click="navigateTo('/library')">
          Get Started Now
        </Button>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
// Meta tags for SEO
useSeoMeta({
  title: 'Journalful - Academic Journal Management System',
  description: 'Organize, track, and discover research articles with Journalful. Your comprehensive academic journal management system.',
  keywords: 'academic, journal, research, articles, library, bibliography, DOI',
})

// Page head
useHead({
  title: 'Journalful - Home'
})
</script>
