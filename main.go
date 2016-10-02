package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	dataInFile      = "hotels.csv"
	dataOutFileJSON = "hotels.json"
	dataOutFileTOML = "hotels.toml"
)

func main() {
	filePath, err := checkArgs()
	if err != nil {
		log.Printf("Error: %s\n", err)
		filePath = dataInFile
		log.Printf("No input file found...going to read: %s\n\n", filePath)
	}
	dir := filepath.Dir(filePath)
	fmt.Printf("dir is:%s\n", dir)
	hotels := extractFromCSVFile(filePath)

	jsonWriter := &JSONWriter{filepath.Join(dir, dataOutFileJSON)}
	tomlWriter := &TOMLWriter{filepath.Join(dir, dataOutFileTOML)}
	loadIntoFile(hotels, jsonWriter)
	loadIntoFile(hotels, tomlWriter)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

/*
 * The checkArgs() function returns a string of file path and
 * error if there is any.
 */
func checkArgs() (string, error) {
	//Fetch the command line arguments.
	args := os.Args

	//Check the length of the arugments, return failure if they are too
	//long or too short.
	if (len(args) < 2) || (len(args) >= 3) {
		return "", errors.New("Invalid number of arguments. \n" +
			"Please provide the file name with relative path of the words list input file!\n")
	}
	file_path := args[1]
	//On success, return the file_path value
	return file_path, nil
}
