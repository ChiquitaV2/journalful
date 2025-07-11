export default defineEventHandler(event => {
  const articleSvr = event.context.articleSvr
  const res = articleSvr.ListArticles({})
  return res
})
