package gt7

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"Go", "oG"},
	}

	for _, tt := range tests {
		result := Reverse(tt.input)
		if result != tt.expected {
			t.Errorf("Reverse(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"madam", true},
	}

	for _, tt := range tests {
		result := IsPalindrome(tt.input)
		if result != tt.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
