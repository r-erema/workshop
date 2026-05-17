package encodeanddecodestrings

import (
	"fmt"
	"strconv"
	"strings"
)

// Encode encodes a list of strings to a single string.
// Time O(N) where N is each string
// Space O(1), since we don't use any additional space.
func Encode(input []string) string {
	var (
		encoded     string
		encodedSb47 strings.Builder
	)
	for _, s := range input {
		_, _ = fmt.Fprintf(&encodedSb47, "%d#%s", len(s), s)
	}

	encoded += encodedSb47.String()

	return encoded
}

// Decode decodes a single string to a list of strings.
// Time O(N) where N is each string
// Space O(N), since we use array to collect sliced words.
func Decode(input string) []string {
	var (
		lengthStr, word string
		res             []string
	)

	for input != "" {
		i := 0
		for input[i] != '#' {
			i++
		}

		lengthStr, input = input[:i], input[i:]

		length, _ := strconv.Atoi(lengthStr)

		word, input = input[1:length+1], input[length+1:]

		res = append(res, word)
	}

	return res
}
