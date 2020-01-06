package ecommerce;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.opencensus.common.Duration;
import io.opencensus.contrib.grpc.metrics.RpcViews;

import java.io.IOException;
import java.util.logging.Logger;

import io.opencensus.exporter.stats.stackdriver.StackdriverStatsConfiguration;
import io.opencensus.exporter.stats.stackdriver.StackdriverStatsExporter;

public class ProductInfoServer {

    private static final Logger logger = Logger.getLogger(ProductInfoServer.class.getName());

    private Server server;

    private void start() throws IOException {
        /* The port on which the server should run */
        int port = 50051;
        server = ServerBuilder.forPort(port)
                .addService(new ProductInfoImpl())
                .build()
                .start();
        logger.info("Server started, listening on " + port);
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            // Use stderr here since the logger may have been reset by its JVM shutdown hook.
            logger.info("*** shutting down gRPC server since JVM is shutting down");
            ProductInfoServer.this.stop();
            logger.info("*** server shut down");
        }));
    }

    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    /**
     * Await termination on the main thread since the grpc library uses daemon threads.
     */
    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    /**
     * Main launches the server from the command line.
     */
    public static void main(String[] args) throws IOException, InterruptedException {
        // Register all the gRPC views and enable stats
        RpcViews.registerAllGrpcViews();

        // Create the Stackdriver stats exporter
        StackdriverStatsExporter.createAndRegister(
                StackdriverStatsConfiguration.builder()
                        .setProjectId("grpc-up-and-running-demo")
                        .setExportInterval(Duration.create(5, 0))
                        .build());

        final ProductInfoServer server = new ProductInfoServer();
        server.start();
        server.blockUntilShutdown();


    }

}
