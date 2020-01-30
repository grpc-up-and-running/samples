package ecommerce;

import io.grpc.CallCredentials;
import io.grpc.Metadata;
import io.grpc.Status;

import java.util.Base64;
import java.util.concurrent.Executor;

public class BasicCallCredentials extends CallCredentials {

    private final String credentials;

    public BasicCallCredentials(String username, String password) {
        this.credentials = username + ":" + password;
    }

    @Override
    public void applyRequestMetadata(RequestInfo requestInfo, Executor executor, MetadataApplier applier) {
        executor.execute(() -> {
            try {
                Metadata headers = new Metadata();
                Metadata.Key<String> authKey = Metadata.Key.of("authorization", Metadata.ASCII_STRING_MARSHALLER);
                headers.put(authKey, "Basic " + Base64.getEncoder().encodeToString(credentials.getBytes()));
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
