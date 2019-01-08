package kata

func ValidBraces(str string) bool {

	// We should only allow *pairs* of braces
	// so any uneven string length is a fail.
	if len(str)%2 != 0 {
		return false
	}

	if matchingViaStack(str) {
		return true
	}

	return false
}

// First attempt (doesn't quite work).
// matchingSplitCompare will start at the center of the string and for each interation,
// walk towards the end comparing each character. Work for all test cases except "(){}[]"
// as it is not symmetrical. Assumes an even number of characters.
func matchingSymmetrical(str string) bool {

	length := len(str)
	chars := []rune(str)
	numIterations := length / 2

	matchers := map[string]string{
		"[": "]",
		"]": "[",
		"{": "}",
		"}": "{",
		"(": ")",
		")": "(",
	}

	// Start at the middle of the array and walk towards the ends looking for matching braces.
	for i, j := numIterations-1, numIterations; i >= 0; i, j = i-1, j+1 {
		if matchers[string(chars[i])] != string(chars[j]) {
			return false
		}
	}

	return true
}

// Second attempt (succcessful).
// matchingViaStack will note every opening brace found and push it onto a stack.
// When a closing brace is found, the last opening brace is popped off the stack
// and if it doesn't match the closing brace we know we have a mismatch.
func matchingViaStack(str string) bool {
	chars := []rune(str)
	bracePairs := map[string]string{
		"[": "]",
		"{": "}",
		"(": ")",
	}

	bracesFound := []string{}
	startBrace := ""

	for _, char := range chars {
		if _, ok := bracePairs[string(char)]; ok {
			bracesFound = append(bracesFound, string(char))
		} else {
			// Check to ensure that we have any matching braces before we try to read them.
			// If we don't we are in a strange state and should exit.
			if len(bracesFound) == 0 {
				return false
			}
			startBrace, bracesFound = bracesFound[len(bracesFound)-1], bracesFound[:len(bracesFound)-1]

			if bracePairs[startBrace] != string(char) {
				// Mismatch, bail out!
				return false
			}
		}
	}

	// At the end there should be no braces left.
	if len(bracesFound) > 0 {
		return false
	}

	return true
}
