// package main

// import (
// 	"os"
// 	"distributed_calculator/internal/auth"
// 	"distributed_calculator/internal/calculator"
// 	"distributed_calculator/internal/grpcclient"
// 	"distributed_calculator/internal/middleware"
// 	"distributed_calculator/internal/storage"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	grpcClient := grpcclient.NewGRPCClient(os.Getenv("GRPC_ADDRESS"))
// 	db, err := storage.NewStorage("calculator.db")
// 	if err != nil {
// 		log.Fatalf("failed to initialize storage: %v", err)
// 	}

// 	authService := auth.NewAuthService(db)
// 	calculatorService := calculator.NewCalculatorService(db, grpcclient.NewGRPCClient("localhost:50051"))

// 	r := mux.NewRouter()

// 	r.HandleFunc("/api/v1/register", auth.RegisterHandler(authService)).Methods("POST")
// 	r.HandleFunc("/api/v1/login", auth.LoginHandler(authService)).Methods("POST")

// 	api := r.PathPrefix("/api/v1").Subrouter()
// 	api.Use(middleware.JWTMiddleware(authService))
// 	api.HandleFunc("/calculate", calculator.CalculateHandler(calculatorService)).Methods("POST")

// 	log.Println("HTTP server started at :8080")
// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		log.Fatalf("failed to start server: %v", err)
// 	}
// }
package main

import (
 "os"
 "distributed_calculator/internal/auth"
 "distributed_calculator/internal/calculator"
 "distributed_calculator/internal/grpcclient"
 "distributed_calculator/internal/middleware"
 "distributed_calculator/internal/storage"
 "log"
 "net/http"

 "github.com/gorilla/mux"
)

func main() {
 // Инициализация gRPC клиента
//  grpcClient := grpcclient.NewGRPCClient(os.Getenv("GRPC_ADDRESS"))
 grpcAddress := os.Getenv("GRPC_ADDRESS")
 if grpcAddress == "" {
        grpcAddress = "localhost:50051" // дефолт для локальной разработки
    }
 grpcClient := grpcclient.NewGRPCClient(grpcAddress)

 // Инициализация базы данных
 db, err := storage.NewStorage("calculator.db")
 if err != nil {
  log.Fatalf("failed to initialize storage: %v", err)
 }

 // Инициализация сервисов
 authService := auth.NewAuthService(db)
 calculatorService := calculator.NewCalculatorService(db, grpcClient)

 // Инициализация маршрутизатора
 r := mux.NewRouter()

 // Роуты для аутентификации
 r.HandleFunc("/api/v1/register", auth.RegisterHandler(authService)).Methods("POST")
 r.HandleFunc("/api/v1/login", auth.LoginHandler(authService)).Methods("POST")

 // Роуты для вычислений с защитой JWT
 api := r.PathPrefix("/api/v1").Subrouter()
 api.Use(middleware.JWTMiddleware(authService))
 api.HandleFunc("/calculate", calculator.CalculateHandler(calculatorService, db)).Methods("POST")

 // Запуск HTTP-сервера
 log.Println("HTTP server started at :8080")
 if err := http.ListenAndServe(":8080", r); err != nil {
  log.Fatalf("failed to start server: %v", err)
 }
}