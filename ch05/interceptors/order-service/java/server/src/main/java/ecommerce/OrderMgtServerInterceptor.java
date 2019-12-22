package ecommerce;

import io.grpc.*;

import java.util.logging.Logger;

public class OrderMgtServerInterceptor implements io.grpc.ServerInterceptor {
    private static final Logger logger = Logger.getLogger(OrderMgtServerInterceptor.class.getName());

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata headers, ServerCallHandler<ReqT, RespT> next) {

        logger.info("======= [Server Interceptor] : Remote Method Invoked - " + call.getMethodDescriptor().getFullMethodName());
        ServerCall<ReqT, RespT> serverCall = new OrderMgtServerCall<>(call);
        return new OrderMgtServerCallListener<>(next.startCall(serverCall, headers));
    }

}

