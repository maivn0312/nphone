package main

import (
	"regexp"
)

type Phone struct {
	Code   string
	Number string
}

var codes = map[string]string{
	"7": "+7",
	"8": "+7",
}

func Parse(rawPhone string) (*Phone, bool) {
	phone := new(Phone)
	digits := determinate(rawPhone)
	if len(digits) < 11 {
		return phone, false
	}
	phone.setCode(digits[:len(digits)-10])
	phone.Number = digits[len(digits)-10:]
	return phone, phone.IsPhone()
}

func (phone *Phone) setCode(rawCode string) {
	if _, ok := codes[rawCode]; ok {
		phone.Code = codes[rawCode]
	}
}

func (phone *Phone) String() string {
	if (phone.IsPhone()) {
		return phone.Code + phone.Number
	}
	return ""
}

func (phone *Phone) IsPhone() bool {
	if inArray(codes, phone.Code) && len(phone.Number) == 10 {
		return true
	}
	return false
}

func determinate(rawPhone string) string {
	re := regexp.MustCompile(`\D+`)
	return re.ReplaceAllString(rawPhone, "")
}

func inArray(values map[string]string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
