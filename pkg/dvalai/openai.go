package dvalai

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type DvalOpenAI struct {
	OpenAIToken string
}

type GetQAndAResponse interface {
	GetQAndAResponse(s string) (string, error)
}

func (d DvalOpenAI) GetQAndAResponse(s string) (string, error) {
	c := openai.NewClient(d.OpenAIToken)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:            openai.GPT3TextDavinci003,
		MaxTokens:        100,
		Temperature:      0,
		TopP:             1,
		FrequencyPenalty: 0.2,
		PresencePenalty:  0,
		Prompt:           "Is " + s + " a valid email address?\n",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("completion error: %v", err)
	}
	return resp.Choices[0].Text, nil
}

func ValidateEmailAI(d GetQAndAResponse, vResponse ValidatorResponse, v Validator, data string) ValidatorResponse {
	resp, err := d.GetQAndAResponse(data)
	vResponse.Rule = strings.TrimSpace(resp)
	if err != nil {
		vResponse.Error = true
		vResponse.ErrorMsg = err.Error()
		return vResponse
	}
	// check if resp starts with "Yes"
	if vResponse.Rule[:3] == "Yes" {
		vResponse.Valid = true
	} else {
		vResponse.Valid = false
	}
	return vResponse
}
