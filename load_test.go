package main

import (
	"encoding/json"
	"testing"
)

type MockWriter struct {
	FilePath string
}

func (mw *MockWriter) GetFilePath() string {
	return mw.FilePath
}

func (mw *MockWriter) GetDataBytes(data DataWriter) ([]byte, error) {
	jsonData, err := json.Marshal(data.GetWritableData())
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func Test_loadIntoFile(t *testing.T) {
	const dataOutFileMock = "hotels.mock"
	exp_hotels := &Hotels{[]Hotel{{"HollowMan", "321 Gotham city", 2, "BoogeyMan", "324-099-334", "http://hole.wall"}}}
	var exp_err error = nil
	mockWriter := &MockWriter{dataOutFileMock}
	act_err := loadIntoFile(exp_hotels, mockWriter)

	if act_err != exp_err {
		t.Errorf("Failed with expected error:%s and actual error:%s\n", exp_err, act_err)
	}
}
