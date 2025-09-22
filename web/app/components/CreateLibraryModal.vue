<template>
  <div class="fixed inset-0 z-50 glass-overlay">
    <div class="flex items-center justify-center min-h-screen p-4">
      <div
        class="glass-elevated w-full max-w-md rounded-2xl p-6 animate-scale-in"
        @click.stop
      >
        <!-- Header -->
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-xl font-bold text-white">Create New Library</h2>
          <button
            @click="$emit('close')"
            class="glass-button !p-2"
          >
            <Icon name="heroicons:x-mark" class="h-5 w-5" />
          </button>
        </div>

        <!-- Form -->
        <form @submit.prevent="createLibrary">
          <div class="space-y-4">
            <!-- Library Name -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Library Name *
              </label>
              <input
                v-model="libraryName"
                type="text"
                required
                class="glass-input"
                placeholder="e.g., Machine Learning Papers"
                maxlength="100"
              />
              <div class="text-xs text-gray-400 mt-1">
                {{ libraryName.length }}/100 characters
              </div>
            </div>

            <!-- Description -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Description (Optional)
              </label>
              <textarea
                v-model="description"
                rows="3"
                class="glass-input resize-none"
                placeholder="Describe what this library is for..."
                maxlength="500"
              ></textarea>
              <div class="text-xs text-gray-400 mt-1">
                {{ description.length }}/500 characters
              </div>
            </div>

            <!-- Privacy Settings -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-3">
                Privacy Settings
              </label>
              <div class="space-y-3">
                <label class="flex items-start space-x-3 cursor-pointer">
                  <input
                    v-model="isPrivate"
                    type="radio"
                    :value="true"
                    name="privacy"
                    class="mt-1 text-purple-500 bg-transparent border-gray-400 focus:ring-purple-500 focus:ring-offset-0"
                  />
                  <div>
                    <div class="font-medium text-white flex items-center">
                      <Icon name="heroicons:lock-closed" class="h-4 w-4 mr-2 text-blue-400" />
                      Private Library
                    </div>
                    <p class="text-sm text-gray-400">Only you can see this library</p>
                  </div>
                </label>
                
                <label class="flex items-start space-x-3 cursor-pointer">
                  <input
                    v-model="isPrivate"
                    type="radio"
                    :value="false"
                    name="privacy"
                    class="mt-1 text-purple-500 bg-transparent border-gray-400 focus:ring-purple-500 focus:ring-offset-0"
                  />
                  <div>
                    <div class="font-medium text-white flex items-center">
                      <Icon name="heroicons:globe-alt" class="h-4 w-4 mr-2 text-green-400" />
                      Public Library
                    </div>
                    <p class="text-sm text-gray-400">Anyone can view this library</p>
                  </div>
                </label>
              </div>
            </div>

            <!-- Tags -->
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-2">
                Tags (Optional)
              </label>
              <div class="flex flex-wrap gap-2 mb-2">
                <span
                  v-for="tag in tags"
                  :key="tag"
                  class="inline-flex items-center px-2 py-1 rounded-full text-xs bg-purple-500/20 text-purple-300 border border-purple-300/30"
                >
                  {{ tag }}
                  <button
                    @click="removeTag(tag)"
                    class="ml-1 hover:text-purple-100"
                  >
                    <Icon name="heroicons:x-mark" class="h-3 w-3" />
                  </button>
                </span>
              </div>
              <div class="flex space-x-2">
                <input
                  v-model="newTag"
                  type="text"
                  class="glass-input flex-1"
                  placeholder="Add a tag..."
                  @keyup.enter="addTag"
                  maxlength="20"
                />
                <button
                  type="button"
                  @click="addTag"
                  :disabled="!newTag.trim() || tags.length >= 5"
                  class="glass-button !py-3 !px-4 disabled:opacity-50"
                >
                  <Icon name="heroicons:plus" class="h-4 w-4" />
                </button>
              </div>
              <div class="text-xs text-gray-400 mt-1">
                Maximum 5 tags, 20 characters each
              </div>
            </div>
          </div>

          <!-- Form Actions -->
          <div class="flex space-x-3 mt-6">
            <button
              type="button"
              @click="$emit('close')"
              class="flex-1 glass-button"
            >
              Cancel
            </button>
            <button
              type="submit"
              :disabled="isCreating || !libraryName.trim()"
              class="flex-1 glass-button bg-gradient-to-r from-purple-500/20 to-blue-500/20 border-purple-300/30 disabled:opacity-50"
            >
              <Icon
                v-if="isCreating"
                name="heroicons:arrow-path"
                class="h-4 w-4 mr-2 animate-spin"
              />
              Create Library
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(['close', 'library-created'])
const { showSuccess, showError } = useNotifications()

// Form state
const libraryName = ref('')
const description = ref('')
const isPrivate = ref(true)
const tags = ref([])
const newTag = ref('')
const isCreating = ref(false)

const addTag = () => {
  const tag = newTag.value.trim().toLowerCase()
  if (tag && !tags.value.includes(tag) && tags.value.length < 5) {
    tags.value.push(tag)
    newTag.value = ''
  }
}

const removeTag = (tagToRemove) => {
  const index = tags.value.indexOf(tagToRemove)
  if (index > -1) {
    tags.value.splice(index, 1)
  }
}

const createLibrary = async () => {
  if (!libraryName.value.trim()) return
  
  isCreating.value = true
  
  try {
    // Call actual API endpoint
    const response = await $fetch('/api/libraries', {
      method: 'POST',
      body: {
        name: libraryName.value.trim(),
        description: description.value.trim() || null,
        isPublic: !isPrivate.value // Backend expects isPublic, frontend uses isPrivate
      }
    })
    
    const newLibrary = {
      id: response.id,
      name: libraryName.value.trim(),
      description: description.value.trim(),
      isPrivate: isPrivate.value,
      tags: tags.value,
      articles: [],
      createdAt: new Date(),
      updatedAt: new Date()
    }
    
    emit('library-created', newLibrary)
    showSuccess('Library created successfully!')
  } catch (error) {
    console.error('Library creation error:', error)
    showError('Failed to create library. Please try again.')
  } finally {
    isCreating.value = false
  }
}

// Reset form when modal opens
onMounted(() => {
  libraryName.value = ''
  description.value = ''
  isPrivate.value = true
  tags.value = []
  newTag.value = ''
})
</script>
