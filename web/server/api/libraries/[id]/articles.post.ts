import {useServices} from "~~/server/proto/useServices";
import {SaveArticleToLibraryRequest, ReadingStatus} from "~~/server/proto/grpc/library/v1/library";

export default defineEventHandler(async (event) => {
  try {
    await requireUserSession(event)
    
    const body = await readBody(event)
    const libraryId = Number(event.context.params?.id)
    
    if (!libraryId || body.articleId === undefined) {
      throw createError({
        statusCode: 400,
        statusMessage: 'Library ID and article ID are required'
      })
    }
    
    const librarySvr = await useServices().getLibraryServiceClient(event)
    
    const result = await librarySvr.SaveArticleToLibrary(SaveArticleToLibraryRequest.fromJSON({
      libraryId,
      articleId: body.articleId,
      readingStatus: body.readingStatus || ReadingStatus.READING_STATUS_TO_READ,
      notes: body.notes
    }))
    
    return {
      success: true,
      id: result.id
    }
  } catch (error: any) {
    console.error("Error adding article to library:", error)
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to add article to library"
    })
  }
});