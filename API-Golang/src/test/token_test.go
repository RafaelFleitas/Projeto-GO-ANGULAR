package test

import (
	"testing"

	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/stretchr/testify/assert"
)

func TestRemoveBearerPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"com prefixo Bearer:", "Bearer: abc123", "abc123"},
		{"sem prefixo", "abc123", "abc123"},
		{"string vazia", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := model.RemoveBearerPrefix(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
