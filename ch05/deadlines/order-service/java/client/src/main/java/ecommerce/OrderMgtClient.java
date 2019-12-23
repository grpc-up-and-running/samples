package ecommerce;

import com.google.protobuf.StringValue;
import io.grpc.*;
import io.grpc.stub.StreamObserver;

import java.util.Iterator;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;

public class OrderMgtClient {

    private static final Logger logger = Logger.getLogger(OrderMgtClient.class.getName());

    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress(
                "localhost", 50051).usePlaintext().build();

        OrderManagementGrpc.OrderManagementBlockingStub stub = OrderManagementGrpc.newBlockingStub(channel).withDeadlineAfter(1000, TimeUnit.MILLISECONDS);
        OrderManagementGrpc.OrderManagementStub asyncStub = OrderManagementGrpc.newStub(channel);

        OrderManagementOuterClass.Order order = OrderManagementOuterClass.Order
                .newBuilder()
                .setId("101")
                .addItems("iPhone XS").addItems("Mac Book Pro")
                .setDestination("San Jose, CA")
                .setPrice(2300)
                .build();

        // Add Order

        try {
            StringValue result = stub.addOrder(order);
            logger.info("AddOrder Response -> : " + result.getValue());
        } catch (StatusRuntimeException e) {
            logger.info("Deadline Exceeded. : " + e.getMessage());
        }

    }

}
