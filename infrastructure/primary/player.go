package main

import (
	"bufio"
	"fmt"
	"os"
	"stockexchangelosses/usecases"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")
	rawStockValues := make([]int, n)
	for i := 0; i < n; i++ {
		v, _ := strconv.ParseInt(inputs[i], 10, 32)
		rawStockValues[i] = int(v)
	}

	maximalLoss := usecases.FindMaximalLossUsecase(rawStockValues)

	fmt.Println(maximalLoss)
}
