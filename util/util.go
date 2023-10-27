package util

import (
	"strconv"
	"strings"
)

// PrepareMapResponse prepares the response by converting the map[int]int to a string
func PrepareMapResponse(m map[int]int) string {
	var builder strings.Builder
	builder.WriteString("{")

	first := true
	for k, v := range m {
		if !first {
			builder.WriteString(", ")
		}
		builder.WriteString(strconv.Itoa(k))
		builder.WriteString(": ")
		builder.WriteString(strconv.Itoa(v))
		first = false
	}

	builder.WriteString("}")
	return builder.String()
}

// PrepareSliceResponse prepares the response by converting the []int to a string
func PrepareSliceResponse(s []int) string {
	strSlice := make([]string, len(s))

	for k, v := range s {
		strSlice[k] = strconv.Itoa(v)
	}

	return strings.Join(strSlice, ", ")
}
