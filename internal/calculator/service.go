package calculator

import (
 "distributed_calculator/internal/grpcclient"
 "distributed_calculator/internal/storage"
)

type CalculatorService struct {
 storage *storage.Storage
 client  *grpcclient.GRPCClient
}

// Создание нового CalculatorService
func NewCalculatorService(storage *storage.Storage, client *grpcclient.GRPCClient) *CalculatorService {
 return &CalculatorService{storage: storage, client: client}
}

// Метод для вычислений
func (s *CalculatorService) Calculate(userID int, expression string) (float64, error) {
 // Вызов метода gRPC клиента для вычисления
 result, err := s.client.Calculate(expression)
 if err != nil {
  return 0, err
 }
 
 // Сохранение результата в базу данных
 if err := s.storage.SaveCalculation(userID, expression, result); err != nil {
  return 0, err
 }
 
 return result, nil
}