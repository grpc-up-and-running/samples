//
// Created by lldx on 1/19/23.
//

#include "ProductInfoImpl.h"
#include <sstream>
grpc::Status ProductInfoImpl::addProduct(::grpc::ServerContext *context, const ::ecommerce::Product *request,
                                         ::ecommerce::ProductID *response) {
    std::cout << request->DebugString();
    std::string id = uuid_gen_test();
    (*server_info)[id]=*request;
    (*server_info)[id].set_id(id);
    response->set_value(id);
    return grpc::Status::OK;
}

grpc::Status ProductInfoImpl::getProduct(::grpc::ServerContext *context, const ::ecommerce::ProductID *request,
                                         ::ecommerce::Product *response) {
    std::cout << request->DebugString();
    if ( server_info->find(request->value()) != server_info->end() ){
        response->CopyFrom((*server_info)[request->value()]);
        return grpc::Status::OK;
    }
    return grpc::Status::CANCELLED;
}

ProductInfoImpl::~ProductInfoImpl() {

}

std::string ProductInfoImpl::uuid_gen_test() {
    static int i = 1;
    std::stringstream ss;
    ss << i ;
    i++;
    return ss.str();
}
