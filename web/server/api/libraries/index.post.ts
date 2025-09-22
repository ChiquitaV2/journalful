import {useServices} from "~~/server/proto/useServices";
import {CreateLibraryRequest} from "~~/server/proto/grpc/library/v1/library";

export default defineEventHandler(async event => {
  try {
    const session = await requireUserSession(event)
    const body = await readBody(event)
    const librarySvr = await useServices().getLibraryServiceClient(event)
    
    // Get user ID from session
    const userSub = session.id
    if (!userSub) {
      throw createError({
        statusCode: 401,
        statusMessage: "User not authenticated"
      })
    }
    
    // Get the user's profile to get the profile ID
    const profileSvr = await useServices().getProfileServiceClient(event)
    let profileId
    try {
      const profileResponse = await profileSvr.GetProfile({})
      profileId = profileResponse.profile?.id
      if (!profileId) {
        throw new Error("Profile ID not found")
      }
    } catch (profileError) {
      // If profile doesn't exist, we need to handle this case
      // For now, we'll return an error asking user to create a profile first
      throw createError({
        statusCode: 400,
        statusMessage: "User profile not found. Please complete your profile setup first."
      })
    }
    
    const { name, description, isPublic = false } = body
    
    if (!name) {
      throw createError({
        statusCode: 400,
        statusMessage: 'Library name is required'
      })
    }
    
    const result = await librarySvr.CreateLibrary(CreateLibraryRequest.fromJSON({
      ownerId: profileId,
      name,
      description: description || null,
      isPublic // This matches the proto field name is_public -> isPublic
    }))
    
    return {
      id: result.libraryId,
      name,
      description,
      isPublic,
      success: true
    }
  } catch (error: any) {
    console.error("library creation error", error)
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to create library"
    })
  }
})
