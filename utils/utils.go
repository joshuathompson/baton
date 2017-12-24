package utils

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// StringInSlice determines whether a given string is within a slice of strings
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// MillisecondsToFormattedTime turns an input of milliseconds into mm:ss format
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

// LeftPaddedString pads a string on the left by a specified amount and pads the string on the right to fill the maxLength
func LeftPaddedString(value string, maxValueLength, padAmount int) string {
	valueLength := utf8.RuneCountInString(value)
	if maxValueLength-padAmount >= valueLength {
		return strings.Repeat(" ", padAmount) + value + strings.Repeat(" ", maxValueLength-valueLength-padAmount)
	} else if maxValueLength-padAmount < valueLength {
		tmp := strings.Trim(value[0:maxValueLength-padAmount-3], " ") + "..."
		tmpLength := utf8.RuneCountInString(tmp)
		return strings.Repeat(" ", padAmount) + tmp + strings.Repeat(" ", maxValueLength-tmpLength-padAmount)
	}

	return value
}
