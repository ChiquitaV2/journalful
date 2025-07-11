import {Client, credentials} from "@grpc/grpc-js";
import type {UnaryCallback} from "@grpc/grpc-js/build/src/client";

export interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

const conn = new Client(
    "localhost:50051",
    credentials.createInsecure()
);

type RpcImpl = (service: string, method: string, data: Uint8Array) => Promise<Uint8Array>;

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

export const rpc: Rpc = { request: sendRequest };