# kanbango-db
A Postgres database docker compose for the kanbango backend

Examples credentials are specified in the `.example` files

# Requirements
- protoc
- export GOPATH=$HOME/go
- export PATH=$PATH:$GOPATH/bin
- go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
- go go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

To run:

`docker compose up -d`

# TODO
- Tests
