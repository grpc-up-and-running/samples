var PROTO_PATH = __dirname + '/../../proto/product_info.proto';
var grpc = require('grpc');
var async = require('async');
var protoLoader = require('@grpc/proto-loader');

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
var client = new productInfo.ProductInfo('localhost:50051',
    grpc.credentials.createInsecure());
let productID = 0;

function addProduct(cb) {
    client.addProduct({
        name: "Apple iPhone 11",
        description: "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode.",
        price: 699.0
    }, (error, response) => {
        if (!error) {
            console.log("addProduct - Response : ", response);
            productID = response.value;
            return cb(null, productID);
        } else {
            console.log("addProduct - Error:", error.message);
            return cb(error);
        }
    });
}

function getProduct(cb) {
    client.getProduct({
        value: productID
    }, (error, response) => {
        if (!error) {
            console.log("getProduct - Response : ", response);
            return cb(null, response);
        } else {
            console.log("getProduct - Error:", error.message);
            return cb(error);
        }
    });
}

function main() {
    async.series([
        addProduct,
        getProduct
    ]);
}

if (require.main === module) {
    main();
}

exports.addProduct = addProduct;
exports.getProduct = getProduct;