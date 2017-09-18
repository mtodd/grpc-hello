build:
	docker build -t grpc-play .

.PHONY: run

run:
	docker run -it -p 8000:8000 -v ${CURDIR}:/go/src/github.com/mtodd/grpc-play grpc-play
