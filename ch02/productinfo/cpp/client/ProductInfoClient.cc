/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>
#include "product_info.grpc.pb.h"


using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
//using ecommerce::Product;
//using ecommerce::ProductInfo;
//using ecommerce::HelloRequest;

class ProductInfoClient {
public:
    ProductInfoClient(std::shared_ptr<Channel> channel)
            : stub_(ecommerce::ProductInfo::NewStub(channel)) {}

    // Assembles the client's payload, sends it and presents the response back
    // from the server.
    std::string addProduct(const std::string & name ,const std::string & desc , const float price) {
        // Data we are sending to the server.
        ecommerce::Product request;
        request.set_name(name);
        request.set_description(desc);
        request.set_price(price);

        // Container for the data we expect from the server.
        ecommerce::ProductID reply;

        // Context for the client. It could be used to convey extra information to
        // the server and/or tweak certain RPC behaviors.
        ClientContext context;

        // The actual RPC.
        Status status = stub_->addProduct(&context, request, &reply);

        // Act upon its status.
        if (status.ok()) {
            return reply.value();
        } else {
            std::cout << status.error_code() << ": " << status.error_message()
                      << std::endl;
            return "RPC failed";
        }
    }

    std::string getProduct(const std::string & id ){
        ecommerce::ProductID request;
        request.set_value(id);
        ecommerce::Product reply;
        ClientContext context;
        Status status = stub_->getProduct(&context,request,&reply);

        if (status.ok()){
            return reply.DebugString();
        }
        else{
            std::cout << status.error_code() << ": " << status.error_message()
                      << std::endl;
            return "RPC failed";
        }
    }

private:
    std::unique_ptr<ecommerce::ProductInfo::Stub> stub_;
};

int main(int argc, char** argv) {
    // Instantiate the client. It requires a channel, out of which the actual RPCs
    // are created. This channel models a connection to an endpoint specified by
    // the argument "--target=" which is the only expected argument.
    // We indicate that the channel isn't authenticated (use of
    // InsecureChannelCredentials()).
    std::string target_str;
    std::string arg_str("--target");
    if (argc > 1) {
        std::string arg_val = argv[1];
        size_t start_pos = arg_val.find(arg_str);
        if (start_pos != std::string::npos) {
            start_pos += arg_str.size();
            if (arg_val[start_pos] == '=') {
                target_str = arg_val.substr(start_pos + 1);
            } else {
                std::cout << "The only correct argument syntax is --target="
                          << std::endl;
                return 0;
            }
        } else {
            std::cout << "The only acceptable argument is --target=" << std::endl;
            return 0;
        }
    } else {
        target_str = "localhost:50051";
    }
    ProductInfoClient product_info_server(
            grpc::CreateChannel(target_str, grpc::InsecureChannelCredentials()));
    std::string reply = product_info_server.addProduct(
            "Apple iPhone 11",
            "Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Nigh mode.",
            1000.0);
    std::cout << "Greeter received: " << reply << std::endl;

    std::string product_info = product_info_server.getProduct(reply);
    std::cout << "Greeter received: " << product_info << std::endl;

    return 0;
}
