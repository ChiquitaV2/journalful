declare module "h3" {

    import {ArticlesService} from "../server/proto/grpc/articles/v1/article";

    interface H3EventContext {
        articlesSvr: ArticlesService
    }
}