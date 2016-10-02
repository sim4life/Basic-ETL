package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/BurntSushi/toml"
)

const (
	dataInFile      = "hotels.csv"
	dataOutFileJSON = "hotels.json"
	dataOutFileTOML = "hotels.toml"
)

/* Match for url regexp */
var Regex = regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)

type Hotels struct {
	Hotels []Hotel
}

type Hotel struct {
	Name    string
	Address string
	Stars   int
	Contact string
	Phone   string
	Uri     string
}

func main() {
	filePath, err := checkArgs()
	if err != nil {
		fmt.Printf("Error: %s", err)

		filePath = dataInFile
		fmt.Printf("No input file found...going to read: %s\n\n", filePath)

	}
	dir := filepath.Dir(filePath)
	fmt.Printf("dir is:%s\n", dir)
	hotels := extractFromCSVFile(filePath)
	// sanity check, display to standard output
	for i, hotel := range hotels.Hotels {
		if i > 0 && i <= 5 {
			fmt.Printf("[%d]-name: %s, address: %s, stars: %s, contact: %s, phone: %s, uri: %s.\n", i, hotel.Name, hotel.Address, hotel.Stars, hotel.Contact, hotel.Phone, hotel.Uri)

		}
	}

	loadIntoFile(hotels, "json", filepath.Join(dir, dataOutFileJSON))
	loadIntoFile(hotels, "toml", filepath.Join(dir, dataOutFileTOML))
}

func extractFromCSVFile(fileName string) Hotels {

	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file(%s): %s", fileName, err)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	//match the rest of the records with the number of records in the first line
	reader.FieldsPerRecord = 0

	//reading first record to ignore the data titles line
	record, err := reader.Read()
	if err == io.EOF {
		log.Fatalf("No data found in file: %s", fileName)
	}
	if err != nil {
		log.Fatalf("Error reading file(%s): %s", fileName, err)
	}

	fmt.Println(record)

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading file(%s): %s", fileName, err)
	}

	hotels := parseToStruct(rawCSVData)

	fmt.Printf("\nhotels read are len:%d\n", len(hotels.Hotels))
	return hotels
}

func parseToStruct(rawData [][]string) Hotels {
	hotels := make([]Hotel, 0)
	for _, each := range rawData {
		/*
			if i <= 5 {
				fmt.Printf("[%d]-name: %s, address: %s, stars: %s, contact: %s, phone: %s, uri: %s.\n", i, each[0], each[1], each[2], each[3], each[4], each[5])

			}*/

		name, stars, uri, err := validateData(each[0], each[2], each[5])
		if err == nil {
			hotel := Hotel{Name: name, Address: each[1], Stars: stars, Contact: each[3], Phone: each[4], Uri: uri}
			hotels = append(hotels, hotel)
		}
	}

	return Hotels{hotels}
}

func loadIntoFile(hotelsData Hotels, format, filePath string) {
	fileHandle, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err)
	defer fileHandle.Close()

	if format == "json" {
		jsonData, err := json.Marshal(hotelsData)
		if err != nil {
			fmt.Println(err)
		}
		fileHandle.Write(jsonData)
		fileHandle.Sync()
	} else if format == "toml" {
		fmt.Println("Writing toml file")
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(hotelsData); err != nil {
			fmt.Println(err)
		}
		fmt.Println(buf.String())
		fileHandle.Write(buf.Bytes())
		fileHandle.Sync()
	}
}

func validateData(name, stars, uri string) (string, int, string, error) {
	vName := name
	if !utf8.ValidString(name) {
		return "", 0, "", errors.New("Name is invalid UTF-8 string")
	}
	vStars, err := strconv.Atoi(stars)
	if err != nil {
		return "", 0, "", err
	}
	if vStars < 0 || vStars > 5 {
		return "", 0, "", errors.New("Stars is -ve or more than 5")
	}

	vUri := uri
	if !validateURL(uri) {
		return "", 0, "", errors.New("URI does not match the regex of URL")
	}
	return vName, vStars, vUri, nil
}

func validateURL(url string) bool {
	return Regex.MatchString(url)
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
