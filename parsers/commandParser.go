package parsers

func ParseCommand(command string) (string, []any, error) {
	commandResult := command
	argCount := 0

	args := make([]any, argCount)

	return commandResult, args, nil

}

func removeEmptyStrings(slice []string) []string {
	results := make([]string, 0)
	for _, str := range slice {
		if str != "" && str != " " {
			results = append(results, str)

		}
	}
	return results
}

func NewTokenize(input string) ([]string, bool) {
	tokenSlice := make([]string, 0)
	groupingLevel := 0
	var token string
	data := false

	resetToken := false
	end := false

	for i, char := range input {

		if char == '(' {
			groupingLevel += 1
			resetToken = true
		} else if char == ')' {
			groupingLevel -= 1
			resetToken = true
		} else if char == '{' {
			resetToken = true
			end = true
			data = true
		} else if char == ' ' && groupingLevel == 0 {
			resetToken = true
		} else {
			token += string(char)
		}

		if resetToken {
			tokenSlice = append(tokenSlice, token)
			token = ""

			resetToken = false
		}

		if end {
			tokenSlice = append(tokenSlice, input[i:])
			data = true
			break
		}
	}

	return removeEmptyStrings(tokenSlice), data

}

func Tokenize(input string) ([]string, bool) {
	tokens := []string{}
	level := 0
	var tok string

	for i, char := range input {
		head := ""
		if i+1 != len(input) {
			head = string(input[i+1])
		}
		if string(char) == " " && head != "(" && head != "{" && level == 0 {
			tokens = append(tokens, tok)
			tok = ""
		} else {
			switch string(head) {
			case "(":
				level += 1
			case ")":
				level -= 1
			case "{":
				if tok != "" {
					tokens = append(tokens, tok)
				}
				tokens = append(tokens, input[i:])
				return tokens, true
			case "":
				if tok != "" {
					tokens = append(tokens, tok)
				}
				return tokens, false
			}
			tok += string(char)
		}
	}

	if tok != "" {
		tokens = append(tokens, tok)
	}

	return tokens, false
}
