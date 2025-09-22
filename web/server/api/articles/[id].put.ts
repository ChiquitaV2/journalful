import {useServices} from "~~/server/proto/useServices";
import {UpdateArticleRequest} from "~~/server/proto/grpc/articles/v1/article";

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
        const articlesSvr = await useServices().getArticlesServiceClient(event)
        
        const request = UpdateArticleRequest.fromJSON({
            id,
            doi: body.doi,
            title: body.title,
            abstract: body.abstract,
            publicationYear: body.publicationYear,
            journalName: body.journalName
        })
        
        await articlesSvr.UpdateArticle(request)
        
        return {
            success: true
        }
    } catch (error: any) {
        console.error("Error updating article:", error)
        throw createError({
            statusCode: error.statusCode || 500,
            statusMessage: error.message || "Failed to update article"
        })
    }
})
