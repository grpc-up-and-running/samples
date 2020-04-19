package ecommerce

import io.grpc.Server
import io.grpc.ServerBuilder
import java.util.UUID
import io.grpc.Status.NOT_FOUND
import io.grpc.StatusException

class ProductInfoServer constructor(
    private val port: Int
) {
  val server: Server = ServerBuilder
      .forPort(port)
      .addService(ProductInfoService())
      .build()

  fun start() {
    server.start()
    println("Server started, listening on $port")
    Runtime.getRuntime().addShutdownHook(
        Thread {
          println("*** shutting down gRPC server since JVM is shutting down")
          this@ProductInfoServer.stop()
          println("*** server shut down")
        }
    )
  }

  private fun stop() {
    server.shutdown()
  }

  fun blockUntilShutdown() {
    server.awaitTermination()
  }

  private class ProductInfoService : ProductInfoGrpcKt.ProductInfoCoroutineImplBase() {
    val productMap:HashMap<String,Product> = HashMap<String,Product>()

    override suspend fun addProduct(request: Product): ProductID {
      var uuid = UUID.randomUUID()
      var product = Product.newBuilder().apply {
        this.id = uuid.toString()
        this.name = request.name
        this.description = request.description
        this.price = request.price
      }.build()
      // Add product to the inmemory map.
      productMap.put(product.id, product)

      return ProductID.newBuilder().apply {
        this.value = product.id
      }.build()
    }

    override suspend fun getProduct(request: ProductID): Product {
      val productID = request.value
      return productMap.get(productID) ?:
      throw StatusException(NOT_FOUND.withDescription("Requested product with id ${productID} doesn't exist"))
    }
  }
}

fun main() {
  val port = 50051
  val server = ProductInfoServer(port)
  server.start()
  server.blockUntilShutdown()
}