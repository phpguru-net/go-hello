package main

import (
	"strings"
)

func AddWelcomeBorder(welcomeMsg string, numStarsPerLine int) string {
	var welcomeBorder = strings.Repeat("*", numStarsPerLine)
	return welcomeBorder + "\n" + welcomeMsg + "\n" + welcomeBorder
}

func CleanUpMessage(message string) string {
	output := strings.ReplaceAll(message, "*", "")
	output = strings.ReplaceAll(output, "\n", "")
	output = strings.TrimLeft(output, " ")
	output = strings.TrimRight(output, " ")
	return output
}
