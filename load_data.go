package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type FileFormatWriter interface {
	GetFilePath() string
	GetDataBytes(DataWriter) ([]byte, error)
}

type TOMLWriter struct {
	FilePath string
}

func (tw *TOMLWriter) GetFilePath() string {
	return tw.FilePath
}

func (tw *TOMLWriter) GetDataBytes(data DataWriter) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(data.GetWritableData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type JSONWriter struct {
	FilePath string
}

func (jw *JSONWriter) GetFilePath() string {
	return jw.FilePath
}

func (jw *JSONWriter) GetDataBytes(data DataWriter) ([]byte, error) {
	jsonData, err := json.Marshal(data.GetWritableData())
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func loadIntoFile(hotelsData DataWriter, fileFormatWriter FileFormatWriter) {
	fileHandle, err := os.OpenFile(fileFormatWriter.GetFilePath(), os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err)
	defer fileHandle.Close()

	bytesData, err := fileFormatWriter.GetDataBytes(hotelsData)
	if err != nil {
		log.Println(err)
	}
	fileHandle.Write(bytesData)
	fileHandle.Sync()
}
