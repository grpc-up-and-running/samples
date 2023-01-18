//
// Created by lldx on 1/19/23.
//

#ifndef PRODUCT_INFO_SERVER_PRODUCTINFOIMPL_H
#define PRODUCT_INFO_SERVER_PRODUCTINFOIMPL_H

#include "product_info.grpc.pb.h"

class ProductInfoImpl final : public ecommerce::ProductInfo::Service{
private:
    std::map<std::string,ecommerce::Product> * server_info = new std::map<std::string,ecommerce::Product>();
public:
    ~ProductInfoImpl() override;
    std::string uuid_gen_test();
public:
    grpc::Status addProduct(::grpc::ServerContext *context, const ::ecommerce::Product *request,
                            ::ecommerce::ProductID *response) override;

    grpc::Status getProduct(::grpc::ServerContext *context, const ::ecommerce::ProductID *request,
                            ::ecommerce::Product *response) override;
};


#endif //PRODUCT_INFO_SERVER_PRODUCTINFOIMPL_H
