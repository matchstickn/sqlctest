package auth

import (
	"context"
	"strings"
)

// CustomClaims contains custom data from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Does nothing, but needs to compile
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// HasScope checks whether the claims have a specific scope.
func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}
