import {useServices} from "~~/server/proto/useServices";
import {UpdateLibraryRequest} from "~~/server/proto/grpc/library/v1/library";

export default defineEventHandler(async event => {
  try {
    const id = Number(event.context.params?.id)
    const session = await requireUserSession(event)
    const body = await readBody(event)
    const librarySvr = await useServices().getLibraryServiceClient(event)
    
    if (!id) {
      throw createError({
        statusCode: 400,
        statusMessage: 'Library ID is required'
      })
    }
    
    const { name, description, isPublic } = body
    
    const result = await librarySvr.UpdateLibrary(UpdateLibraryRequest.fromJSON({
      libraryId: id,
      name: name || null,
      description: description || null,
      isPublic: isPublic !== undefined ? isPublic : null
    }))
    
    return {
      success: result.success,
      id
    }
  } catch (error: any) {
    console.error("library update error", error)
    if (error.message?.includes("not found")) {
      throw createError({
        statusCode: 404,
        statusMessage: "Library not found"
      })
    }
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to update library"
    })
  }
})
