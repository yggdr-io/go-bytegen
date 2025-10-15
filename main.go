package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	sizeFlag := flag.String("size", "1MB", "Size of the file to generate (e.g., 1MB, 1024KB)")
	outFlag := flag.String("o", "random_bytes.bin", "Output file path")
	flag.Parse()

	size, err := parseSize(*sizeFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse size %q: %v\n", *sizeFlag, err)
		os.Exit(1)
	}

	out, err := os.Create(*outFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create %q: %v\n", *outFlag, err)
		os.Exit(1)
	}
	defer func() {
		if cerr := out.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "close %q: %v\n", *outFlag, cerr)
			os.Exit(1)
		}
	}()

	if err := writeRandom(out, size); err != nil {
		fmt.Fprintf(os.Stderr, "write random bytes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Random bytes file created successfully")
}

// writeRandom writes exactly n random bytes to w.
func writeRandom(w io.Writer, n int64) error {
	_, err := io.CopyN(w, rand.Reader, n)
	return err
}

// parseSize turns size strings into byte values.
func parseSize(s string) (int64, error) {
	cfg := map[string]int64{
		"KB": 1024,
		"MB": 1024 * 1024,
		"GB": 1024 * 1024 * 1024,
	}

	s = strings.TrimSpace(strings.ToUpper(s))

	if size, err := strconv.ParseInt(s, 10, 64); err == nil {
		return size, nil
	}

	for unit, mult := range cfg {
		if !strings.HasSuffix(s, unit) {
			continue
		}

		s1 := strings.TrimSuffix(s, unit)
		size, err := strconv.ParseInt(s1, 10, 64)
		if err != nil {
			return 0, err
		}
		return size * mult, nil
	}

	return 0, errors.New("unrecognized size format")
}
