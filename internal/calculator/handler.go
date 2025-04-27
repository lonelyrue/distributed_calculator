// package calculator

// import (
//     "encoding/json"
//     "net/http"
//     "distributed_calculator/internal/auth"
// )

// // Структура для запроса
// type CalculationRequest struct {
//     Expression string `json:"expression"`
// }

// // Обработчик для вычислений
// func CalculateHandler(service *CalculatorService) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         userID := auth.GetUserID(r.Context())

//         var req CalculationRequest
//         if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//             http.Error(w, "invalid request", http.StatusBadRequest)
//             return
//         }

//         result, err := service.Calculate(userID, req.Expression)
//         if err != nil {
//             http.Error(w, "calculation error", http.StatusInternalServerError)
//             return
//         }

//         json.NewEncoder(w).Encode(map[string]interface{}{
//             "result": result,
//         })
//     }
// }

// package calculator

// import (
//  "encoding/json"
//  "net/http"
//  "distributed_calculator/internal/auth"
// )

// // Структура для запроса
// type CalculateRequest struct {
//  Expression string `json:"expression"`
// }
package calculator

// import (
//  "encoding/json"
//  "net/http"
// )

// Структура для запроса
type CalculateRequest struct {
 Expression string `json:"expression"`
}

// Этот файл теперь не содержит функции CalculateHandler, так как она уже присутствует в calculate.go

// Этот файл теперь не содержит функцию CalculateHandler, так как она уже присутствует в calculate.go