package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const (
	dataInFile      = "hotels.csv"
	dataOutFileJSON = "hotels.json"
	dataOutFileTOML = "hotels.toml"
)

func main() {
	filePath, isSort, err := checkArgs()
	if err != nil {
		log.Printf("Error: %s\n", err)
		filePath = dataInFile
		log.Printf("No input file found...going to read: %s\n\n", filePath)
	}
	dir := filepath.Dir(filePath)
	fmt.Printf("dir is:%s\n", dir)
	hotels := extractFromCSVFile(filePath)

	if isSort {
		sort.Sort(NameSorter(*hotels))
	}

	jsonWriter := &JSONWriter{filepath.Join(dir, dataOutFileJSON)}
	tomlWriter := &TOMLWriter{filepath.Join(dir, dataOutFileTOML)}
	loadIntoFile(hotels, jsonWriter)
	loadIntoFile(hotels, tomlWriter)
}

/*
 * The checkArgs() function returns a string of file path and
 * error if there is any.
 */
func checkArgs() (string, bool, error) {
	//Fetch the command line arguments.
	args := os.Args

	//Check the length of the arugments, return failure if they are too
	//long or too short.
	if (len(args) < 2) || (len(args) > 3) {
		return "", false, errors.New("Invalid number of arguments. \n" +
			" Please provide sort option (-s) and the file name with relative path" +
			" of the csv data input file!\n")
	}
	file_path := args[1]
	var isSort bool
	if len(args) == 3 {
		if args[1] != "-s" {
			return "", false, errors.New("Invalid number of arguments. \n" +
				" Please provide sort option (-s) and the file name with relative path" +
				" of the csv data input file!\n")
		}
		isSort = true
		file_path = args[2]
	}
	//On success, return the file_path and isSort value
	return file_path, isSort, nil
}
