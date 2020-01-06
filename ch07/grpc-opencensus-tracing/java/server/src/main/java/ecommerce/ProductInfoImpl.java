package ecommerce;

import io.grpc.Status;
import io.grpc.StatusException;
import io.opencensus.common.Scope;
import io.opencensus.trace.Tracer;
import io.opencensus.trace.Tracing;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class ProductInfoImpl extends ProductInfoGrpc.ProductInfoImplBase {

    private static final Tracer tracer = Tracing.getTracer();
    private Map productMap = new HashMap<String, ProductInfoOuterClass.Product>();

    @Override
    public void addProduct(ProductInfoOuterClass.Product request,
                           io.grpc.stub.StreamObserver<ProductInfoOuterClass.ProductID> responseObserver) {

        try (Scope ignored = ProductInfoImpl.tracer.spanBuilder("ecommerce.ProductInfoImpl.addProduct").startScopedSpan()) {
            UUID uuid = UUID.randomUUID();
            String randomUUIDString = uuid.toString();
            request = request.toBuilder().setId(randomUUIDString).build();
            productMap.put(randomUUIDString, request);
            ProductInfoOuterClass.ProductID id
                    = ProductInfoOuterClass.ProductID.newBuilder().setValue(randomUUIDString).build();
            responseObserver.onNext(id);
        } finally {
            responseObserver.onCompleted();
        }
    }

    @Override
    public void getProduct(ProductInfoOuterClass.ProductID request,
                           io.grpc.stub.StreamObserver<ProductInfoOuterClass.Product> responseObserver) {
        String id = request.getValue();
        if (productMap.containsKey(id)) {
            try (Scope ignored = ProductInfoImpl.tracer.spanBuilder("ecommerce.ProductInfoImpl.getProduct").startScopedSpan()) {
                responseObserver.onNext((ProductInfoOuterClass.Product) productMap.get(id));
            } finally {
                responseObserver.onCompleted();
            }
        } else {
            responseObserver.onError(new StatusException(Status.NOT_FOUND));
        }
    }
}
