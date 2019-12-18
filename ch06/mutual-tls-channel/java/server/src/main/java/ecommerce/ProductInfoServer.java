package ecommerce;

import io.grpc.Server;
import io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.NettyServerBuilder;
import io.netty.handler.ssl.ClientAuth;

import java.io.File;
import java.io.IOException;
import java.nio.file.Paths;
import java.util.logging.Logger;

public class ProductInfoServer {

    private static final Logger logger = Logger.getLogger(ProductInfoServer.class.getName());

    private Server server;

    private void start() throws IOException {
        File certFile = Paths.get("mutual-tls-channel", "certs", "server.crt").toFile();
        File keyFile = Paths.get("mutual-tls-channel", "certs", "server.pem").toFile();
        File caFile = Paths.get("mutual-tls-channel", "certs", "ca.crt").toFile();
        /* The port on which the server should run */
        int port = 50051;
        server = NettyServerBuilder.forPort(port)
                .addService(new ProductInfoImpl())
                .sslContext(GrpcSslContexts.forServer(certFile, keyFile)
                        .trustManager(caFile)
                        .clientAuth(ClientAuth.OPTIONAL)
                        .build())
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
        final ProductInfoServer server = new ProductInfoServer();
        server.start();
        server.blockUntilShutdown();
    }

}
