## Generate the protobuf

```
brew install protobuf-c

go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin

./script/generate-protos
```

## Build and run the container (server)

```
make
make run
```

## Build and run the client

```
go run client/main.go "Your Name"
```
