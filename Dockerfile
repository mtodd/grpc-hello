FROM golang:1.9.0

WORKDIR /go/src/github.com/mtodd/grpc-play
COPY . .

RUN /bin/bash script/build

EXPOSE 8000

CMD ["go", "run", "server/main.go"]
