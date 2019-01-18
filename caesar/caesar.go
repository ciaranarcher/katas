package caesar

const StartLowerAlpha int = 97
const EndLowerAlpha int = 122
const StartUpperAlpha int = 65
const EndUpperAlpha int = 90
const StartNumbers int = 48 // 0
const EndNumbers int = 57   // 9

// Shift ...
func Shift(s string, shiftFactor int) string {
	chars := []rune(s)
	var shiftedChars []rune
	var shiftedChar rune

	for _, char := range chars {
		shiftedChar = shiftBasedOnRange(char, StartLowerAlpha, EndLowerAlpha, shiftFactor)
		if shiftedChar != char {
			shiftedChars = append(shiftedChars, shiftedChar)
			continue
		}

		shiftedChar = shiftBasedOnRange(char, StartUpperAlpha, EndUpperAlpha, shiftFactor)
		if shiftedChar != char {
			shiftedChars = append(shiftedChars, shiftedChar)
			continue
		}

		shiftedChars = append(shiftedChars, char)
	}

	return string(shiftedChars)

}

func shiftBasedOnRange(char rune, startRange, endRange, shiftFactor int) rune {
	shiftedChar := int(char)
	if shiftedChar >= startRange && shiftedChar <= endRange {
		shiftedChar = shiftedChar + shiftFactor
		if shiftedChar > endRange {
			shiftedChar = (startRange - 1) + (shiftedChar % endRange)
		}
		return rune(shiftedChar)
	}
	return char
}

// MovingShift ...
func MovingShift(s string, shift int) []string {
	chars := []rune(s)

	var shiftedStr string
	var rotatingShift int

	for i, char := range chars {
		// Make sure not to shift by more than the number of letters in the alphabet.
		rotatingShift = (shift + i) % 26
		shiftedStr = shiftedStr + Shift(string(char), rotatingShift)
	}

	// Now we have a shifted array of runes / chars. We will split this into 5 groups.

	var splitStr []string

	splitLen := len(shiftedStr)/5 + 1
	chars = []rune(shiftedStr)

	// All but the last group are fixed size.
	for i := range chars {
		if i > 0 && i%splitLen == 0 && i < len(chars)-1 {
			splitStr = append(splitStr, string(chars[i-splitLen:i]))
		}

	}

	// The last group is the remaining letters.
	charsLen := len(chars)
	from := charsLen - charsLen%splitLen
	splitStr = append(splitStr, string(chars[from:charsLen]))

	return splitStr
}

// DemovingShift ...
func DemovingShift(arr []string, shift int) string {
	return ""
}
