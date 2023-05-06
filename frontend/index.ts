import { ProtoGrpcType } from './items';
import { Empty } from './items/Empty';

var PROTO_PATH = __dirname + '/../items.proto';

var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });


var { items: { ItemsService, Empty } }: ProtoGrpcType = grpc.loadPackageDefinition(packageDefinition);
const client = new ItemsService(
    "localhost:5002",
    grpc.credentials.createInsecure())


client.GetAll(<Empty>{}, (err, res) => { console.log(res?.items) })





