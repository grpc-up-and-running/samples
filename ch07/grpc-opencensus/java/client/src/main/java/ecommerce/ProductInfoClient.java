package ecommerce;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.opencensus.common.Duration;
import io.opencensus.contrib.grpc.metrics.RpcViews;
import io.opencensus.exporter.stats.stackdriver.StackdriverStatsConfiguration;
import io.opencensus.exporter.stats.stackdriver.StackdriverStatsExporter;

import java.io.IOException;
import java.util.logging.Logger;

/**
 * gRPC client sample for productInfo service.
 */
public class ProductInfoClient {

    private static final Logger logger = Logger.getLogger(ProductInfoClient.class.getName());

    public static void main(String[] args) throws IOException, InterruptedException {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 50051)
                .usePlaintext()
                .build();
        setupOpencensusAndExporters();
        ProductInfoGrpc.ProductInfoBlockingStub stub =
                ProductInfoGrpc.newBlockingStub(channel);

        for(int i = 0; i < 10000; i++) {
            Thread.sleep(1000);
            ProductInfoOuterClass.ProductID productID = stub.addProduct(
                    ProductInfoOuterClass.Product.newBuilder()
                            .setName("Samsung S10 " + i)
                            .setDescription("Samsung Galaxy S10 is the latest smart phone, " +
                                    "launched in February 2019")
                            .setPrice(700.0f)
                            .build());
            logger.info("Product ID: " + productID.getValue() + " added successfully.");

            ProductInfoOuterClass.Product product = stub.getProduct(productID);
            logger.info("Product: " + product.toString());
        }
        channel.shutdown();
    }

    private static void setupOpencensusAndExporters() throws IOException {

        // Register all the gRPC views and enable stats
        RpcViews.registerAllGrpcViews();

        // Create the Stackdriver stats exporter
        StackdriverStatsExporter.createAndRegister(
                StackdriverStatsConfiguration.builder()
                        .setProjectId("grpc-up-and-running-demo")
                        .setExportInterval(Duration.create(5, 0))
                        .build());
    }
}
