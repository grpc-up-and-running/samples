import ballerina/grpc;
import ballerina/uuid;

@grpc:ServiceDescriptor {descriptor: ROOT_DESCRIPTOR, descMap: getDescriptorMap()}
service "ProductInfo" on new grpc:Listener(9090) {

    private final map<Product> productMap = {};

    remote function addProduct(Product value) returns ProductID|error {
        string uuid = uuid:createType1AsString();
        value.id = uuid;
        lock {
            self.productMap[uuid] = value.cloneReadOnly();
        }
        return {value: uuid};
    }

    remote function getProduct(ProductID productId) returns Product|error {
        lock {
            if (self.productMap.hasKey(productId.value)) {
                return self.productMap.get(productId.value).cloneReadOnly();
            } else {
                return error grpc:NotFoundError("product doesn't exist for the product id: " + productId.value);
            }
        }
    }
}

