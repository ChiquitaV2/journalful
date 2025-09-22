import {useServices} from "~~/server/proto/useServices";
import {GetProfileRequest} from "~~/server/proto/grpc/profile/v1/profile";

export default defineEventHandler(async event => {
  try {
    await requireUserSession(event)
    
    const profileSvr = await useServices().getProfileServiceClient(event)
    const response = await profileSvr.GetProfile(GetProfileRequest.fromJSON({}))
    
    return response.profile
  } catch (error: any) {
    console.error("Error fetching profile:", error)
    if (error.message?.includes("not found")) {
      throw createError({
        statusCode: 404,
        statusMessage: "Profile not found"
      })
    }
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to fetch profile"
    })
  }
})