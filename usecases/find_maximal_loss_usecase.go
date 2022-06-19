package usecases

import (
	"stockexchangelosses/domain"
)

func FindMaximalLossUsecase(rawStockValues []int) int {
	maximalLoss := domain.NewStockValues(rawStockValues).FindMaximalLoss()

	return maximalLoss
}
