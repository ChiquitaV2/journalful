import {generateMockLibrary} from '../mock/data'
import {useServices} from "~~/server/proto/useServices";
import {GetUserLibraryRequest, ReadingStatus} from "~~/server/proto/grpc/library/v1/library";

export default defineEventHandler(async event => {
    try {
        const session = await requireUserSession(event)
        console.log('Session:', session)

        const librarySvr = await useServices().getLibraryServiceClient(event)

        // Get user ID from session - we need to get the profile ID from the user_id
        const userSub = session.id
        if (!userSub) {
            return createError({
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
            return createError({
                statusCode: 400,
                statusMessage: "User profile not found. Please complete your profile setup first."
            })
        }

        const libraries = await librarySvr.GetUserLibrary(GetUserLibraryRequest.fromJSON({userId: profileId}))
        const defaultLibrary = libraries.defaultLibrary
        const privateLibraries = libraries.privateLibraries

        if (!defaultLibrary) {
            return createError({
                statusCode: 404,
                statusMessage: "Default library not found"
            })
        }

        return {
            defaultLibrary,
            privateLibraries,
            stats: {
                total: 1 + privateLibraries.length,
                articles: defaultLibrary.articles.length + privateLibraries.reduce((sum, lib) => sum + lib.articles.length, 0),
                completed: [defaultLibrary, ...privateLibraries]
                    .flatMap(lib => lib.articles)
                    .filter(article => article.readingStatus === ReadingStatus.READING_STATUS_READ).length
            }
        }
    } catch (e: any) {
        console.error("library lookup error", e)
        return createError({
            statusCode: 500,
            statusMessage: e.message || "Failed to fetch libraries"
        })
    }
})
