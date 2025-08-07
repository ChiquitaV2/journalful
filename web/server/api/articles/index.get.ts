import {useServices} from "~~/server/proto/useServices";

export default defineEventHandler(async event => {
    const articlesSvr = await useServices().getArticlesServiceClient(event)
    return articlesSvr.ListArticles({})
})