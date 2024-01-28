package main

import (
	"fmt"
	"testing"
)

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
		tc := tc
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			t.Parallel()

			got, err := parseSize(tc.s)
			if err != nil {
				t.Fatalf("Got unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("Expected %d bytes for input '%s', but got %d bytes", tc.want, tc.s, got)
			}
		})
	}
}
