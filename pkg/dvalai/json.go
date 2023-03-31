package dvalai

import (
	"encoding/json"
)

func ValidateJSON(vResponse ValidatorResponse, v Validator, data string) ValidatorResponse {
	vResponse.Valid = json.Valid([]byte(data))
	return vResponse
}
