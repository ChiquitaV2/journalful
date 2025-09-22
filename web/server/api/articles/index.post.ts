import {useServices} from "~~/server/proto/useServices";
import {CreateArticleRequest} from "~~/server/proto/grpc/articles/v1/article";

export default defineEventHandler(async event => {
    try {
        await requireUserSession(event)
        
        const body = await readBody(event)
        const articlesSvr = await useServices().getArticlesServiceClient(event)
        
        // Validate required fields
        if (!body.doi || !body.title) {
            throw createError({
                statusCode: 400,
                statusMessage: 'DOI and title are required'
            })
        }
        
        // Transform authors from strings to Author objects
        const authorObjects = (body.authors || []).map((authorName: string) => ({
            id: 0, // Will be assigned by backend
            name: authorName,
            profileId: 0 // No profile linked initially
        }))
        
        const request = CreateArticleRequest.fromJSON({
            doi: body.doi,
            title: body.title,
            authors: authorObjects,
            abstract: body.abstract,
            publicationYear: body.publicationYear,
            journalName: body.journalName
        })
        
        const response = await articlesSvr.CreateArticle(request)
        
        return {
            success: true,
            id: response.id
        }
    } catch (error: any) {
        console.error("Error creating article:", error)
        throw createError({
            statusCode: error.statusCode || 500,
            statusMessage: error.message || "Failed to create article"
        })
    }
})
