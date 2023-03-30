package dvalai

import (
	"context"
	"fmt"
	"regexp"

	openai "github.com/sashabaranov/go-openai"
)

type DvalOpenAI struct {
	OpenAIToken string
}

func (d DvalOpenAI) GetRegularExpression(s string) (string, error) {
	c := openai.NewClient(d.OpenAIToken)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:            openai.GPT3TextDavinci003,
		MaxTokens:        100,
		Temperature:      0,
		TopP:             1,
		FrequencyPenalty: 0.2,
		PresencePenalty:  0,
		Prompt:           "Regular expression for " + s,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("completion error: %v", err)
	}
	return resp.Choices[0].Text, nil
}

// Struct for ValidatorRequest
type Validator struct {
	Name string `json:"name"`
	Type string `json:"type"` // "regex" or "genregex"
	Rule string `json:"rule"`
}

// Struct for ValidatorResponse
type ValidatorResponse struct {
	Name     string `json:"name"`
	Error    bool   `json:"error"`
	ErrorMsg string `json:"errorMsg"`
	Rule     string `json:"rule"`
	Match    bool   `json:"match"`
}

func ValidateRegEx(vResponse ValidatorResponse, v Validator, data string) ValidatorResponse {
	vResponse.Rule = v.Rule
	m, err := regexp.MatchString(v.Rule, data)
	if err != nil {
		vResponse.Error = true
		vResponse.ErrorMsg = err.Error()
		return vResponse
	}
	vResponse.Match = m
	return vResponse
}

func (d DvalOpenAI) ValidateGenRegEx(vResponse ValidatorResponse, v Validator, data string) ValidatorResponse {
	generatedRegex, err := d.GetRegularExpression(v.Rule)
	vResponse.Rule = generatedRegex
	if err != nil {
		vResponse.Error = true
		vResponse.ErrorMsg = err.Error()
		return vResponse
	}
	return ValidateRegEx(vResponse, v, data)
}
