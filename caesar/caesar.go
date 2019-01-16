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
	return []string{}
}

// DemovingShift ...
func DemovingShift(arr []string, shift int) string {
	return ""
}
