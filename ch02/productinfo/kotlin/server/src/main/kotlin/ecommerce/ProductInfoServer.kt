package ecommerce

import io.grpc.Server
import io.grpc.ServerBuilder
import java.util.UUID
import io.grpc.Status.NOT_FOUND
import io.grpc.StatusException

class ProductInfoServer(
    private val port: Int,
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
    val productMap = HashMap<String,Product>()

    override suspend fun addProduct(request: Product): ProductID {
      val uuid = UUID.randomUUID()
      val product = product {
        id = uuid.toString()
        name = request.name
        description = request.description
        price = request.price
      }
      // Add product to the inmemory map.
      productMap[product.id] = product

      return productID {
        value = product.id
      }
    }

    override suspend fun getProduct(request: ProductID): Product {
      val productID = request.value
      return productMap[productID] ?:
        throw StatusException(NOT_FOUND.withDescription("Requested product with id $productID doesn't exist"))
    }
  }
}

fun main() {
  val port = 50051
  val server = ProductInfoServer(port)
  server.start()
  server.blockUntilShutdown()
}