package utils

import (
	"fmt"
	"regexp"
	"strings"
)

const alphaNumeric = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
const onlyAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func (v *Validator) VerifyEmail(email string) {
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

func (v *Validator) CheckSymbols(word string) {
	for _, l := range word {
		if !strings.ContainsAny(string(l), alphaNumeric) {
			v.err = append(v.err, fmt.Errorf("%s should contain only alpha numeric symbols", word))
		}
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

func (v *Validator) Check() []error {
	if len(v.err) > 0 {
		return v.err
	}
	return nil
}
