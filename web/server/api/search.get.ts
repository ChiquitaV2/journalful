import { useServices } from '~~/server/proto/useServices'
import { ListArticlesRequest } from '~~/server/proto/grpc/articles/v1/article'

export default defineEventHandler(async (event) => {
  try {
    await requireUserSession(event)
    
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

    const startTime = Date.now()
    
    // Get articles service
    const articlesSvr = await useServices().getArticlesServiceClient(event)
    
    // Get all articles (for now - in production, implement proper search)
    const response = await articlesSvr.ListArticles(ListArticlesRequest.fromJSON({}))
    
    // Filter articles by search term (basic text matching)
    const searchTermLower = searchTerm.toLowerCase()
    const filteredResults = response.articles.filter(article => {
      return (
        article.title.toLowerCase().includes(searchTermLower) ||
        article.abstract?.toLowerCase().includes(searchTermLower) ||
        article.authors.some(author => 
          author.name.toLowerCase().includes(searchTermLower)
        ) ||
        article.journalName?.toLowerCase().includes(searchTermLower)
      )
    })
    
    // Limit results
    const results = filteredResults.slice(0, limit)
    
    const searchTime = Date.now() - startTime

    return {
      results,
      total: filteredResults.length,
      searchTime,
      suggestions: [
        `${searchTerm} applications`,
        `${searchTerm} algorithms`, 
        `${searchTerm} methodology`,
        `${searchTerm} survey`
      ]
    }
  } catch (error: any) {
    console.error('Search error:', error)
    return {
      results: [],
      total: 0,
      searchTime: 0,
      error: error.message || 'Search failed'
    }
  }
})
