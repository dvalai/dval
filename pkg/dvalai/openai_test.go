package dvalai

import (
	"fmt"
	"testing"
)

type (
	mockDvalOpenAI struct {
		OpenAIToken string
		mockValid   bool
	}
)

func (mdoAI mockDvalOpenAI) GetQAndAResponse(s string) (string, error) {
	if mdoAI.mockValid {
		return "Yes, johndoe@gmail.com is a valid email address.", nil
	}
	return "No, invalidemail@something.thiscannotbethecase is not a valid email address.", nil
}

// Test for ValidateGenRegEx
func TestValidateGenRegEx(t *testing.T) {
	// Create a mock for GetRegularExpression
	openai := mockDvalOpenAI{OpenAIToken: "testtoken", mockValid: true}
	v := Validator{
		Name: "test",
		Type: "emailai",
	}
	vResponse := ValidatorResponse{
		Name: v.Name,
		Rule: v.Rule,
	}
	vResponse = ValidateEmailAI(openai, vResponse, v, "johndoe@gmail.com")
	if vResponse.Valid != true {
		fmt.Println(vResponse)
		t.Errorf("Expected true, got %v", vResponse.Valid)

	}
	openai.mockValid = false
	vResponse = ValidateEmailAI(openai, vResponse, v, "invalidemail@test@gmail.comabcd")
	if vResponse.Valid != false {
		t.Errorf("Expected false, got %v", vResponse.Valid)

	}
}
