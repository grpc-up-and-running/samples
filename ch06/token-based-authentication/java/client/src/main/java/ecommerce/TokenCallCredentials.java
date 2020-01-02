package ecommerce;

import io.grpc.CallCredentials;
import io.grpc.Metadata;
import io.grpc.Status;

import java.util.Base64;
import java.util.concurrent.Executor;

public class TokenCallCredentials extends CallCredentials {

    private final String credentials;

    TokenCallCredentials(String token) {
        this.credentials = token;
    }

    @Override
    public void applyRequestMetadata(RequestInfo requestInfo, Executor executor, MetadataApplier applier) {
        executor.execute(() -> {
            try {
                Metadata headers = new Metadata();
                Metadata.Key<String> authKey = Metadata.Key.of("Authorization", Metadata.ASCII_STRING_MARSHALLER);
                headers.put(authKey, "Bearer " + credentials);
                applier.apply(headers);
            } catch (Throwable e) {
                applier.fail(Status.UNAUTHENTICATED.withCause(e));
            }
        });
    }

    @Override
    public void thisUsesUnstableApi() {

    }
}
