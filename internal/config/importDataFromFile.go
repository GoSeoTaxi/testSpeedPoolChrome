package config

import (
	"encoding/csv"
	"log"
	"os"
)

func uRLsListImport(fileName string) (arrStr []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		arrStr = append(arrStr, record[0])
	}
	return arrStr
}
