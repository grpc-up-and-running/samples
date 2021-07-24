import ballerina/grpc;

public isolated client class ProductInfoClient {
    *grpc:AbstractClientEndpoint;

    private final grpc:Client grpcClient;

    public isolated function init(string url, *grpc:ClientConfiguration config) returns grpc:Error? {
        self.grpcClient = check new (url, config);
        check self.grpcClient.initStub(self, ROOT_DESCRIPTOR, getDescriptorMap());
    }

    isolated remote function addProduct(Product|ContextProduct req) returns (ProductID|grpc:Error) {
        map<string|string[]> headers = {};
        Product message;
        if (req is ContextProduct) {
            message = req.content;
            headers = req.headers;
        } else {
            message = req;
        }
        var payload = check self.grpcClient->executeSimpleRPC("ecommerce.ProductInfo/addProduct", message, headers);
        [anydata, map<string|string[]>] [result, _] = payload;
        return <ProductID>result;
    }

    isolated remote function addProductContext(Product|ContextProduct req) returns (ContextProductID|grpc:Error) {
        map<string|string[]> headers = {};
        Product message;
        if (req is ContextProduct) {
            message = req.content;
            headers = req.headers;
        } else {
            message = req;
        }
        var payload = check self.grpcClient->executeSimpleRPC("ecommerce.ProductInfo/addProduct", message, headers);
        [anydata, map<string|string[]>] [result, respHeaders] = payload;
        return {content: <ProductID>result, headers: respHeaders};
    }

    isolated remote function getProduct(ProductID|ContextProductID req) returns (Product|grpc:Error) {
        map<string|string[]> headers = {};
        ProductID message;
        if (req is ContextProductID) {
            message = req.content;
            headers = req.headers;
        } else {
            message = req;
        }
        var payload = check self.grpcClient->executeSimpleRPC("ecommerce.ProductInfo/getProduct", message, headers);
        [anydata, map<string|string[]>] [result, _] = payload;
        return <Product>result;
    }

    isolated remote function getProductContext(ProductID|ContextProductID req) returns (ContextProduct|grpc:Error) {
        map<string|string[]> headers = {};
        ProductID message;
        if (req is ContextProductID) {
            message = req.content;
            headers = req.headers;
        } else {
            message = req;
        }
        var payload = check self.grpcClient->executeSimpleRPC("ecommerce.ProductInfo/getProduct", message, headers);
        [anydata, map<string|string[]>] [result, respHeaders] = payload;
        return {content: <Product>result, headers: respHeaders};
    }
}

public client class ProductInfoProductIDCaller {
    private grpc:Caller caller;

    public isolated function init(grpc:Caller caller) {
        self.caller = caller;
    }

    public isolated function getId() returns int {
        return self.caller.getId();
    }

    isolated remote function sendProductID(ProductID response) returns grpc:Error? {
        return self.caller->send(response);
    }

    isolated remote function sendContextProductID(ContextProductID response) returns grpc:Error? {
        return self.caller->send(response);
    }

    isolated remote function sendError(grpc:Error response) returns grpc:Error? {
        return self.caller->sendError(response);
    }

    isolated remote function complete() returns grpc:Error? {
        return self.caller->complete();
    }

    public isolated function isCancelled() returns boolean {
        return self.caller.isCancelled();
    }
}

public client class ProductInfoProductCaller {
    private grpc:Caller caller;

    public isolated function init(grpc:Caller caller) {
        self.caller = caller;
    }

    public isolated function getId() returns int {
        return self.caller.getId();
    }

    isolated remote function sendProduct(Product response) returns grpc:Error? {
        return self.caller->send(response);
    }

    isolated remote function sendContextProduct(ContextProduct response) returns grpc:Error? {
        return self.caller->send(response);
    }

    isolated remote function sendError(grpc:Error response) returns grpc:Error? {
        return self.caller->sendError(response);
    }

    isolated remote function complete() returns grpc:Error? {
        return self.caller->complete();
    }

    public isolated function isCancelled() returns boolean {
        return self.caller.isCancelled();
    }
}

public type ContextProduct record {|
    Product content;
    map<string|string[]> headers;
|};

public type ContextProductID record {|
    ProductID content;
    map<string|string[]> headers;
|};

public type Product record {|
    string id = "";
    string name = "";
    string description = "";
    float price = 0.0;
|};

public type ProductID record {|
    string value = "";
|};

const string ROOT_DESCRIPTOR = "0A1270726F647563745F696E666F2E70726F746F120965636F6D6D6572636522650A0750726F64756374120E0A0269641801200128095202696412120A046E616D6518022001280952046E616D6512200A0B6465736372697074696F6E180320012809520B6465736372697074696F6E12140A0570726963651804200128025205707269636522210A0950726F64756374494412140A0576616C7565180120012809520576616C7565327D0A0B50726F64756374496E666F12360A0A61646450726F6475637412122E65636F6D6D657263652E50726F647563741A142E65636F6D6D657263652E50726F64756374494412360A0A67657450726F6475637412142E65636F6D6D657263652E50726F6475637449441A122E65636F6D6D657263652E50726F64756374620670726F746F33";

isolated function getDescriptorMap() returns map<string> {
    return {"product_info.proto": "0A1270726F647563745F696E666F2E70726F746F120965636F6D6D6572636522650A0750726F64756374120E0A0269641801200128095202696412120A046E616D6518022001280952046E616D6512200A0B6465736372697074696F6E180320012809520B6465736372697074696F6E12140A0570726963651804200128025205707269636522210A0950726F64756374494412140A0576616C7565180120012809520576616C7565327D0A0B50726F64756374496E666F12360A0A61646450726F6475637412122E65636F6D6D657263652E50726F647563741A142E65636F6D6D657263652E50726F64756374494412360A0A67657450726F6475637412142E65636F6D6D657263652E50726F6475637449441A122E65636F6D6D657263652E50726F64756374620670726F746F33"};
}

