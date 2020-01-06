package ecommerce;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.opencensus.common.Scope;
import io.opencensus.exporter.trace.stackdriver.StackdriverTraceConfiguration;
import io.opencensus.exporter.trace.stackdriver.StackdriverTraceExporter;
import io.opencensus.trace.Tracer;
import io.opencensus.trace.Tracing;
import io.opencensus.trace.config.TraceConfig;
import io.opencensus.trace.samplers.Samplers;

import java.io.IOException;
import java.util.logging.Logger;

/**
 * gRPC client sample for productInfo service.
 */
public class ProductInfoClient {

    private static final Tracer tracer = Tracing.getTracer();
    private static final Logger logger = Logger.getLogger(ProductInfoClient.class.getName());

    public static void main(String[] args) throws IOException, InterruptedException {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 50051)
                .usePlaintext()
                .build();
        setupOpencensusAndExporters();
        ProductInfoGrpc.ProductInfoBlockingStub stub =
                ProductInfoGrpc.newBlockingStub(channel);
        ProductInfoOuterClass.ProductID productID;
        try (Scope ignored = ProductInfoClient.tracer.spanBuilder("ecommerce.ProductInfoClient.addProduct").startScopedSpan()) {
            productID = stub.addProduct(
                    ProductInfoOuterClass.Product.newBuilder()
                            .setName("Samsung S10 ")
                            .setDescription("Samsung Galaxy S10 is the latest smart phone, " +
                                    "launched in February 2019")
                            .setPrice(700.0f)
                            .build());
            logger.info("Product ID: " + productID.getValue() + " added successfully.");
        }

        try (Scope ignored =
                     ProductInfoClient.tracer.spanBuilder("ecommerce.ProductInfoClient.getProduct").startScopedSpan()) {
            ProductInfoOuterClass.Product product = stub.getProduct(productID);
            logger.info("Product: " + product.toString());
        }

        channel.shutdown();
    }

    private static void setupOpencensusAndExporters() throws IOException {

        // For demo purposes, always sample
        TraceConfig traceConfig = Tracing.getTraceConfig();
        traceConfig.updateActiveTraceParams(
                traceConfig.getActiveTraceParams()
                        .toBuilder()
                        .setSampler(Samplers.alwaysSample())
                        .build());

        // Create the Stackdriver trace exporter
        StackdriverTraceExporter.createAndRegister(
                StackdriverTraceConfiguration.builder()
                        .setProjectId("grpc-up-and-running-demo")
                        .build());
    }
}
