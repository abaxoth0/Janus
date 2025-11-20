package ascii

type Char = rune

// Return true if given Char is one of ASCII control characters,
// otherwise return false.
func IsControlChar(c Char) bool {
	return c <= 31
}

// Returns true if given Char is a letter of the Latin alphabet,
// otherwise returns false.
func IsAlpha(c Char) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}
