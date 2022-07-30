package ecommerce

import ecommerce.ProductInfoGrpcKt.ProductInfoCoroutineStub
import io.grpc.ManagedChannel
import io.grpc.ManagedChannelBuilder
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import kotlinx.coroutines.runBlocking
import java.io.Closeable
import java.util.concurrent.TimeUnit

class ProductInfoClient (
    private val channel: ManagedChannel,
) : Closeable {
  private val stub: ProductInfoCoroutineStub = ProductInfoCoroutineStub(channel)

  suspend fun addProduct(name: String, description: String, price: Float): String {
    val product = product {
      this.name = name
      this.description = description
      this.price = price
    }

    val response = stub.addProduct(product)
    val productID = response.value
    println("Product ID:  $productID added successfully.")
    return productID
  }

  suspend fun getProduct(productID: String): Product {
    val request = productID {
      value = productID
    }
    val response = stub.getProduct(request)
    println("Product: $response")
    return response
  }

  override fun close() {
    channel.shutdown().awaitTermination(5, TimeUnit.SECONDS)
  }
}


fun main() = runBlocking {
  val port = 50051

  val client = ProductInfoClient(
        ManagedChannelBuilder.forAddress("localhost", port)
            .usePlaintext()
            .executor(Dispatchers.Default.asExecutor())
            .build())


  val productID = client.addProduct("Samsung S10", "Samsung Galaxy S10 is the latest smart phone, " +
          "launched in February 2019", 700.0f)
  val product = client.getProduct(productID)
  client.close()
}