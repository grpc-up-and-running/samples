package ecommerce

import io.grpc.ManagedChannel
import io.grpc.ManagedChannelBuilder
import ecommerce.ProductInfoGrpcKt.ProductInfoCoroutineStub
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import kotlinx.coroutines.runBlocking
import java.io.Closeable
import java.util.concurrent.TimeUnit
import ecommerce.Product
import ecommerce.ProductID

class ProductInfoClient constructor(
    private val channel: ManagedChannel
) : Closeable {
  private val stub: ProductInfoCoroutineStub = ProductInfoCoroutineStub(channel)

  suspend fun addProduct(name: String, description: String, price: Float): String = runBlocking {
    var product = Product.newBuilder().apply {
      this.name = name
      this.description = description
      this.price = price
    }.build()

    var response = stub.addProduct(product)
    var productID = response.value
    println("Product ID:  ${productID} added successfully.")
    return@runBlocking productID
  }

  suspend fun getProduct(productID: String): Product = runBlocking {
    var request = ProductID.newBuilder().apply {
      this.value = productID;
    }.build()
    var response = stub.getProduct(request)
    println("Product: ${response}")
    return@runBlocking response
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


  var productID = client.addProduct("Samsung S10", "Samsung Galaxy S10 is the latest smart phone, " +
          "launched in February 2019", 700.0f)
  var product = client.getProduct(productID)
  client.close()
}