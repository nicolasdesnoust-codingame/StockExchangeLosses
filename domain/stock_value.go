package domain

import (
	"log"
	"math"
)

type StockValue int

type StockValues struct {
	Content []StockValue
}

func (stockValues *StockValues) Append(stockValue StockValue) {
	stockValues.Content = append(stockValues.Content, stockValue)
}

func (stockValues *StockValues) Size() int {
	return len(stockValues.Content)
}

func NewStockValues(rawStockValues []int) *StockValues {
	var stockValues StockValues

	for _, rawStockValue := range rawStockValues {
		stockValues.Append(StockValue(rawStockValue))
	}

	return &stockValues
}

func (stockValues *StockValues) FindMaximalLoss() int {
	minimums := stockValues.findMinimumAmongFollowersForEachStockValue()

	losses := stockValues.findMaximalLossForEachStockValue(minimums)

	maximalLoss := findMaximumPositiveValueIn(losses)
	return -maximalLoss
}

func (stockValues *StockValues) findMinimumAmongFollowersForEachStockValue() []int {
	minimums := make([]int, stockValues.Size()-1)
	currentMinimum := math.MaxInt32

	for stockValueIndex := stockValues.Size() - 1; stockValueIndex > 0; stockValueIndex-- {
		if int(stockValues.Content[stockValueIndex]) < currentMinimum {
			currentMinimum = int(stockValues.Content[stockValueIndex])
		}

		minimums[stockValueIndex-1] = currentMinimum
	}

	log.Println(minimums)
	return minimums
}

func (stockValues *StockValues) findMaximalLossForEachStockValue(minimums []int) []int {
	losses := make([]int, 0)

	for stockValueIndex := 0; stockValueIndex < stockValues.Size()-1; stockValueIndex++ {
		minimumAmongFollowers := minimums[stockValueIndex]
		loss := int(stockValues.Content[stockValueIndex]) - int(minimumAmongFollowers)
		losses = append(losses, loss)
	}

	return losses
}

func findMaximumPositiveValueIn(array []int) int {
	maximum := 0

	for _, item := range array {
		if item > maximum {
			maximum = item
		}
	}

	return maximum
}
