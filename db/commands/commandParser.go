package commands

import (
	"errors"
	"strings"

	"github.com/deuce109/mercurydb/v2/common/collections"
)

var commandNames = []string{CONTAINS, INTERSECTS, PING, WITHIN, INSERT}

func ParseCommand(command string) (string, []string, error) {

	var foundCommand string
	var commandLine string
	commandLineUpper := strings.ToUpper(command)

	for _, c := range commandNames {
		if c == commandLineUpper {
			return c, []string{}, nil
		} else if len(c)+1 <= len(commandLineUpper) && commandLineUpper[0:len(c)+1] == c+" " {
			foundCommand = c
			commandLine = command[len(c)+1:]
			break
		}
	}

	if foundCommand == "" {
		return "", nil, errors.New("no command found in beginning of string")
	}

	tokens := tokenize(commandLine)

	return strings.ToUpper(foundCommand), tokens, nil

}

func tokenize(input string) []string {
	tokenSlice := make([]string, 0)
	groupingLevel := 0
	var token string

	resetToken := func() {
		tokenSlice = append(tokenSlice, token)
		token = ""
	}

	for _, char := range input {

		if char == '(' {
			groupingLevel += 1
			token += string(char)
		} else if char == ')' && groupingLevel > 0 {
			groupingLevel -= 1
			token += string(char)
		} else if char == '{' && groupingLevel == 0 {
			groupingLevel += 1
			resetToken()
			token += string(char)
		} else if char == '}' && groupingLevel == 1 {
			groupingLevel -= 1
			token += string(char)
			resetToken()
		} else if char == ' ' && groupingLevel == 0 {
			resetToken()
		} else {
			token += string(char)
		}
	}

	tokenSlice = append(tokenSlice, token)

	return collections.
		NewCollection(tokenSlice).
		FilterNonZero().
		ToSlice()
}
