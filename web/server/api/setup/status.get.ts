import {useServices} from "~~/server/proto/useServices";
import {GetProfileRequest} from "~~/server/proto/grpc/profile/v1/profile";

export default defineEventHandler(async event => {
  try {
    const session = await requireUserSession(event)
    
    const profileSvr = await useServices().getProfileServiceClient(event)
    
    try {
      const response = await profileSvr.GetProfile(GetProfileRequest.fromJSON({}))
      
      return {
        needsSetup: false,
        hasProfile: true,
        profile: response.profile
      }
    } catch (error: any) {
      // If profile doesn't exist, user needs setup
      if (error.message?.includes("not found")) {
        return {
          needsSetup: true,
          hasProfile: false,
          profile: null
        }
      }
      
      // For other errors, assume they don't need setup to avoid redirect loops
      return {
        needsSetup: false,
        hasProfile: false,
        error: error.message
      }
    }
  } catch (error: any) {
    console.error("Setup status check failed:", error)
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.message || "Failed to check setup status"
    })
  }
})