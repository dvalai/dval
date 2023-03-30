package dvalai

import "testing"

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
	if vResponse.Match != true {
		t.Errorf("Expected true, got %v", vResponse.Match)

	}
	vResponse = ValidateRegEx(vResponse, v, "abc123")
	if vResponse.Match != false {
		t.Errorf("Expected false, got %v", vResponse.Match)

	}
}

// Test for ValidateGenRegEx
func TestValidateGenRegEx(t *testing.T) {
	// Create a mock for GetRegularExpression
	openai := DvalOpenAI{OpenAIToken: "testtoken"}
	v := Validator{
		Name: "test",
		Type: "genregex",
		Rule: "email",
	}
	vResponse := ValidatorResponse{
		Name: v.Name,
		Rule: v.Rule,
	}
	vResponse = openai.ValidateGenRegEx(vResponse, v, "johndoe@gmail.com")
	if vResponse.Match != true {
		t.Errorf("Expected true, got %v", vResponse.Match)

	}
	vResponse = openai.ValidateGenRegEx(vResponse, v, "invalidemail@test@gmail.comabcd")
	if vResponse.Match != false {
		t.Errorf("Expected false, got %v", vResponse.Match)

	}
}
