import type {ArticlesService} from "~/shared/grpc/articles/v1/article";
declare module "h3" {

    interface H3EventContext {
        articleSvr: ArticlesService
    }
}