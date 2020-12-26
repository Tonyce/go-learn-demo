
const grpc = require('@grpc/grpc-js');

const helloServices = require('./hello-world');

const server = new grpc.Server();

const GRPC_SERVER_ADDR = '127.0.0.1';
const GRPC_SERVER_PORT = 5051;

helloServices.addToServer(server);

server.bindAsync(`${GRPC_SERVER_ADDR}:${GRPC_SERVER_PORT}`, grpc.ServerCredentials.createInsecure(), (err, port) => {
    // if (err) return cb(err);
    server.start();
    // return cb(null, port);
}); 