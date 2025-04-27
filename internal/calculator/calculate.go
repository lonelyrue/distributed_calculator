// package calculator

// import (
//     "encoding/json"
//     "net/http"
//     "distributed_calculator/internal/auth"
// )

// // Структура для запроса
// type CalculationRequest struct {
//     Expression string  `json:"expression"`
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

//         // Расчёт результата
//         result, err := service.Calculate(userID, req.Expression)
//         if err != nil {
//             http.Error(w, "calculation error", http.StatusInternalServerError)
//             return
//         }

//         // Ответ с результатом
//         json.NewEncoder(w).Encode(map[string]interface{}{
//             "result": result,
//         })
//     }
// }
package calculator

import (
 "distributed_calculator/internal/auth"
 "distributed_calculator/internal/storage"
 "encoding/json"
 "net/http"
)

// Структура для запроса вычислений
type CalculationRequest struct {
 Expression string `json:"expression"`
}

// Обработчик для вычислений
func CalculateHandler(service *CalculatorService, storage *storage.Storage) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {
  // Получаем ID пользователя из контекста
  userID := auth.GetUserID(r.Context())

  var req CalculationRequest
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
   http.Error(w, "invalid request", http.StatusBadRequest)
   return
  }

  // Вызов метода для расчёта
  result, err := service.Calculate(userID, req.Expression)
  if err != nil {
   http.Error(w, "calculation error", http.StatusInternalServerError)
   return
  }

  // Сохраняем результат в базе данных
  err = storage.SaveCalculation(userID, req.Expression, result)
  if err != nil {
   http.Error(w, "failed to save calculation", http.StatusInternalServerError)
   return
  }

  // Ответ с результатом
  json.NewEncoder(w).Encode(map[string]interface{}{
   "result": result,
  })
 }
}