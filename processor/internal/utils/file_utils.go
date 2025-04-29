package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var Output_path string

func ReadCsvFile(filePath, delimiter string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = []rune(delimiter)[0]
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
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
