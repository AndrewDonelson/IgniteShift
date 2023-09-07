package syslist

import (
	"testing"
)

func TestOS_String(t *testing.T) {
	tests := []struct {
		os       OS
		expected string
	}{
		{OS_UNKNOWN, "unknown"},
		{OS_LINUX, "linux"},
		{OS_DARWIN, "darwin"},
		{OS_WINDOWS, "windows"},
		{OS(-1), "unknown"}, // Test an invalid OS value
	}

	for _, test := range tests {
		got := test.os.String()
		if got != test.expected {
			t.Errorf("For OS %d, expected %s but got %s", test.os, test.expected, got)
		}
	}
}

func TestArch_String(t *testing.T) {
	tests := []struct {
		arch     Arch
		expected string
	}{
		{ARCH_UNKNOWN, "unknown"},
		{ARCH_AMD64, "amd64"},
		{ARCH_ARM64, "arm64"},
		{Arch(-1), "unknown"}, // Test an invalid Arch value
	}

	for _, test := range tests {
		got := test.arch.String()
		if got != test.expected {
			t.Errorf("For Arch %d, expected %s but got %s", test.arch, test.expected, got)
		}
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		osarch   string
		expected bool
	}{
		{"linux/amd64", true},
		{"linux/arm64", true},
		{"darwin/amd64", true},
		{"darwin/arm64", true},
		{"windows/amd64", true},
		{"windows/arm64", true},
		{"linux/invalid", false},
		{"invalid/amd64", false},
		{"invalid/invalid", false},
		{"", false},
	}

	for _, test := range tests {
		got := Validate(test.osarch)
		if got != test.expected {
			t.Errorf("For osarch %s, expected %v but got %v", test.osarch, test.expected, got)
		}
	}
}
