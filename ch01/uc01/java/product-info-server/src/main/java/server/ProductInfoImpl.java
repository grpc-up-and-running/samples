package server;

import ecommerce.ProductInfoGrpc;
import ecommerce.ProductInfoOuterClass;
import io.grpc.Status;
import io.grpc.StatusException;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class ProductInfoImpl extends ProductInfoGrpc.ProductInfoImplBase {

    private Map<String, ProductInfoOuterClass.Product> productMap = new HashMap();

    @Override
    public void addProduct(ProductInfoOuterClass.Product request, io.grpc.stub.StreamObserver<ProductInfoOuterClass.ProductID> responseObserver) {
        UUID uuid = UUID.randomUUID();
        String randomUUIDString = uuid.toString();
        productMap.put(randomUUIDString, request);
        ProductInfoOuterClass.ProductID id =
                ProductInfoOuterClass.ProductID.newBuilder().setValue(randomUUIDString).build();
        responseObserver.onNext(id);
        responseObserver.onCompleted();
    }

    @Override
    public void getProduct(ProductInfoOuterClass.ProductID request, io.grpc.stub.StreamObserver<ProductInfoOuterClass.Product> responseObserver) {
        String id = request.getValue();
        if (productMap.containsKey(id)) {
            responseObserver.onNext(productMap.get(id));
        } else {
            responseObserver.onError(new StatusException(Status.NOT_FOUND));
        }
        super.getProduct(request, responseObserver);
    }
}
