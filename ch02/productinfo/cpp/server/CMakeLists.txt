# Copyright 2018 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# cmake build file for C++ product_info example.
# Assumes protobuf and gRPC have been installed using cmake.
# See cmake_externalproject/CMakeLists.txt for all-in-one cmake build
# that automatically builds all the dependencies before building product_info.

cmake_minimum_required(VERSION 3.5.1)

project(product_info_server C CXX)

include(../cmake/common.cmake)

# Proto file
get_filename_component(pd_proto "../../proto/product_info.proto" ABSOLUTE)
get_filename_component(pd_proto_path "${pd_proto}" PATH)
# Generated sources
set(pd_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/product_info.pb.cc")
set(pd_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/product_info.pb.h")
set(pd_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/product_info.grpc.pb.cc")
set(pd_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/product_info.grpc.pb.h")
add_custom_command(
      OUTPUT "${pd_proto_srcs}" "${pd_proto_hdrs}" "${pd_grpc_srcs}" "${pd_grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
        --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
        -I "${pd_proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${pd_proto}"
      DEPENDS "${pd_proto}")

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# pd_grpc_proto
add_library(pd_grpc_proto
  ${pd_grpc_srcs}
  ${pd_grpc_hdrs}
  ${pd_proto_srcs}
  ${pd_proto_hdrs})
target_link_libraries(pd_grpc_proto
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF})

add_executable(product_info_server ProductInfoServer.cc ProductInfoImpl.cpp ProductInfoImpl.h)
target_link_libraries(product_info_server pd_grpc_proto ${_REFLECTION} ${_GRPC_GRPCPP} ${_PROTOBUF_LIBPROTOBUF})