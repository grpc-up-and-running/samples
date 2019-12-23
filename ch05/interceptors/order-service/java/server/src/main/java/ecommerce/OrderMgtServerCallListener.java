package ecommerce;

import io.grpc.ForwardingServerCallListener;
import io.grpc.ServerCall;

import java.util.logging.Logger;

public class OrderMgtServerCallListener<R> extends ForwardingServerCallListener<R> {
    private static final Logger logger = Logger.getLogger(OrderMgtServerCallListener.class.getName());

    private final ServerCall.Listener<R> delegate;

    OrderMgtServerCallListener(ServerCall.Listener<R> delegate) {
        this.delegate = delegate;
    }

    @Override
    protected ServerCall.Listener<R> delegate() {
        return delegate;
    }

    @Override
    public void onMessage(R message) {
        logger.info("Message Received from Client -> Service " + message);
        super.onMessage(message);
    }
}
