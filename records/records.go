package records

import (
	"encoding/csv"
	"os"
	"strings"
)

func getCSVHeaders() ([]string, error) {
	f, err := os.Open("binlist-data.csv")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	record, err := csvReader.Read()
	return record, err
}

const (
	Bin = iota
	Brand
	Type
	Category
	Issuer
	Alpha_2
	Alpha_3
	Country
	Latitude
	Longitude
	Bank_phone
	Bank_url
)

func GetRecords() ([][]string, error) {
	f, err := os.Open("binlist-data.csv")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	if _, err := csvReader.Read(); err != nil {
		return nil, err
	}

	record, err := csvReader.ReadAll()
	return record, err
}

func GetBinDataFromRecord(cardNumber string, record [][]string) []string {
	for _, r := range record {
		if strings.Contains(cardNumber[:8], r[Bin]) {
			return r
		}
	}

	return nil
}
