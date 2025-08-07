import 'h3'

import {ArticlesService} from "./proto/grpc/articles/v1/article";
declare module "h3" {
    interface H3EventContext {
        articlesSvr: ArticlesService
    }
}
