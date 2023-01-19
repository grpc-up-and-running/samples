
# Install Environment
```shell
export MY_INSTALL_DIR=$HOME/.local
mkdir -p $MY_INSTALL_DIR
export PATH="$MY_INSTALL_DIR/bin:$PATH"
sudo apt install -y cmake
sudo apt install -y build-essential autoconf libtool pkg-config
git clone --recurse-submodules -b v1.50.0 --depth 1 --shallow-submodules https://github.com/grpc/grpc
cd grpc
mkdir -p cmake/build
pushd cmake/build
cmake -DgRPC_INSTALL=ON \
      -DgRPC_BUILD_TESTS=OFF \
      -DCMAKE_INSTALL_PREFIX=$MY_INSTALL_DIR \
      ../..
make -j 8
sudo make install
popd
```

如果你github的下载速度不够，可以参考一下方法给命令行设置代理
```shell
export http_proxy=http://ip:port
export https_proxy=http://ip:port
```

# Build And Run Server
```shell
cd server
mkdir build
cd build 
cmake ..
make -j4
./product_info_server
```

# Build And Run Client
```shell
cd client
mkdir build
cd build 
cmake ..
make -j4
./product_info_client
```

# proto
```shell
cd server
protoc -I ../../proto --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` ../../proto/product_info.proto
protoc -I ../../proto --cpp_out=. ../../proto/product_info.proto
```