package dvalai

import (
	"fmt"
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

type (
	mockDvalOpenAI struct {
		OpenAIToken string
	}
)

func (mdoAI mockDvalOpenAI) GetRegularExpression(s string) (string, error) {
	return "^[a-z0-9._%+-]+@[a-z0-9.-]+.[a-z]{2,4}$", nil
}

// Test for ValidateGenRegEx
func TestValidateGenRegEx(t *testing.T) {
	// Create a mock for GetRegularExpression
	openai := mockDvalOpenAI{OpenAIToken: "testtoken"}
	v := Validator{
		Name: "test",
		Type: "genregex",
		Rule: "email",
	}
	vResponse := ValidatorResponse{
		Name: v.Name,
		Rule: v.Rule,
	}
	vResponse = ValidateGenRegEx(openai, vResponse, v, "johndoe@gmail.com")
	if vResponse.Valid != true {
		fmt.Println(vResponse)
		t.Errorf("Expected true, got %v", vResponse.Valid)

	}
	vResponse = ValidateGenRegEx(openai, vResponse, v, "invalidemail@test@gmail.comabcd")
	if vResponse.Valid != false {
		t.Errorf("Expected false, got %v", vResponse.Valid)

	}
}
