import {useGrpcClient} from "~~/server/proto/useGrpcClient";
import {H3Event} from "h3";
import {ArticlesServiceClientImpl} from "~~/server/proto/grpc/articles/v1/article";
import {LibraryServiceClientImpl} from "~~/server/proto/grpc/library/v1/library";
import {ProfileServiceClientImpl} from "~~/server/proto/grpc/profile/v1/profile";
import {AuthorServiceClientImpl} from "~~/server/proto/grpc/profile/v1/author";

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

    const getProfileServiceClient = async (event: H3Event) => {
        const accessToken = (await getUserSession(event)).secure?.accessToken

        //Create rpc client
        const client = createAuthenticatedClient(accessToken)
        return new ProfileServiceClientImpl(client)
    }

    const getAuthorServiceClient = async (event: H3Event) => {
        const accessToken = (await getUserSession(event)).secure?.accessToken

        //Create rpc client
        const client = createAuthenticatedClient(accessToken)
        return new AuthorServiceClientImpl(client)
    }

    return {
        getArticlesServiceClient,
        getLibraryServiceClient,
        getProfileServiceClient,
        getAuthorServiceClient
    }

}