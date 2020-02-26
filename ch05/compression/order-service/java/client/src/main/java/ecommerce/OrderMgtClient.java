package ecommerce;

import com.google.protobuf.StringValue;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.util.logging.Logger;

public class OrderMgtClient {

    private static final Logger logger = Logger.getLogger(OrderMgtClient.class.getName());

    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress(
                "localhost", 50051).usePlaintext().build();
        OrderManagementGrpc.OrderManagementBlockingStub stub = OrderManagementGrpc.newBlockingStub(channel);

        OrderManagementOuterClass.Order order = OrderManagementOuterClass.Order
                .newBuilder()
                .setId("101")
                .addItems("iPhone XS").addItems("Mac Book Pro")
                .setDestination("San Jose, CA")
                .setPrice(2300)
                .build();

        // Add Order
        StringValue result = stub.withCompression("gzip").addOrder(order);
        logger.info("AddOrder Response -> : " + result.getValue());

        // Get Order
        StringValue id = StringValue.newBuilder().setValue("101").build();
        OrderManagementOuterClass.Order orderResponse = stub.withCompression("gzip").getOrder(id);
        logger.info("GetOrder Response -> : " + orderResponse.toString());
    }
}
