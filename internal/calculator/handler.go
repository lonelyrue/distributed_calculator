package calculator

import (
	"distributed_calculator/internal/auth"
	"encoding/json"
	"net/http"
)

type CalculationRequest struct {
	Expression string `json:"expression"`
}

func CalculateHandler(service *CalculatorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := auth.GetUserID(r.Context())

		var req CalculationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		result, err := service.Calculate(userID, req.Expression)
		if err != nil {
			http.Error(w, "calculation error", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"result": result,
		})
	}
}
