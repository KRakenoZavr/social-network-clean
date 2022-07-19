package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const alphaNumeric = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
const onlyAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const minAge = 10
const maxAge = 150

type Validator struct {
	err []error
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) MustBeGreaterThan(high, value int) {
	if high > value {
		v.err = append(v.err, fmt.Errorf("%v should be greater than %v", value, high))
	}
}

func (v *Validator) CheckEmail(email string) {
	emailRegex := regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")
	if !emailRegex.MatchString(email) {
		v.err = append(v.err, fmt.Errorf("email: %v is not correct", email))
	}
}

func (v *Validator) CheckLen(word string, wordLen int) {
	if len(word) < wordLen {
		v.err = append(v.err, fmt.Errorf("%v should contain at lest %v characters", word, wordLen))
	}
}

func (v *Validator) CheckLenMultiple(wordLen int, words ...string) {
	for _, word := range words {
		v.CheckLen(word, wordLen)
	}
}

func (v *Validator) CheckSymbols(word string) {
	for _, l := range word {
		if !strings.ContainsAny(string(l), alphaNumeric) {
			v.err = append(v.err, fmt.Errorf("%s should contain only alpha numeric symbols", word))
		}
	}
}

func (v *Validator) CheckSymbolsMultiple(words ...string) {
	for _, word := range words {
		v.CheckSymbols(word)
	}
}

func (v *Validator) OnlyAlphabet(word string) {
	for _, l := range word {
		if !strings.ContainsAny(string(l), onlyAlphabet) {
			v.err = append(v.err, fmt.Errorf("%s should contain only alphabet characters", word))
		}
	}
}

func (v *Validator) CheckPass(password string) {
	v.CheckLen(password, 6)
}

func (v *Validator) CheckNull(word, field string) {
	if len(word) == 0 {
		v.err = append(v.err, fmt.Errorf("field: %s should not be empty", field))
	}
}

func (v *Validator) CheckBDay(date time.Time) {
	dateNow := time.Now()

	dy := date.Year()
	dny := dateNow.Year()

	if dny > dy+maxAge {
		v.err = append(v.err, fmt.Errorf("you are too old man"))
		return
	}

	if dny < dy+minAge {
		v.err = append(v.err, fmt.Errorf("you should be at least 10 y.o., you are too young man"))
		return
	}
}

func (v *Validator) Errors() []error {
	if len(v.err) > 0 {
		return v.err
	}
	return nil
}
