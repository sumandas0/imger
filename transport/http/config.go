package http

import (
	"time"

	"github.com/RexterR/imger/pkg/errors"
)

// Configuration for running the server
type Configuration struct {
	Address      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Validate the server configuration
func (c *Configuration) Validate() error {
	if c.Address == "" {
		return errors.EValidation("Missing Address", nil)
	}

	return nil
}
