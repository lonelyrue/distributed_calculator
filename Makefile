run:
 go run ./cmd/calc_service/...

run-compute:
 go run ./cmd/compute_service/...

protoc:
 protoc --go_out=. --go-grpc_out=. proto/calculator.proto

test:
 go test ./...