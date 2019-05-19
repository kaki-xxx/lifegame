package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// FromFile creates table from reading file
func FromFile(filePath string) [][]bool {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Can't open file")
	}
	defer file.Close()

	const bufsize = 100
	newtable := make([][]bool, 0, bufsize)
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error:", err)
		}
		cols := make([]bool, len(record))
		for x, v := range record {
			if v == "1" {
				cols[x] = true
			}
		}
		newtable = append(newtable, cols)
	}
	return newtable
}
