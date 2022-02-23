package imger

// Parameter contains all properties of a single effect parameter
type Parameter struct {
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Type        string      `json:"type"`
	Example     interface{} `json:"example"`
	Default     interface{} `json:"default,omitempty"`
	Values      interface{} `json:"values,omitempty"`
}

// Parameters it's a map that contains all parameters of an effect
type Parameters map[string]Parameter

// Effect represents an image transformation (ex: rotate, resize, overlay...)
type Effect interface {
	// ID that identifies the effect
	ID() string
}
