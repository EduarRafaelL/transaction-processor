package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CheckTransacctionType(number float64, creditTransactions, debitTransactions *[]float64) {
	if number > 0 {
		*creditTransactions = append(*creditTransactions, number)
	}
	if number < 0 {
		*debitTransactions = append(*debitTransactions, number)
	}

}

func GetMonthByNumber(number int) string {
	month := time.Month(number)
	return month.String()
}

func GateTransactionMonth(date string) int {
	dateParts := strings.Split(date, "/")
	month := dateParts[0]
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return monthInt
}
