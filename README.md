# libra_example

```
# install dependencies
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go

# extract the .proto's from libra and compile to go
# set LIBRA_SRC to wherever you cloned libra/libra
$ LIBRA_SRC=$HOME/libra ./gen_protos.sh

# make a basic request to a testnet node
$ go run libra_client/main.go
```
