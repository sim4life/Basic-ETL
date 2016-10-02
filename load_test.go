package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

type MockWriter struct {
	FilePath string
}

func (mw *MockWriter) GetFilePath() string {
	return mw.FilePath
}

func (mw *MockWriter) GetDataBytes(data DataWriter) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, data.GetWritableData().Hotels)
	if err != nil {
		return nil, err
	}
	fmt.Printf("% x", buffer.Bytes())
	return buffer.Bytes(), nil
}

func (mw *MockWriter) ReadDataBytes() (*Hotels, error) {
	// buffer := bytes.NewBuffer(buffer.Bytes())
	buffer := new(bytes.Buffer)
	hotels := &Hotels{}
	err := binary.Read(buffer, binary.LittleEndian, hotels)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func Test_loadIntoFile(t *testing.T) {
	const dataOutFileMock = "hotels.mock"
	exp_hotels := &Hotels{[]Hotel{{"HollowMan", "321 Gotham city", 2, "BoogeyMan", "324-099-334", "http://hole.wall"}}}
	exp_len := len(exp_hotels.Hotels)
	mockWriter := &MockWriter{dataOutFileMock}
	loadIntoFile(exp_hotels, mockWriter)

	act_hotels, _ := mockWriter.ReadDataBytes()
	act_len := len(act_hotels.Hotels)
	if act_len != exp_len {
		t.Errorf("Failed with expected length:%d and actual length:%d\n", exp_len, act_len)
	}
}
