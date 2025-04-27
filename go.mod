module distributed_calculator

go 1.22

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/golang-jwt/jwt/v5 v5.1.0
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.20
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.63.2
