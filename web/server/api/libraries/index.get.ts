import { generateMockLibrary } from '../mock/data'

export default defineEventHandler(event => {
  // Generate mock libraries for the index page
  const defaultLibrary = generateMockLibrary(1, {
    name: 'My Reading List',
    description: 'Your default collection for articles',
    articleCount: 42,
    isDefault: true
  })

  const privateLibraries = [
    generateMockLibrary(2, {
      name: 'Machine Learning Papers',
      description: 'Latest research in ML and AI',
      articleCount: 18,
      icon: 'heroicons:cpu-chip'
    }),
    generateMockLibrary(3, {
      name: 'Climate Research', 
      description: 'Environmental science and climate change studies',
      articleCount: 12,
      icon: 'heroicons:globe-alt'
    }),
    generateMockLibrary(4, {
      name: 'Computer Vision',
      description: 'Image processing and visual recognition research',
      articleCount: 8,
      icon: 'heroicons:camera'
    })
  ]

  return {
    defaultLibrary,
    privateLibraries,
    stats: {
      total: 1 + privateLibraries.length,
      articles: defaultLibrary.articles.length + privateLibraries.reduce((sum, lib) => sum + lib.articles.length, 0),
      completed: [defaultLibrary, ...privateLibraries]
        .flatMap(lib => lib.articles)
        .filter(article => article.readingStatus === 'completed').length
    }
  }
})
