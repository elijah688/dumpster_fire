import type { ProtoGrpcType } from '../items';
import type { Empty } from '../items/Empty';
const PROTO_PATH = '../items.proto';

import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import { onMount } from 'svelte';
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
    
const { items: { ItemsService, Empty } }: ProtoGrpcType = grpc.loadPackageDefinition(packageDefinition);

const client = new ItemsService(
    "localhost:5002",
    grpc.credentials.createInsecure())





export const GET = async (): any => {

    const ress = await new Promise((resolve,) => {
        client.GetAll(<Empty>{}, (err, res) => {
            resolve(res)
        })
    })
    return new Response(JSON.stringify(ress));

}




