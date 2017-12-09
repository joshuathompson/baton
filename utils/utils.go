package utils

import (
	"strconv"
	"strings"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func MillisecondsToFormattedTime(i int) string {
	var output []string
	totalSeconds := i / 1000

	minutes := (totalSeconds) / 60
	seconds := (totalSeconds % 60) % 60

	output = append(output, strconv.Itoa(minutes))

	if seconds < 10 {
		output = append(output, "0"+strconv.Itoa(seconds))
	} else {
		output = append(output, strconv.Itoa(seconds))
	}

	return strings.Join(output, ":")
}
