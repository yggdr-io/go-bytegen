package main

import (
	"fmt"
	"os"
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

func TestWriteRandomBytes(t *testing.T) {
	testCases := []int64{
		2048,
		1536,
		0,
	}

	for i, size := range testCases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			t.Parallel()

			file, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer t.Cleanup(func() {
				os.Remove(file.Name())
			})

			if err := writeRandomBytes(file, size); err != nil {
				t.Fatalf("Failed to write random bytes: %v", err)
			}

			if err := file.Close(); err != nil {
				t.Fatalf("Failed to close the file: %v", err)
			}

			stat, err := os.Stat(file.Name())
			if err != nil {
				t.Fatalf("Failed to stat temp file: %v", err)
			}
			if stat.Size() != size {
				t.Errorf("Expected file size to be %d bytes, got %d bytes", size, stat.Size())
			}
		})
	}
}
