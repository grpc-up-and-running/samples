var PROTO_PATH = __dirname + '/../../proto/product_info.proto';
var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var uuidv4 = require('uuid/v4');

let productMap = new Map();

// Suggested options for similarity to existing grpc.load behavior
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);

// The protoDescriptor object has the full package hierarchy
var productInfo = protoDescriptor.ecommerce;

function addProduct(call, callback) {
    var newProduct = call.request;
    var id = uuidv4();
    newProduct.id = id;
    productMap.set(id, newProduct);
    // Constructing the output
    var productID = {
        value: id,
    };
    callback(null, productID);
}

function getProduct(call, callback) {
    var productID = call.request;
    if (productMap.has(productID.value)) {
        callback(null, productMap.get(productID.value));
    } else {
        callback({
            code: 400,
            message: "product doesn't exist for the product id: " + productID.value,
            status: grpc.status.NOT_FOUND
        })
    }
}

function getServer() {
    var server = new grpc.Server();
    server.addService(productInfo.ProductInfo.service, {
        addProduct: addProduct,
        getProduct: getProduct
    });
    return server;
}

var productInfoServer = getServer();
productInfoServer.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure());
productInfoServer.start();