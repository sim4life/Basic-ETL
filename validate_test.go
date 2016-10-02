package main

import "testing"

func Test_validName(t *testing.T) {
	const val = "HelloMan"
	nameValidator := &NameValidator{val}
	exp_val := true
	act_val := validateData(nameValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}

func Test_invalidName(t *testing.T) {
	const val = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	nameValidator := &NameValidator{val}
	exp_val := false
	act_val := validateData(nameValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}

func Test_validStars1(t *testing.T) {
	const val = "5"
	starsValidator := &StarsValidator{val}
	exp_val := true
	act_val := validateData(starsValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}
func Test_validStars2(t *testing.T) {
	const val = "0"
	starsValidator := &StarsValidator{val}
	exp_val := true
	act_val := validateData(starsValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}

func Test_invalidStars1(t *testing.T) {
	const val = "6"
	starsValidator := &StarsValidator{val}
	exp_val := false
	act_val := validateData(starsValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}
func Test_invalidStars2(t *testing.T) {
	const val = "-1"
	starsValidator := &StarsValidator{val}
	exp_val := false
	act_val := validateData(starsValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}
func Test_invalidStars3(t *testing.T) {
	const val = "x"
	starsValidator := &StarsValidator{val}
	exp_val := false
	act_val := validateData(starsValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}

func Test_validURL(t *testing.T) {
	const val = "http://premier.de/about"
	urlValidator := &URLValidator{val}
	exp_val := true
	act_val := validateData(urlValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}

func Test_invalidURL(t *testing.T) {
	const val = "6"
	urlValidator := &URLValidator{val}
	exp_val := false
	act_val := validateData(urlValidator)

	if act_val != exp_val {
		t.Errorf("Failed:%s with expected value:%t and actual value:%t\n", val, exp_val, act_val)
	}
}
