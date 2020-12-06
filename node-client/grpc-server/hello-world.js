/**
 * hello world 
 * 验证程序跑起来了
 */
const path = require('path');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const helloworldProtoFileName = path.resolve(__dirname, '../../proto/helloworld.proto');
const helloworldPackageDefinition = protoLoader.loadSync(helloworldProtoFileName, {});
const helloworldObject = grpc.loadPackageDefinition(helloworldPackageDefinition);

// console.log(helloworldObject)

const Greater = helloworldObject.pb.Greeter;

function addToServer(server) {
    server.addService(Greater.service, {
        sayHello: (call, callback) => {
            const { metadata, request } = call; // {Http2ServerCallStream, cancelled:false}
            call.sendMetadata(metadata);
            const { name } = request;
            callback(null, { message: `hello 5051 ${name}` });
        },
    });
}

exports.addToServer = addToServer;