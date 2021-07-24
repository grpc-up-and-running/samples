import ballerina/log;

public function main() returns error? {
    ProductInfoClient productClient = check new ("http://localhost:9090");

    ProductID productId = check productClient->addProduct({name: "Samsung S10", 
    description: "Samsung Galaxy S10 is the latest smart phone, launched in February 2019", 
    price: 700.0f});
    log:printInfo("Product added successfully", productID = productId);

    Product product = check productClient->getProduct(productId);
    log:printInfo("Product retrieved successfully", product = product);
}
