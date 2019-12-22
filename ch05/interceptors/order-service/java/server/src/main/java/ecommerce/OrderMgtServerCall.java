package ecommerce;

import io.grpc.ForwardingServerCall;
import io.grpc.MethodDescriptor;
import io.grpc.ServerCall;

import java.util.logging.Logger;

public class OrderMgtServerCall<ReqT, RespT> extends ForwardingServerCall.SimpleForwardingServerCall<ReqT, RespT>  {

    private static final Logger logger = Logger.getLogger(OrderMgtServerCall.class.getName());

    OrderMgtServerCall(ServerCall<ReqT, RespT> delegate) {
        super(delegate);
    }

    @Override
    protected ServerCall<ReqT, RespT> delegate() {
        return super.delegate();
    }

    @Override
    public MethodDescriptor<ReqT, RespT> getMethodDescriptor() {
        return super.getMethodDescriptor();
    }

    @Override
    public void sendMessage(RespT message) {
        logger.info("Message from Service -> Client : " + message);
        super.sendMessage(message);
    }
}
