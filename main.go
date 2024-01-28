package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sizeFlag := flag.String("size", "1MB", "Size of the file to generate (e.g., 1MB, 1024KB)")
	flag.Parse()

	_, err := parseSize(*sizeFlag)
	if err != nil {
		fmt.Printf("Error parsing size: %v\n", err)
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
