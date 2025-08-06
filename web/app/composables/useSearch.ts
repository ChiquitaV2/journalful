import { debounce } from '../../utils'

/**
 * Composable for search functionality
 */
export const useSearch = () => {
  const searchQuery = ref('')
  const searchResults = ref<any[]>([])
  const isSearching = ref(false)
  const searchError = ref<string | null>(null)

  const performSearch = async (query: string) => {
    if (!query || query.trim().length < 2) {
      searchResults.value = []
      return
    }

    isSearching.value = true
    searchError.value = null

    try {
      const data = await $fetch('/api/search', {
        query: { q: query.trim() }
      })
      
      searchResults.value = data?.results || []
    } catch (error) {
      searchError.value = 'Failed to search articles. Please try again.'
      searchResults.value = []
    } finally {
      isSearching.value = false
    }
  }

  const debouncedSearch = debounce(performSearch, 300)

  const searchArticles = (query: string) => {
    searchQuery.value = query
    debouncedSearch(query)
  }

  const clearSearch = () => {
    searchQuery.value = ''
    searchResults.value = []
    searchError.value = null
    isSearching.value = false
  }

  const searchByDOI = async (doi: string) => {
    if (!doi) throw new Error('DOI is required')

    try {
      const response = await $fetch('/api/articles/lookup', {
        method: 'POST',
        body: { doi }
      })
      
      return response
    } catch (error) {
      throw new Error('Failed to lookup article by DOI')
    }
  }

  return {
    searchQuery: readonly(searchQuery),
    searchResults: readonly(searchResults),
    isSearching: readonly(isSearching),
    searchError: readonly(searchError),
    searchArticles,
    clearSearch,
    searchByDOI
  }
}
