package dvalai

import (
	"regexp"
)

func ValidateRegEx(vResponse ValidatorResponse, v Validator, data string) ValidatorResponse {
	vResponse.Rule = v.Rule
	m, err := regexp.MatchString(v.Rule, data)
	if err != nil {
		vResponse.Error = true
		vResponse.ErrorMsg = err.Error()
		return vResponse
	}
	vResponse.Valid = m
	return vResponse
}
