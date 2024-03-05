package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sizeFlag := flag.String("size", "1MB", "Size of the file to generate (e.g., 1MB, 1024KB)")
	flag.Parse()

	size, err := parseSize(*sizeFlag)
	if err != nil {
		fmt.Printf("Error parsing size: %v\n", err)
		return
	}

	out, err := os.Create("random_bytes.bin")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer out.Close()

	if err := gen(out, size); err != nil {
		fmt.Printf("Error writing random bytes: %v\n", err)
		return
	}

	fmt.Println("Random bytes file created successfully")
}

func parseSize(s string) (int64, error) {
	units := map[string]int64{
		"KB": 1024,
		"MB": 1024 * 1024,
		"GB": 1024 * 1024 * 1024,
	}

	s = strings.ToUpper(s)
	for unit, multiplier := range units {
		if strings.HasSuffix(s, unit) {
			size, err := strconv.ParseInt(strings.TrimSuffix(s, unit), 10, 64)
			if err != nil {
				return 0, err
			}
			return size * multiplier, nil
		}
	}

	return strconv.ParseInt(s, 10, 64)
}

// gen populates file f with n bytes of random data.
func gen(f *os.File, n int64) error {
	b := make([]byte, 1024)
	for n > 0 {
		ni := min(n, int64(len(b)))
		_, err := rand.Read(b[:ni])
		if err != nil {
			return err
		}
		_, err = f.Write(b[:ni])
		if err != nil {
			return err
		}
		n -= ni
	}

	return nil
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
