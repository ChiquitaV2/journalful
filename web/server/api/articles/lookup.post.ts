import {useServices} from "~~/server/proto/useServices";
import {
  CreateArticleRequest,
  GetArticleByDOIRequest,
  GetArticleRequest
} from "~~/server/proto/grpc/articles/v1/article";

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const { doi } = body

  if (!doi) {
    throw createError({
      statusCode: 400,
      statusMessage: 'DOI is required'
    })
  }

  const articlesSvr = await useServices().getArticlesServiceClient(event)
  try {
    return await articlesSvr.GetArticleByDOI(GetArticleByDOIRequest.fromJSON({doi}))
  } catch (e) {
    console.error("lookup error", e)
    // For all errors besdies not found return the error
    if (!e.message.includes("not found")) {
      return createError({
        statusCode: 500,
        statusMessage: e.message
      })
    }
  }
  // If an article is not found, we attempt to add it to the service based on doi
  try {
    const cratedArticle = await articlesSvr.CreateArticle(CreateArticleRequest.fromJSON({doi}))
    return await articlesSvr.GetArticle(GetArticleRequest.fromJSON({id: cratedArticle.id}))
  } catch (e) {
    return createError({
      statusCode: 404,
      statusMessage: "Article dose not exist"
    })
  }
  // Simulate API delay
  await new Promise(resolve => setTimeout(resolve, 1000))

  // Mock DOI lookup - in reality this would call CrossRef API
  if (doi.includes('10.')) {
    return {
      success: true,
      article: {
        doi,
        title: 'Advances in Machine Learning for Academic Research',
        authors: [
          { id: 1, name: 'Dr. Jane Smith', profileId: 1 },
          { id: 2, name: 'Prof. John Doe', profileId: 2 }
        ],
        publicationYear: 2024,
        journalName: 'Nature Machine Intelligence',
        abstract: 'This paper presents novel approaches to machine learning applications in academic research, focusing on automated literature analysis and citation pattern recognition.'
      }
    }
  } else {
    throw createError({
      statusCode: 404,
      statusMessage: 'Article not found for the provided DOI'
    })
  }
})
