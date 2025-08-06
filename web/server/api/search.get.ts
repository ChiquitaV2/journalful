import { generateMockArticle } from './mock/data'

export default defineEventHandler(async (event) => {
  const query = getQuery(event)
  const searchTerm = query.q as string
  const limit = Number(query.limit) || 10
  
  if (!searchTerm || searchTerm.length < 2) {
    return {
      results: [],
      total: 0,
      searchTime: 0
    }
  }

  // Simulate search delay
  await new Promise(resolve => setTimeout(resolve, 500))

  // Generate mock search results
  const results = Array.from({ length: Math.min(limit, 8) }, (_, index) => {
    return generateMockArticle(Date.now() + index, {
      title: `${searchTerm} Research: Paper ${index + 1}`,
      relevanceScore: Math.random() * 0.4 + 0.6 // 0.6-1.0 relevance
    })
  })

  return {
    results,
    total: results.length,
    searchTime: Math.random() * 300 + 100, // 100-400ms
    suggestions: [
      `${searchTerm} applications`,
      `${searchTerm} algorithms`,
      `${searchTerm} methodology`,
      `${searchTerm} survey`
    ]
  }
})
