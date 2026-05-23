package helpers

import (
	"os"
	"testing"
)

// TestIsCliModeActive tests the IsCliModeActive helper which detects
// if the program is running in CLI mode (--version, --health, etc.)
// to allow early exit from init() functions without requiring DB.
func TestIsCliModeActive(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected bool
	}{
		{
			name:     "no args",
			args:     []string{"./shadow"},
			expected: false,
		},
		{
			name:     "--version flag",
			args:     []string{"./shadow", "--version"},
			expected: true,
		},
		{
			name:     "-version flag",
			args:     []string{"./shadow", "-version"},
			expected: true,
		},
		{
			name:     "-v flag",
			args:     []string{"./shadow", "-v"},
			expected: true,
		},
		{
			name:     "--health flag",
			args:     []string{"./shadow", "--health"},
			expected: true,
		},
		{
			name:     "-health flag",
			args:     []string{"./shadow", "-health"},
			expected: true,
		},
		{
			name:     "other flags",
			args:     []string{"./shadow", "--help"},
			expected: false,
		},
		{
			name:     "version with extra args",
			args:     []string{"./shadow", "--version", "extra"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			// Set test args
			os.Args = tt.args

			result := IsCliModeActive()
			if result != tt.expected {
				t.Errorf("IsCliModeActive() = %v, want %v", result, tt.expected)
			}
		})
	}
}
