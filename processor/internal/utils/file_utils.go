package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var Output_path string

func ReadCsvFile(filePath, delimiter string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %w", filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = []rune(delimiter)[0]

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to parse file %s as CSV: %w", filePath, err)
	}

	return records, nil
}

func ValidateCsvFile(records [][]string) error {
	if len(records) < 2 {
		return fmt.Errorf("file must contain at least a header and one data row")
	}

	for i, row := range records {
		if i == 0 {
			continue
		}
		if len(row) < 3 {
			return fmt.Errorf("row %d has invalid column count: expected 3, got %d", i+1, len(row))
		}

		amount := row[2]
		amount = strings.ReplaceAll(amount, "+", "") // remove '+' for parsing
		_, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			return fmt.Errorf("row %d has invalid amount value '%s': %v", i+1, row[2], err)
		}

	}

	return nil
}

func LogError(fileName string, err error) {
	if _, statErr := os.Stat(Output_path); os.IsNotExist(statErr) {
		os.MkdirAll(Output_path, 0755)
	}

	logFilePath := filepath.Join(Output_path, fileName+".error.log")

	f, openErr := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if openErr != nil {
		log.Printf("Error opening log file: %v", openErr)
		return
	}
	defer f.Close()

	_, writeErr := f.WriteString(fmt.Sprintf("Error processing %s: %v\n", fileName, err))
	if writeErr != nil {
		log.Printf("Error writing to log file: %v", writeErr)
	}
}
