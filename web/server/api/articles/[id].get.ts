import {useServices} from "~~/server/proto/useServices";
import {GetArticleRequest} from "~~/server/proto/grpc/articles/v1/article";

export default defineEventHandler(async event => {
  const articlesSvr = await useServices().getArticlesServiceClient(event)
  const id = Number(event.context.params?.id)
  if (!id) createError({
    statusCode: 400,
    statusMessage: 'ID is required'
  })
  return await articlesSvr.GetArticle(GetArticleRequest.fromJSON({id}))

})
