package dvalai

import (
	"testing"
)

// Test for ValidateRegEx
func TestValidateRegEx(t *testing.T) {
	v := Validator{
		Name: "test",
		Type: "regex",
		Rule: "^[a-z]+$",
	}
	vResponse := ValidatorResponse{
		Name: v.Name,
		Rule: v.Rule,
	}
	vResponse = ValidateRegEx(vResponse, v, "abc")
	if vResponse.Valid != true {
		t.Errorf("Expected true, got %v", vResponse.Valid)

	}
	vResponse = ValidateRegEx(vResponse, v, "abc123")
	if vResponse.Valid != false {
		t.Errorf("Expected false, got %v", vResponse.Valid)

	}
}
