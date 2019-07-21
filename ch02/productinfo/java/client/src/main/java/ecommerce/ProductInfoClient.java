package ecommerce;

import com.google.protobuf.StringValue;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

/**
 * gRPC client sample for productInfo service.
 */
public class ProductInfoClient {

    public static void main(String[] args) throws InterruptedException {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 50051)
                .usePlaintext(true)
                .build();

        ProductInfoGrpc.ProductInfoBlockingStub stub =
                ProductInfoGrpc.newBlockingStub(channel);

        StringValue productID = stub.addProduct(
                ProductInfoOuterClass.Product.newBuilder()
                        .setName("Sumsung S10")
                        .setDescription("Samsung Galaxy S10 is the latest smart phone, launched in February 2019")
                        .setPrice(700.0f)
                        .build());
        System.out.println(productID.getValue());

        ProductInfoOuterClass.Product product = stub.getProduct(productID);
        System.out.println(product.toString());
        channel.shutdown();
    }
}
