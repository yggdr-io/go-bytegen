package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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

func TestGen(t *testing.T) {
	testCases := []int64{
		2048,
		1536,
		0,
	}

	for i, n := range testCases {
		t.Run(fmt.Sprintf("Case%d", i), func(t *testing.T) {
			t.Parallel()

			f, err := os.CreateTemp("", "test")
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer t.Cleanup(func() {
				os.Remove(f.Name())
			})

			if err := gen(f, n); err != nil {
				t.Fatalf("Failed to write random bytes: %v", err)
			}

			if err := f.Close(); err != nil {
				t.Fatalf("Failed to close the temp file: %v", err)
			}

			stat, err := os.Stat(f.Name())
			if err != nil {
				t.Fatalf("Failed to stat temp file: %v", err)
			}
			if stat.Size() != n {
				t.Errorf("Expected file size to be %d bytes, got %d bytes", n, stat.Size())
			}
		})
	}
}

func TestGen_Random(t *testing.T) {
	n := int64(1024)

	genb := func() ([]byte, error) {
		f, err := os.CreateTemp("", "test")
		if err != nil {
			return nil, err
		}
		defer os.Remove(f.Name())

		if err := gen(f, n); err != nil {
			return nil, err
		}

		if err := f.Close(); err != nil {
			return nil, err
		}

		return os.ReadFile(f.Name())
	}

	b1, err := genb()
	if err != nil {
		t.Fatalf("Error in first generation/read: %v", err)
	}

	b2, err := genb()
	if err != nil {
		t.Fatalf("Error in second generation/read: %v", err)
	}

	if bytes.Equal(b1, b2) {
		t.Errorf("Expected gen to generate different random bytes, but both files are identical")
	}
}
