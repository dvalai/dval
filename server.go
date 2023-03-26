package main

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/", validate)
	e.Logger.Fatal(e.Start(":1323"))
}

// Struct for post request
type PostRequest struct {
	Data  string `json:"data"`
	Regex string `json:"regex"`
}

type PostResponse struct {
	Match bool `json:"match"`
}

// e.POST("/", validate)
func validate(c echo.Context) error {
	// Get data and regex from request
	req := new(PostRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	m, err := regexp.MatchString(req.Regex, req.Data)
	if err != nil {
		return err
	}
	if m {
		return c.JSON(http.StatusOK, PostResponse{Match: true})
	}
	return c.JSON(http.StatusOK, PostResponse{Match: false})
}
