package ecommerce;

import io.grpc.Metadata;
import io.grpc.ServerCall;
import io.grpc.ServerCallHandler;
import io.grpc.ServerInterceptors;

import java.util.logging.Logger;

import static io.grpc.Metadata.ASCII_STRING_MARSHALLER;

public class MDInterceptor implements io.grpc.ServerInterceptor {

    private static final Logger logger = Logger.getLogger(MDInterceptor.class.getName());


    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata metadata, ServerCallHandler<ReqT, RespT> next) {

        String ret_MD = metadata.get(Metadata.Key.of("MY_MD_1", ASCII_STRING_MARSHALLER));
        logger.info("Metadata Retrived : " + ret_MD);
        return next.startCall(call, metadata);
    }
}
