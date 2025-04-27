package calculator

import (
 "encoding/json"
 "net/http"
 "distributed_calculator/internal/storage"
)

// Структура для запроса
type CalculateRequest struct {
 UserID    int    `json:"user_id"`
 Expression string `json:"expression"`
}

// Обработчик для вычислений
func CalculateHandler(service *CalculatorService) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {
  var req CalculateRequest

  // Декодирование тела запроса в структуру
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
   http.Error(w, "Invalid request body", http.StatusBadRequest)
   return
  }

  // Вызов метода для расчёта
  result, err := service.Calculate(req.UserID, req.Expression)
  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }

  // Ответ с результатом
  resp := struct {
   Result float64 `json:"result"`
  }{
   Result: result,
  }

  // Установка типа контента и отправка ответа
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(resp)
 }
}