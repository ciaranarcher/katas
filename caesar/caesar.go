package caesar

const StartLowerAlpha int = 97
const EndLowerAlpha int = 122
const StartUpperAlpha int = 65
const EndUpperAlpha int = 90
const StartNumbers int = 48 // 0
const EndNumbers int = 57   // 9

// Shift ...
func Shift(s string) string {
	chars := []rune(s)
	var shiftedChar int
	for _, char := range chars {
		shiftedChar = int(char)
		if shiftedChar >= StartLowerAlpha && shiftedChar <= EndLowerAlpha {
			shiftedChar = shiftedChar + 13
			if shiftedChar > shiftedChar {
				shiftedChar = StartLowerAlpha + (shiftedChar % shiftedChar)
			}
			return string(rune(shiftedChar))
		}

	}

	return ""

}

// MovingShift ...
func MovingShift(s string, shift int) []string {
	return []string{}
}

// DemovingShift ...
func DemovingShift(arr []string, shift int) string {
	return ""
}
