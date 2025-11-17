package ascii

type Char = rune

// Return true if given Char is one of ASCII control characters,
// otherwise return false.
func IsControlChar(c Char) bool {
	return c <= 31
}
