package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func extractFromCSVFile(filePath string) *Hotels {
	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file(%s): %s", filePath, err)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	//match the rest of the records with the number of records in the first line
	reader.FieldsPerRecord = 0

	//reading first record to ignore the data titles line
	record, err := reader.Read()
	checkReadErr(err)

	fmt.Printf("Record first line is:%s-\n", record)

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading file(%s): %s", filePath, err)
	}

	hotels := parseToStruct(rawCSVData)

	fmt.Printf("\nhotels read are len:%d\n", len(hotels.Hotels))
	return hotels
}

func parseToStruct(rawData [][]string) *Hotels {
	hotels := make([]Hotel, 0)
	for _, each := range rawData {
		name, address, starsStr, contact, phone, uri := each[0], each[1], each[2], each[3], each[4], each[5]
		nameValidator := &NameValidator{name}
		starsValidator := &StarsValidator{starsStr}
		urlValidator := &URLValidator{uri}
		if validateData(nameValidator) && validateData(starsValidator) && validateData(urlValidator) {
			stars, _ := strconv.Atoi(starsStr)
			hotel := Hotel{Name: name, Address: address, Stars: stars, Contact: contact, Phone: phone, Uri: uri}
			hotels = append(hotels, hotel)
		}
	}
	return &Hotels{hotels}
}

func checkReadErr(e error) {
	if e == io.EOF {
		log.Fatal("No data found in file")
	}
	if e != nil {
		log.Fatalf("Error reading file: %s", e)
	}
}
