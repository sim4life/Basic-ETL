package main

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

/* Match for url regexp */
var Regex = regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)

type Validator interface {
	Validate() bool
}

type NameValidator struct {
	name string
}

func (nv *NameValidator) Validate() bool {
	return utf8.ValidString(nv.name)
}

type URLValidator struct {
	url string
}

func (uv *URLValidator) Validate() bool {
	return Regex.MatchString(uv.url)
}

type StarsValidator struct {
	stars string
}

func (sv *StarsValidator) Validate() bool {
	vStars, err := strconv.Atoi(sv.stars)
	if err != nil {
		return false
	}
	if vStars < 0 || vStars > 5 {
		return false
	}
	return true
}

func validateData(validator Validator) bool {
	return validator.Validate()
}
