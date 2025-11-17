package ansix364

import "strconv"

type EscapeSequence = string

// NOTE:
// Technically SavePosition and RestorePosition sequences are
// not a part of ANSI X3.64, but rather a part of private standart used in DEC VT100.
// (which supports ANSI escape codes, but anyway)
// So they are not belong to any other public standart.

const (
	Up              EscapeSequence = "\033[A"  // Moves cursor up one line.
	Down            EscapeSequence = "\033[B"  // Moves cursor down one line.
	Right           EscapeSequence = "\033[C"  // Moves cursor right one column.
	Left            EscapeSequence = "\033[D"  // Moves cursor left one column.
	EraseToEnd      EscapeSequence = "\033[K"  // Erase from cursor to end of line.
	EraseFromStart  EscapeSequence = "\033[1K" // Erase from start of line to cursor.
	EraseAll        EscapeSequence = "\033[2K" // Erase entire line.
	SavePosition    EscapeSequence = "\033[7"  // Save current cursor position.
	RestorePosition EscapeSequence = "\033[8"  // Restore saved cursor position.
)

// Moves cursor up N line(-s).
// Sets N to 1 if N is less than 0.
func UpN(n int) EscapeSequence {
	if n < 0 {
		n = 1
	}
	return EscapeSequence("\033[" + strconv.Itoa(n) + "A")
}

// Moves cursor down N line(-s).
// Sets N to 1 if N is less than 0.
func DownN(n int) EscapeSequence {
	if n < 0 {
		n = 1
	}
	return EscapeSequence("\033[" + strconv.Itoa(n) + "B")
}

// Moves cursor right N column(-s).
// Sets N to 1 if N is less than 0.
func RightN(n int) EscapeSequence {
	if n < 0 {
		n = 1
	}
	return EscapeSequence("\033[" + strconv.Itoa(n) + "C")
}

// Moves cursor left N column(-s).
// Sets N to 1 if N is less than 0.
func LeftN(n int) EscapeSequence {
	if n < 0 {
		n = 1
	}
	return EscapeSequence("\033[" + strconv.Itoa(n) + "D")
}
