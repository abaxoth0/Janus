package ascii

type Char = rune

func IsControlChar(c Char) bool {
	return c <= 31
}
