import {useServices} from "~~/server/proto/useServices";
import {UpdateProfileRequest} from "~~/server/proto/grpc/profile/v1/profile";

export default defineEventHandler(async event => {
  try {
    await requireUserSession(event)
    
    const id = Number(event.context.params?.id)
    if (!id) {
      throw createError({
        statusCode: 400,
        statusMessage: 'ID is required'
      })
    }
    
    const body = await readBody(event)
    const profileSvr = await useServices().getProfileServiceClient(event)
    
    const request = UpdateProfileRequest.fromJSON({
      id,
      name: body.name,
      bio: body.bio,
      institution: body.institution
    })
    
    await profileSvr.UpdateProfile(request)
    
    return {
      success: true
    }
  } catch (error: any) {
    console.error("Error updating profile:", error)
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to update profile"
    })
  }
})
