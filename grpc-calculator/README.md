refer: https://medium.com/@DreamsOfImran/learn-grpc-and-protobuf-with-golang-8456a2e64977


generate go from protobuf

    $ protoc -I proto/ --go_out=. --go-grpc_out=. proto/*.proto
