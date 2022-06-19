package usecases

import (
	"testing"
)

func TestFindMaximalLoss_ShouldReturnNoMaximalLoss(t *testing.T) {
	rawStockValues := []int{1, 2, 3, 4, 5, 6}

	actualMaximalLoss := FindMaximalLossUsecase(rawStockValues)

	expectedMaximalLoss := 0
	if actualMaximalLoss != expectedMaximalLoss {
		t.Fatalf(`Actual maximal loss "%d" does not match expected maximal loss "%d"`, actualMaximalLoss, expectedMaximalLoss)
	}
}

func TestFindMaximalLoss_ShouldReturnExpectedMaximalLoss(t *testing.T) {
	rawStockValues := []int{3, 2, 4, 2, 1, 5}

	actualMaximalLoss := FindMaximalLossUsecase(rawStockValues)

	expectedMaximalLoss := -3
	if actualMaximalLoss != expectedMaximalLoss {
		t.Fatalf(`Actual maximal loss "%d" does not match expected maximal loss "%d"`, actualMaximalLoss, expectedMaximalLoss)
	}
}
