package piano

const KeyboardLen int = 88
const SectionLen int = 12
const StartKeyboardLen int = 3
const White string = "white"
const Black string = "black"

// BlackOrWhiteKey will return if a given key is black or white.
func BlackOrWhiteKey(keyPressCount int) string {

	// We know that we have an 88 key piano. So:
	// 1) Figure out how many piano lengths there are.
	// 2) Work out the remainder on an 88 key piano given it is comprised of repeating "sections" excpet for the start and the end of the keyboard.
	keyPosition := keyPressCount % KeyboardLen

	// Deal with start and end of the keyboard
	// Remember we did a modulo with KeyboardLen!
	if keyPosition <= 3 { // Start of the keyboard.
		if keyPosition != 0 && keyPosition%2 == 0 {
			return Black
		}
		return White
	}

	// Once we take away the first few keys, now we are dealing with a set of
	// "sections". Each section has 12 keys. So mod by 12 to get the key
	// position in the section.
	keyPositionSection := (keyPosition - StartKeyboardLen) % SectionLen

	// We know we have black keys at positions 2, 4, 7, 9 and 11.
	switch keyPositionSection {
	case 2, 4, 7, 9, 11:
		return Black
	default:
		return White
	}
}
