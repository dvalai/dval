package dvalai

// Struct for ValidatorRequest
type Validator struct {
	Name string `json:"name"`
	Type string `json:"type"` // "regex", "json" or "emailai"
	Rule string `json:"rule,omitempty"`
}

// Struct for ValidatorResponse
type ValidatorResponse struct {
	Name     string `json:"name"`
	Error    bool   `json:"error"`
	ErrorMsg string `json:"errorMsg"`
	Rule     string `json:"rule,omitempty"`
	Valid    bool   `json:"valid"`
}
