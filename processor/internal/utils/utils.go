package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetTransactionType(number float64) int {
	if number > 0 {
		return 1
	}
	if number < 0 {
		return 2
	}
	return 0
}

func GetMonthByNumber(number int) string {
	month := time.Month(number)
	return month.String()
}

func GetTransactionMonth(date string) int {
	dateParts := strings.Split(date, "/")
	month := dateParts[0]
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return monthInt
}

func ParseDate(date string) time.Time {
	yearString := strconv.Itoa(time.Now().Year())
	date = date + "/" + yearString
	parsedDate, err := time.Parse("1/2/2006", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return parsedDate
}
