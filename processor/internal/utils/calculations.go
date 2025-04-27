package utils

func GetAverageTransaction(transactions []float64) float64 {
	if len(transactions) == 0 {
		return 0
	}
	var total float64
	for _, transaction := range transactions {
		total += transaction
	}
	return total / float64(len(transactions))
}

func GetTotalTransactions(transactions []float64) int {
	return len(transactions)
}
