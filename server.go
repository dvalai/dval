package main

import (
	"dvalai/dval/pkg/dvalai"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/", validate)
	e.Logger.Fatal(e.Start(":1323"))
}

// Struct for post request
type PostRequest struct {
	Data       string             `json:"data"`
	Validators []dvalai.Validator `json:"validators"`
}

type PostResponse struct {
	ValidatorResponses []dvalai.ValidatorResponse `json:"validatorResponses"`
}

// OpenAI instance
var openaiInstance = dvalai.DvalOpenAI{
	OpenAIToken: os.Getenv("OPENAI_TOKEN"),
}

// e.POST("/", validate)
func validate(c echo.Context) error {
	// Get data and regex from request
	req := new(PostRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	resp := new(PostResponse)
	for _, v := range req.Validators {
		vResponse := dvalai.ValidatorResponse{Name: v.Name, Error: false}
		if v.Type == "regex" {
			resp.ValidatorResponses = append(resp.ValidatorResponses, dvalai.ValidateRegEx(vResponse, v, req.Data))
		} else if v.Type == "genregex" {
			resp.ValidatorResponses = append(resp.ValidatorResponses, dvalai.ValidateGenRegEx(openaiInstance, vResponse, v, req.Data))
		}
	}

	return c.JSON(http.StatusOK, resp)
}
