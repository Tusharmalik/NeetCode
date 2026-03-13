package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var sizes []string
	var encode string

	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}
	encode = strings.Join(sizes, ",")
	encode = encode + "#" + strings.Join(strs, "")

	return encode
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	parts := strings.SplitN(encoded, "#", 2)
	if len(parts) != 2 {
		return []string{}
	}

	sizesPart, dataPart := parts[0], parts[1]
	sizeStrs := strings.Split(sizesPart, ",")
	var decoded []string

	prevSize := 0

	for _, sizeStr := range sizeStrs {
		size, _ := strconv.Atoi(sizeStr)

		decoded = append(decoded, dataPart[prevSize:prevSize+size])
		prevSize = prevSize + size
	}
	return decoded
}

func main() {
	solution := Solution{}
	encoded := solution.Encode([]string{"Go💙Go💙", "Go💙"})
	decode := solution.Decode(encoded)
	fmt.Printf("%#v\n", decode)
}
