## Generate the protobuf

```
brew install protobuf-c

go get -u github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$GOPATH/bin

./script/generate-protos
```

## Build and run the container (server)

```
docker build -t grpc-play .

docker run -it -p 8000:8000 -v $(pwd):/go/src/github.com/mtodd/grpc-play grpc-play
```
