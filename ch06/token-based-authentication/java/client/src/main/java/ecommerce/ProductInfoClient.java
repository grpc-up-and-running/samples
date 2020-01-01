package ecommerce;

import io.grpc.ManagedChannel;
import io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.NettyChannelBuilder;
import io.netty.handler.ssl.SslContext;

import java.io.File;
import java.nio.file.Paths;
import java.util.logging.Logger;
import javax.net.ssl.SSLException;

/**
 * gRPC client sample for productInfo service.
 */
public class ProductInfoClient {

    private static final Logger logger = Logger.getLogger(ProductInfoClient.class.getName());

    public static void main(String[] args) throws SSLException {
        File certFile = Paths.get("basic-authentication", "certs", "server.crt").toFile();
        SslContext sslContext = GrpcSslContexts.forClient().trustManager(certFile).build();
        ManagedChannel channel = NettyChannelBuilder.forAddress("localhost", 50051)
                .useTransportSecurity()
                .sslContext(sslContext)
                .build();

        TokenCallCredentials callCredentials = new TokenCallCredentials("some-secret-token");

        ProductInfoGrpc.ProductInfoBlockingStub stub =
                ProductInfoGrpc.newBlockingStub(channel).withCallCredentials(callCredentials);

        ProductInfoOuterClass.ProductID productID = stub.addProduct(
                ProductInfoOuterClass.Product.newBuilder()
                        .setName("Samsung S10")
                        .setDescription("Samsung Galaxy S10 is the latest smart phone, " +
                                "launched in February 2019")
                        .setPrice(700.0f)
                        .build());
        logger.info("Product ID: " + productID.getValue() + " added successfully.");

        ProductInfoOuterClass.Product product = stub.getProduct(productID);
        logger.info("Product: " + product.toString());
        channel.shutdown();
    }
}
