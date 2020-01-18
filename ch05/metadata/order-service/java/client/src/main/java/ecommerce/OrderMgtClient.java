package ecommerce;

import com.google.protobuf.StringValue;
import io.grpc.*;
import io.grpc.stub.MetadataUtils;
import io.grpc.stub.StreamObserver;

import java.util.Iterator;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.logging.Logger;
import static io.grpc.Metadata.ASCII_STRING_MARSHALLER;


public class OrderMgtClient {

    private static final Logger logger = Logger.getLogger(OrderMgtClient.class.getName());

    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress(
                "localhost", 50051)
                .usePlaintext()
                .intercept(new OrderClientInterceptor())
                .build();




//        Metadata metadata = new Metadata();
//        metadata.put(Metadata.Key.of("MY_MD_1", ASCII_STRING_MARSHALLER), "This is metadata of MY_MD_1");
        OrderManagementGrpc.OrderManagementBlockingStub stub = OrderManagementGrpc.newBlockingStub(channel).withDeadlineAfter(1000, TimeUnit.MILLISECONDS);



        // Creating an order with invalid Order ID.
        OrderManagementOuterClass.Order order = OrderManagementOuterClass.Order
                .newBuilder()
                .setId("101")
                .addItems("iPhone XS").addItems("Mac Book Pro")
                .setDestination("San Jose, CA")
                .setPrice(2300)
                .build();



        StringValue result = stub.addOrder(order);
        logger.info("AddOrder Response -> : " + result.getValue());


    }

}
