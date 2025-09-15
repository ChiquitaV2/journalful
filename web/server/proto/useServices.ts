import {useGrpcClient} from "~~/server/proto/useGrpcClient";
import {H3Event} from "h3";
import {ArticlesServiceClientImpl} from "~~/server/proto/grpc/articles/v1/article";
import {LibraryServiceClientImpl} from "~~/server/proto/grpc/library/v1/library";

export const useServices = () => {
    const {createAuthenticatedClient} = useGrpcClient();

    const getArticlesServiceClient = async (event: H3Event) => {
        const accessToken = (await getUserSession(event)).secure?.accessToken
        //Create rpc client
        const client = createAuthenticatedClient(accessToken)
        return new ArticlesServiceClientImpl(client)
    }

    const getLibraryServiceClient = async (event: H3Event) => {
        const accessToken = (await getUserSession(event)).secure?.accessToken

        //Create rpc client
        const client = createAuthenticatedClient(accessToken)
        return new LibraryServiceClientImpl(client)
    }

    return {
        getArticlesServiceClient,
        getLibraryServiceClient
    }

}