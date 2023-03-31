package dvalai

import "testing"

// Test for dvalai.ValidateJSON
func TestValidateJSON(t *testing.T) {
	v := Validator{
		Name: "test",
		Type: "json",
		Rule: "",
	}
	vResponse := ValidatorResponse{
		Name: v.Name,
		Rule: v.Rule,
	}
	vResponse = ValidateJSON(vResponse, v, `{"name": "John Doe", "age": 25}`)
	if vResponse.Valid != true {
		t.Errorf("Expected true, got %v", vResponse.Valid)

	}
	vResponse = ValidateJSON(vResponse, v, `{"name": "John Doe", "age": 25`)
	if vResponse.Valid != false {
		t.Errorf("Expected false, got %v", vResponse.Valid)

	}
}
