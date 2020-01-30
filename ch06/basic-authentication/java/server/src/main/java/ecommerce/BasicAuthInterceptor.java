package ecommerce;

import io.grpc.Context;
import io.grpc.Contexts;
import io.grpc.Metadata;
import io.grpc.ServerCall;
import io.grpc.ServerCallHandler;
import io.grpc.ServerInterceptor;
import io.grpc.Status;

import java.util.Base64;
import java.util.logging.Logger;

import static io.grpc.Metadata.ASCII_STRING_MARSHALLER;

public class BasicAuthInterceptor implements ServerInterceptor {

    private static final ServerCall.Listener NOOP_LISTENER = new ServerCall.Listener() {
    };
    private static final String ADMIN_USER_CREDENTIALS = "admin:admin";
    private static final Context.Key<String> USER_ID_CTX_KEY = Context.key("userId");
    private static final String ADMIN_USER_ID = "admin";
    private static final Logger logger = Logger.getLogger(BasicAuthInterceptor.class.getName());

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata headers, ServerCallHandler<ReqT, RespT> next) {
        String basicAuthString = headers.get(Metadata.Key.of("authorization", ASCII_STRING_MARSHALLER));
        if (basicAuthString == null) {
            call.close(Status.UNAUTHENTICATED.withDescription("Basic authentication value is missing in Metadata"),
                    headers);
            return NOOP_LISTENER;
        }
        if (validUser(basicAuthString)) {
            Context ctx = Context.current().withValue(USER_ID_CTX_KEY, ADMIN_USER_ID);
            return Contexts.interceptCall(ctx, call, headers, next);
        } else {
            logger.info("Verification failed - Unauthenticated!");
            call.close(Status.UNAUTHENTICATED.withDescription("Invalid basic credentials"), headers);
            return NOOP_LISTENER;
        }
    }

    private boolean validUser(String basicAuthString) {
        if (basicAuthString == null) {
            return false;
        }
        String token = basicAuthString.substring("Basic ".length()).trim();
        byte[] byteArray = Base64.getDecoder().decode(token.getBytes());
        return ADMIN_USER_CREDENTIALS.equals(new String(byteArray));
    }
}
