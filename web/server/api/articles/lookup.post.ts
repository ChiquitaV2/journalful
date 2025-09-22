import {useServices} from "~~/server/proto/useServices";
import {
  CreateArticleRequest,
  GetArticleByDOIRequest,
  GetArticleRequest
} from "~~/server/proto/grpc/articles/v1/article";

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event)
    const { doi } = body
    
    const session = await requireUserSession(event)

    if (!doi) {
      throw createError({
        statusCode: 400,
        statusMessage: 'DOI is required'
      })
    }

    const articlesSvr = await useServices().getArticlesServiceClient(event)
    
    // Try to find existing article by DOI
    try {
      const response = await articlesSvr.GetArticleByDOI(GetArticleByDOIRequest.fromJSON({doi}))
      return {
        success: true,
        article: response.article
      }
    } catch (e: any) {
      console.error("lookup error", e)
      // For all errors besides not found, return the error
      if (!e.message.includes("not found")) {
        throw createError({
          statusCode: 500,
          statusMessage: e.message
        })
      }
    }
    
    // If article is not found, attempt to create it based on DOI
    try {
      const createdArticle = await articlesSvr.CreateArticle(CreateArticleRequest.fromJSON({
        doi,
        title: `Article with DOI: ${doi}`, // Backend should fetch metadata
        authors: []
      }))
      
      const response = await articlesSvr.GetArticle(GetArticleRequest.fromJSON({id: createdArticle.id}))
      return {
        success: true,
        article: response.article
      }
    } catch (e: any) {
      console.error("creation error", e)
      throw createError({
        statusCode: 404,
        statusMessage: "Article does not exist and could not be created"
      })
    }
  } catch (error: any) {
    console.error("Article lookup error:", error)
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to lookup article"
    })
  }
})
