package printErrors

import (
	"fmt"
	"strings"
)

func QuoteError(textAfterCommand string) bool {

	if !strings.Contains(textAfterCommand, "\"") {
		fmt.Println("")
		fmt.Println("Task must be wrapped in quotes")
		fmt.Println("")
		
		return true
	}

	return false
}

func ClosingQuoteError(start, end int ) bool {

	if start == end {
		fmt.Println("")
		fmt.Println("Task must be wrapped in quotes")
		fmt.Println("")

		return true
	}

	return false
}

func BadCommand(invalidCommand string) bool {
	if invalidCommand != "" {
		fmt.Println("")
        fmt.Println(invalidCommand, "is not a valid argument")
        fmt.Println("")

		return true
	}

	return false
}