import {useServices} from "~~/server/proto/useServices";
import {DeleteArticleRequest} from "~~/server/proto/grpc/articles/v1/article";

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
        
        const articlesSvr = await useServices().getArticlesServiceClient(event)
        
        await articlesSvr.DeleteArticle(DeleteArticleRequest.fromJSON({id}))
        
        return {
            success: true
        }
    } catch (error: any) {
        console.error("Error deleting article:", error)
        throw createError({
            statusCode: error.statusCode || 500,
            statusMessage: error.message || "Failed to delete article"
        })
    }
})
