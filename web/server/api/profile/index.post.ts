import {useServices} from "~~/server/proto/useServices";
import {CreateProfileRequest} from "~~/server/proto/grpc/profile/v1/profile";
import {CreateLibraryRequest} from "~~/server/proto/grpc/library/v1/library";

export default defineEventHandler(async event => {
  try {
    const session = await requireUserSession(event)
    
    const body = await readBody(event)
    const profileSvr = await useServices().getProfileServiceClient(event)
    const librarySvr = await useServices().getLibraryServiceClient(event)
    
    // Use user's name from session if not provided
    const sessionUser = session.user as any // Type assertion for auth provider fields
    const userName = body.name || sessionUser?.name || sessionUser?.given_name || 'User'
    
    if (!userName.trim()) {
      throw createError({
        statusCode: 400,
        statusMessage: 'Name is required'
      })
    }
    
    // Create the profile
    const request = CreateProfileRequest.fromJSON({
      name: userName.trim(),
      bio: body.bio || null,
      institution: body.institution || null
    })
    
    const profileResponse = await profileSvr.CreateProfile(request)
    
    // Create a default library for the new user
    try {
      const defaultLibraryResult = await librarySvr.CreateLibrary(CreateLibraryRequest.fromJSON({
        ownerId: profileResponse.id,
        name: "My Reading List",
        description: "Your personal academic reading collection",
        isPublic: false
      }))
      
      console.log(`Created default library ${defaultLibraryResult.libraryId} for new user profile ${profileResponse.id}`)
    } catch (libraryError) {
      console.error("Failed to create default library for new user:", libraryError)
      // Don't fail the entire request if library creation fails
    }
    
    return {
      success: true,
      id: profileResponse.id,
      name: userName.trim(),
      bio: body.bio || null,
      institution: body.institution || null,
      setupComplete: true
    }
  } catch (error: any) {
    console.error("Error creating profile:", error)
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to create profile"
    })
  }
})
