package ecommerce;

import io.grpc.*;
import static io.grpc.Metadata.ASCII_STRING_MARSHALLER;


public class OrderClientInterceptor implements ClientInterceptor {
    @Override
    public <ReqT, RespT> ClientCall<ReqT, RespT> interceptCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions, Channel channel) {

        return new ForwardingClientCall.SimpleForwardingClientCall<ReqT, RespT>(channel.newCall(method, callOptions)) {
            @Override
            public void start(Listener<RespT> responseListener, Metadata headers) {
                headers.put(Metadata.Key.of("MY_MD_1", ASCII_STRING_MARSHALLER), "This is metadata of MY_MD_1");
                super.start(responseListener, headers);
            }
        };
    }
}
