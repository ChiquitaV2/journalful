import {rpc} from "#shared/grpc/rpc";
import {ArticlesServiceClientImpl} from "#shared/grpc/articles/v1/article";

export default defineNitroPlugin(nitroApp => {

    const articleSvr = new ArticlesServiceClientImpl(rpc)

    nitroApp.hooks.hook('request', (event) => {
        event.context.articleSvr = articleSvr
        })
})
