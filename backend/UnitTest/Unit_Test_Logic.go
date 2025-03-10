package unittest

import (
	"testing"
)

func idIsTrue(id int) bool {
	var isTrue bool = true
	if id <= 144 {
		isTrue = false
	}
	return isTrue
}

func TestidIsTrue(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"ID less than or equal to 144", 100, false},
		{"ID greater than 144", 150, true},
		{"ID equal to 144", 144, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := idIsTrue(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
