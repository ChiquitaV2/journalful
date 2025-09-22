import {useServices} from "~~/server/proto/useServices";
import {GetArticleRequest} from "~~/server/proto/grpc/articles/v1/article";

export default defineEventHandler(async event => {
    try {
        const articlesSvr = await useServices().getArticlesServiceClient(event)
        const id = Number(event.context.params?.id)
        if (!id) {
            throw createError({
                statusCode: 400,
                statusMessage: 'ID is required'
            })
        }
        
        const response = await articlesSvr.GetArticle(GetArticleRequest.fromJSON({id}))
        return response.article
    } catch (error: any) {
        console.error("Error fetching article:", error)
        throw createError({
            statusCode: error.statusCode || 500,
            statusMessage: error.message || "Failed to fetch article"
        })
    }
})
