package main

import "testing"

func Test_extractFromCSVFile(t *testing.T) {

	filePath := "data_files/hotels.csv"

	exp_len := 4000
	hotels := extractFromCSVFile(filePath)
	act_len := len(hotels.Hotels)

	if act_len != exp_len {
		t.Errorf("Failed with expected length:%d and actual length:%d\n", exp_len, act_len)
	}
}

func Test_parseToStructValid1(t *testing.T) {
	rawData := [][]string{{"WorldMan", "123 Globe street", "4", "FooMan", "123-456-789", "http://premier.ch/about"}}
	exp_len := 1
	hotels := parseToStruct(rawData)
	act_len := len(hotels.Hotels)

	if act_len != exp_len {
		t.Errorf("Failed with expected length:%d and actual length:%d\n", exp_len, act_len)
	}
}

func Test_parseToStructValid2(t *testing.T) {
	rawData := [][]string{{"WorldMan", "123 Globe street", "4", "FooMan", "123-456-789", "http://premier.ch/about"},
		{"BazMan", "13 Life street", "0", "JazzMan", "123-987-789", "http://prime.ch/about"}}
	exp_len := 2
	hotels := parseToStruct(rawData)
	act_len := len(hotels.Hotels)

	if act_len != exp_len {
		t.Errorf("Failed with expected length:%d and actual length:%d\n", exp_len, act_len)
	}
}

func Test_parseToStructInValid1(t *testing.T) {
	rawData := [][]string{{"WorldMan", "123 Globe street", "4", "FooMan", "123-456-789", "http://premier.ch/about"},
		{"BazMan", "13 Life street", "x", "JazzMan", "123-987-789", "http://prime.ch/about"}}
	exp_len := 1
	hotels := parseToStruct(rawData)
	act_len := len(hotels.Hotels)

	if act_len != exp_len {
		t.Errorf("Failed with expected length:%d and actual length:%d\n", exp_len, act_len)
	}
}

func Test_parseToStructInValid2(t *testing.T) {
	rawData := [][]string{{"WorldMan", "123 Globe street", "4", "FooMan", "123-456-789", "http://premier.ch/about"},
		{"\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98", "13 Life street", "4", "JazzMan", "123-987-789", "http://prime.ch/about"}}
	exp_len := 1
	hotels := parseToStruct(rawData)
	act_len := len(hotels.Hotels)

	if act_len != exp_len {
		t.Errorf("Failed with expected length:%d and actual length:%d\n", exp_len, act_len)
	}
}

func Test_parseToStructInValid3(t *testing.T) {
	rawData := [][]string{{"WorldMan", "123 Globe street", "4", "FooMan", "123-456-789", "//premier.ch/about"}}
	exp_len := 0
	hotels := parseToStruct(rawData)
	act_len := len(hotels.Hotels)

	if act_len != exp_len {
		t.Errorf("Failed with expected length:%d and actual length:%d\n", exp_len, act_len)
	}
}
