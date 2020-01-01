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

public class TokenAuthInterceptor implements ServerInterceptor {

    private static final ServerCall.Listener NOOP_LISTENER = new ServerCall.Listener() {
    };
    private static final String ADMIN_USER_TOKEN = "some-secret-token";
    private static final Context.Key<String> USER_ID_CTX_KEY = Context.key("userId");
    private static final String ADMIN_USER_ID = "admin";
    private static final Logger logger = Logger.getLogger(TokenAuthInterceptor.class.getName());

    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata headers, ServerCallHandler<ReqT, RespT> next) {
        String tokenString = headers.get(Metadata.Key.of("Authorization", ASCII_STRING_MARSHALLER));
        if (tokenString == null) {
            call.close(Status.UNAUTHENTICATED.withDescription("Token value is missing in Metadata"),
                    headers);
            return NOOP_LISTENER;
        }
        if (validUser(tokenString)) {
            Context ctx = Context.current().withValue(USER_ID_CTX_KEY, ADMIN_USER_ID);
            return Contexts.interceptCall(ctx, call, headers, next);
        } else {
            logger.info("Verification failed - Unauthenticated!");
            call.close(Status.UNAUTHENTICATED.withDescription("Invalid user token"), headers);
            return NOOP_LISTENER;
        }
    }

    private boolean validUser(String basicAuthString) {
        if (basicAuthString == null) {
            return false;
        }
        // Perform the token validation here. For the sake of this example, the code
        // here forgoes any of the usual OAuth2 token validation and instead checks
        // for a token matching an arbitrary string.
        String token = basicAuthString.substring("Bearer ".length()).trim();
        return ADMIN_USER_TOKEN.equals(token);
    }
}
