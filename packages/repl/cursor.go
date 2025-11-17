package repl

import (
	"fmt"
	"strings"

	ansix364 "github.com/abaxoth0/Janus/packages/ansix3.64"
	"github.com/abaxoth0/Janus/packages/ascii"
)

type cursor struct {

}

// Writes given char at current cursor position.
func (c *cursor) WriteChar(char rune) *cursor {
	fmt.Printf("%c", char)
	return c
}

// Writes given string at current cursor position.
func (c *cursor) Write(s string) *cursor {
	fmt.Print(s)
	return c
}

// Moves cursor to the beginning of line.
func (c *cursor) Rewind() *cursor {
	fmt.Print(ascii.CarriageReturn)
	return c
}

// Moves cursor 1 character back.
func (c *cursor) Back() *cursor {
	fmt.Print(ansix364.Left)
	return c
}

// Moves cursor N character(-s) back.
// If N equals zero does nothing.
func (c *cursor) BackN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.LeftN(int(n)))
	return c
}

// Moves cursor 1 character forward.
func (c *cursor) Forward() *cursor {
	fmt.Print(ansix364.Right)
	return c
}

// Moves cursor N character(-s) forward.
// If N equals zero does nothing.
func (c *cursor) ForwardN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.RightN(int(n)))
	return c
}

// Moves cursor 1 line up.
func (c *cursor) Up() *cursor {
	fmt.Print(ansix364.Up)
	return c
}

// Moves cursor N line(-s) up.
// If N equals zero does nothing.
func (c *cursor) UpN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.UpN(int(n)))
	return c
}

// Moves cursor 1 line down.
func (c *cursor) Down() *cursor {
	fmt.Print(ansix364.Down)
	return c
}

// Moves cursor N line(-s) down.
// If N equals zero does nothing.
func (c *cursor) DownN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.DownN(int(n)))
	return c
}

// Moves cursor to the beginning of new line.
func (c *cursor) NewLine() *cursor {
	fmt.Print(ascii.LineFeed)
	return c
}

func (c *cursor) repeat(char rune, n uint8) *cursor {
	if n == 0 {
		return c
	}

	str := strings.Builder{}
	str.Grow(int(n))

	for range n {
		if _, err := str.WriteRune(char); err != nil {
			// TODO FIX: Replace this after adding a proper log system
			panic(err)
		}
	}

	fmt.Print(str.String())

	return c
}

// Moves cursor to the beginning of new line, skipping N - 1 lines.
// If N equals zero does nothing.
func (c *cursor) NewLineN(n uint8) *cursor {
	return c.repeat(ascii.LineFeed, n)
}

// Saves current cursor position.
func (c *cursor) SavePosition() *cursor {
	fmt.Print(ansix364.SavePosition)
	return c
}

// Restores saved cursor position.
func (c *cursor) RestorePosition() *cursor {
	fmt.Print(ansix364.RestorePosition)
	return c
}

// Removes all characters in inline and moves cursor to it's beginning.
func (c *cursor) FlushLine() *cursor {
	fmt.Print(ansix364.EraseAll+string(ascii.CarriageReturn))
	return c
}
