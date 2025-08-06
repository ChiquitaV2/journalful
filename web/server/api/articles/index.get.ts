export default defineEventHandler(event => {
    return event.context.articlesSvr.ListArticles({})
})