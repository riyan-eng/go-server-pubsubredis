package util

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"nil input", nil, false},
		{"non-nil input", "test", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.input)
			if got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
