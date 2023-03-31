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

type GetRegularExpression interface {
	GetRegularExpression(s string) (string, error)
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

func ValidateGenRegEx(d GetRegularExpression, vResponse ValidatorResponse, v Validator, data string) ValidatorResponse {
	generatedRegex, err := d.GetRegularExpression(v.Rule)
	v.Rule = generatedRegex
	if err != nil {
		vResponse.Error = true
		vResponse.ErrorMsg = err.Error()
		return vResponse
	}
	return ValidateRegEx(vResponse, v, data)
}
