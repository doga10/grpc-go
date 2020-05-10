## gRPC - GoLang 

#### Author: Douglas Dennys `douglasdennys@yahoo.com`

## DEV Commands:
- `protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:framework/pb`

### install client gRPC evans

### UP MongoDB: `docker-compose up -d`
### UP Server gRPC: `go run ./framework/cmd/server.go`
### UP Client gRPC `evans -r -p 8080` 