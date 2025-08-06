import {Client, credentials} from "@grpc/grpc-js";
import type {UnaryCallback} from "@grpc/grpc-js/build/src/client";
import {ArticlesServiceClientImpl} from "~~/server/proto/grpc/articles/v1/article";

export interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
type RpcImpl = (service: string, method: string, data: Uint8Array) => Promise<Uint8Array>;

export default defineNitroPlugin(nitroApp => {
    const config = useRuntimeConfig()
    const conn = new Client(
        config.backendURL,
        credentials.createInsecure()
    );

    const sendRequest: RpcImpl = (service, method, data) => {
        // Conventionally in gRPC, the request path looks like
        //   "package.names.ServiceName/MethodName",
        // we therefore construct such a string
        const path = `/${service}/${method}`;

        return new Promise((resolve, reject) => {
            // makeUnaryRequest transmits the result (and error) with a callback
            // transform this into a promise!
            const resultCallback: UnaryCallback<any> = (err, res) => {
                if (err) {
                    return reject(err);
                }
                resolve(res);
            };

            function passThrough(argument: any) {
                return argument;
            }

            // Using passThrough as the deserialize functions
            conn.makeUnaryRequest(path, d => Buffer.from(d), passThrough, data, resultCallback);
        });
    };
    const rpc: Rpc = { request: sendRequest };

    const articlesSvr = new ArticlesServiceClientImpl(rpc)
    nitroApp.hooks.hook('request', (event) => {
        event.context.articlesSvr = articlesSvr
    })
})
