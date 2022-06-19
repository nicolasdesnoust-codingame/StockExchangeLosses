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
	log.Println(rawStockValues)

	for _, rawStockValue := range rawStockValues {
		stockValues.Append(StockValue(rawStockValue))
	}
	log.Println(stockValues.Content)

	return &stockValues
}

func (stockValues *StockValues) FindMaximalLoss() int {
	losses := make([]int, 0)
	log.Println(stockValues.Content)

	// Parcourir le tableau dans le sens inverse et en remplir un nouveau avec à chaque case le minimum observé dans la séquence
	minimums := make([]int, stockValues.Size()-1)
	currentMinimum := math.MaxInt32
	for stockValueIndex := stockValues.Size() - 1; stockValueIndex > 0; stockValueIndex-- {
		if int(stockValues.Content[stockValueIndex]) < currentMinimum {
			currentMinimum = int(stockValues.Content[stockValueIndex])
		}

		minimums[stockValueIndex-1] = currentMinimum
	}
	log.Println(minimums)

	for stockValueIndex := 0; stockValueIndex < stockValues.Size()-1; stockValueIndex++ {
		minimumAmongFollowingValues := minimums[stockValueIndex]

		losses = append(losses, int(stockValues.Content[stockValueIndex])-int(minimumAmongFollowingValues))
	}

	maximalLoss := 0
	for _, loss := range losses {
		if loss > maximalLoss {
			maximalLoss = loss
		}
	}

	return -maximalLoss
}
