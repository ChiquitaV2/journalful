import type { Article, ListArticlesRequest, ListArticlesResponse } from '~~/shared/grpc/articles/v1/article'
import type { Author } from '~~/shared/grpc/profile/v1/author'

export const useArticles = () => {
  const articles = ref<Article[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  const fetchArticles = async (params?: ListArticlesRequest) => {
    loading.value = true
    error.value = null
    
    try {
      // TODO: Implement actual gRPC API call
      // For now, return mock data
      await new Promise(resolve => setTimeout(resolve, 500))
      
      const mockArticles: Article[] = [
        {
          id: 1,
          doi: '10.1000/182',
          title: 'Machine Learning in Healthcare: A Comprehensive Review',
          authors: [
            { 
              id: 1, 
              name: 'Dr. Sarah Johnson', 
              profileId: 1, 
              createdAt: new Date('2023-01-01'), 
              updatedAt: new Date('2023-01-01') 
            },
            { 
              id: 2, 
              name: 'Dr. Michael Chen', 
              profileId: 2, 
              createdAt: new Date('2023-01-01'), 
              updatedAt: new Date('2023-01-01') 
            }
          ],
          abstract: 'This paper presents a comprehensive review of machine learning applications in healthcare, covering recent advances in medical imaging, drug discovery, and patient diagnosis. We examine the current state of the art and discuss future directions for the field.',
          publicationYear: 2023,
          journalName: 'Journal of Medical Informatics',
          createdAt: new Date('2023-01-15'),
          updatedAt: new Date('2023-01-15')
        },
        {
          id: 2,
          doi: '10.1000/183',
          title: 'Climate Change Impact on Agricultural Productivity',
          authors: [
            { 
              id: 3, 
              name: 'Dr. Emily Rodriguez', 
              profileId: 3, 
              createdAt: new Date('2023-01-01'), 
              updatedAt: new Date('2023-01-01') 
            },
            { 
              id: 4, 
              name: 'Dr. James Wilson', 
              profileId: 4, 
              createdAt: new Date('2023-01-01'), 
              updatedAt: new Date('2023-01-01') 
            }
          ],
          abstract: 'We analyze the effects of climate change on global agricultural productivity over the past two decades. Our findings suggest significant impacts on crop yields and food security worldwide.',
          publicationYear: 2023,
          journalName: 'Environmental Science & Policy',
          createdAt: new Date('2023-02-20'),
          updatedAt: new Date('2023-02-20')
        },
        {
          id: 3,
          doi: '10.1000/184',
          title: 'Quantum Computing: Breaking the Barriers',
          authors: [
            { 
              id: 5, 
              name: 'Dr. David Thompson', 
              profileId: 5, 
              createdAt: new Date('2023-01-01'), 
              updatedAt: new Date('2023-01-01') 
            }
          ],
          abstract: 'This review explores recent breakthroughs in quantum computing technology and their potential applications in cryptography, optimization, and scientific simulation.',
          publicationYear: 2024,
          journalName: 'Nature Quantum Information',
          createdAt: new Date('2024-01-10'),
          updatedAt: new Date('2024-01-10')
        }
      ]
      
      articles.value = mockArticles
      
      return {
        articles: mockArticles,
        total: mockArticles.length,
        hasMore: false
      }
      
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'An error occurred'
      console.error('Error fetching articles:', err)
      return {
        articles: [],
        total: 0,
        hasMore: false
      }
    } finally {
      loading.value = false
    }
  }

  const getArticleById = async (id: number) => {
    loading.value = true
    error.value = null
    
    try {
      // TODO: Implement actual gRPC API call
      await new Promise(resolve => setTimeout(resolve, 300))
      
      const article = articles.value.find(a => a.id === id)
      if (!article) {
        throw new Error('Article not found')
      }
      
      return article
      
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'An error occurred'
      console.error('Error fetching article:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  const searchArticles = async (query: string, filters?: Record<string, any>) => {
    loading.value = true
    error.value = null
    
    try {
      // TODO: Implement actual search API call
      await new Promise(resolve => setTimeout(resolve, 400))
      
      const filteredArticles = articles.value.filter(article => 
        article.title.toLowerCase().includes(query.toLowerCase()) ||
        article.abstract?.toLowerCase().includes(query.toLowerCase()) ||
        article.authors.some(author => author.name.toLowerCase().includes(query.toLowerCase()))
      )
      
      return {
        articles: filteredArticles,
        total: filteredArticles.length,
        hasMore: false
      }
      
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'An error occurred'
      console.error('Error searching articles:', err)
      return {
        articles: [],
        total: 0,
        hasMore: false
      }
    } finally {
      loading.value = false
    }
  }

  const addArticleByDOI = async (doi: string) => {
    loading.value = true
    error.value = null
    
    try {
      // TODO: Implement actual DOI lookup and article creation
      await new Promise(resolve => setTimeout(resolve, 1000))
      
      // Mock article creation
      const newArticle: Article = {
        id: Date.now(),
        doi,
        title: 'New Article Added by DOI',
        authors: [{ 
          id: 999, 
          name: 'Unknown Author', 
          profileId: 999, 
          createdAt: new Date(), 
          updatedAt: new Date() 
        }],
        abstract: 'This article was added by DOI lookup.',
        publicationYear: 2024,
        journalName: 'Unknown Journal',
        createdAt: new Date(),
        updatedAt: new Date()
      }
      
      articles.value.unshift(newArticle)
      return newArticle
      
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'An error occurred'
      console.error('Error adding article by DOI:', err)
      return null
    } finally {
      loading.value = false
    }
  }

  return {
    articles: readonly(articles),
    loading: readonly(loading),
    error: readonly(error),
    fetchArticles,
    getArticleById,
    searchArticles,
    addArticleByDOI
  }
}
