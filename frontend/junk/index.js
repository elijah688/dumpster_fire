var PROTO_PATH = __dirname + '/../items.proto';

// var parseArgs = require('minimist');
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
var hello_proto = grpc.loadPackageDefinition(packageDefinition);

function main() {
    const client = new hello_proto.items.ItemsService(
        "localhost:5002",
        grpc.credentials.createInsecure()
    )
    client.GetAll(undefined, (err, res) => { console.log(err, res) })
}

main();
