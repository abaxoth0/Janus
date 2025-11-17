package ansix364

import "strconv"

type EscapeSequence = string

// NOTE:
// Technically SavePosition and RestorePosition sequences are
// not a part of ANSI X3.64, but rather a part of private standart used in DEC VT100.
// (which supports ANSI escape codes, but anyway)
// So they are not belong to any other public standart.

const (
	Up 				EscapeSequence = "\033[A"
	Down 			EscapeSequence = "\033[B"
	Right 			EscapeSequence = "\033[C"
	Left 			EscapeSequence = "\033[D"
	EraseToEnd		EscapeSequence = "\033[K"
	EraseToStart	EscapeSequence = "\033[1K"
	EraseAll		EscapeSequence = "\033[2K"
	SavePosition	EscapeSequence = "\033[7"
	RestorePosition	EscapeSequence = "\033[8"
)

func UpN(n int) EscapeSequence {
	return EscapeSequence("\033["+strconv.Itoa(n)+"A")
}

func DownN(n int) EscapeSequence {
	return EscapeSequence("\033["+strconv.Itoa(n)+"B")
}

func RightN(n int) EscapeSequence {
	return EscapeSequence("\033["+strconv.Itoa(n)+"C")
}

func LeftN(n int) EscapeSequence {
	return EscapeSequence("\033["+strconv.Itoa(n)+"D")
}
