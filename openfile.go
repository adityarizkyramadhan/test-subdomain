package main

import (
	"encoding/csv"
	"os"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	panicIfError(err)
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	panicIfError(err)
	var subdomain []string
	for _, each := range data {
		subdomain = append(subdomain, each[0])
	}
	return subdomain
}
