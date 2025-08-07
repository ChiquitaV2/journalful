import {Client, credentials, Metadata, CallCredentials} from "@grpc/grpc-js";
import type {UnaryCallback} from "@grpc/grpc-js/build/src/client";

export interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
    handleError?(service: string, method: string, error: Error): Error;

}

export const useGrpcClient = () => {
    const config = useRuntimeConfig();

    //
    const createUserCallCredentials = (token: string): CallCredentials => {
        return credentials.createFromMetadataGenerator((options, callback) => {
            const metadata = new Metadata();
            metadata.add('authorization', `Bearer ${token}`);
            callback(null, metadata);
        });
    };

    const createAuthenticatedClient = (token?: string): Rpc => {
        let channelCredentials;
        if (token) {
            const sslCreds = credentials.createSsl();
            const userCallCreds = createUserCallCredentials(token);
            channelCredentials = credentials.combineChannelCredentials(sslCreds, userCallCreds);
        } else {
            channelCredentials = credentials.createInsecure();
        }
        const conn = new Client(config.backendURL, channelCredentials);

        const sendRequest = (service: string, method: string, data: Uint8Array): Promise<Uint8Array> => {
            const path = `/${service}/${method}`;

            return new Promise((resolve, reject) => {
                const resultCallback: UnaryCallback<any> = (err, res) => {
                    if (err) {
                        return reject(err);
                    }
                    resolve(res);
                };

                function passThrough(argument: any) {
                    return argument;
                }

                conn.makeUnaryRequest(path, d => Buffer.from(d), passThrough, data, resultCallback);
            });
        };

        const handleError: Rpc['handleError'] = (service, method, error) => {
            console.error("Serivce: ", service, "Method: ", method, "Error:", error.message)
            if (error.message.includes("not found")) {
                return createError({
                    statusCode: 404,
                    statusMessage: error.message
                })
            } else {
                return error
            }
        }

        return { request: sendRequest, handleError };
    }

    return {
        createAuthenticatedClient
    }
}