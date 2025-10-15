package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestParseSize ensures parseSize correctly translates size strings,
// including units, to byte values.
func TestParseSize(t *testing.T) {
	testCases := []struct {
		s    string // Input size as string
		want int64  // Expected result in bytes
	}{
		{"512", 512}, // No unit, assumed bytes
		{"1KB", 1024},
		{"10MB", 10 * 1024 * 1024},
		{"2GB", 2 * 1024 * 1024 * 1024},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			t.Parallel()

			got, err := parseSize(tc.s)
			if err != nil {
				t.Fatalf("Got unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Expected %d bytes for input %q, but got %d bytes", tc.want, tc.s, got)
			}
		})
	}
}

// TestParseSize_Error tests the function parseSize() with inputs expected
// to cause errors.
func TestParseSize_Error(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{"abc", "unrecognized size format"},
		{"NKB", "parsing \"N\": invalid syntax"},
		{"1TB", "unrecognized size format"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			t.Parallel()

			_, err := parseSize(tc.s)
			if err == nil {
				t.Fatalf("parseSize(%s) expected an error, but got none", tc.s)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Errorf("parseSize(%s) expected error message to contain %q, got %q", tc.s, tc.want, err.Error())
			}
		})
	}
}
