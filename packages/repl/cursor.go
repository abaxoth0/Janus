package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ansix364 "github.com/abaxoth0/Janus/packages/ansix3.64"
	"github.com/abaxoth0/Janus/packages/ascii"
)

type cursor struct {
	x      int
	minX   int
	savedX int
}

// IMPORTANT:
// After creating a cursor use only it's own methods for the output.
// Otherwise it will lead to invalid cursor behaviour.
func newCursor(minX int) *cursor {
	cur := new(cursor)

	cur.x, _ = cur.FetchPosition()

	if cur.x < minX {
		cur.x = minX
	}

	cur.minX = minX
	cur.savedX = -1

	return cur
}

// Returns X position of a cursor.
func (c *cursor) GetX() int {
	return c.x
}

// Uses ANSI X3.64 escape sequence to get current X and Y
// cursor position in terminal and return it.
func (c *cursor) FetchPosition() (x int, y int) {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print(ansix364.OutputPosition)

	str, err := inputReader.ReadString('R')
	if err != nil {
		panic(err)
	}

	_, err = fmt.Sscanf(str, "\033[%d;%dR", &y, &x)
	if err != nil {
		panic(err)
	}

	return x, y
}

// Writes given char at current cursor position.
func (c *cursor) WriteChar(char rune) *cursor {
	fmt.Printf("%c", char)
	if char == ascii.CarriageReturn || char == ascii.LineFeed {
		c.x = c.minX
		return c
	}
	c.x++
	return c
}

// Writes given value at current cursor position.
// Wrapper for fmt.Print()
func (c *cursor) RawWrite(a ...any) *cursor {
	fmt.Print(a...)
	c.x, _ = c.FetchPosition()
	return c
}

// Writes given string at current cursor position.
func (c *cursor) Write(s string) *cursor {
	fmt.Print(s)
	end := s[len(s)-1]
	if ascii.Char(end) == ascii.CarriageReturn || ascii.Char(end) == ascii.LineFeed {
		c.x = c.minX
		return c
	}
	c.x += len(s)
	return c
}

// Writes given string at current cursor position and moves cursor to a new line.
func (c *cursor) Writeln(s string) *cursor {
	return c.Write(s).NewLine()
}

// Moves cursor to the beginning of line.
func (c *cursor) Rewind() *cursor {
	fmt.Print(ascii.CarriageReturn)
	c.x = c.minX
	return c
}

// Moves cursor 1 character back.
// Does nothing if cursor X position is equal to it's min possible value.
func (c *cursor) Back() *cursor {
	if c.x == c.minX {
		return c
	}
	fmt.Print(ansix364.Left)
	c.x--
	return c
}

// Moves cursor N character(-s) back.
// If N equals zero does nothing.
func (c *cursor) BackN(n int) *cursor {
	if n == 0 {
		return c
	}
	if c.x == c.minX {
		return c
	}
	if c.x-n <= c.minX {
		n = c.x - c.minX
	}
	fmt.Print(ansix364.LeftN(n))
	return c
}

// Moves cursor 1 character forward.
// Does nothing if cursor X position is equal to it's max possible value.
func (c *cursor) Forward() *cursor {
	fmt.Print(ansix364.Right)
	c.x++
	return c
}

// Moves cursor N character(-s) forward.
// If N equals zero does nothing.
func (c *cursor) ForwardN(n int) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.RightN(n))
	c.x += n
	return c
}

// Moves cursor 1 line up.
func (c *cursor) Up() *cursor {
	fmt.Print(ansix364.Up)
	return c
}

// Moves cursor N line(-s) up.
// If N equals zero does nothing.
func (c *cursor) UpN(n int) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.UpN(n))
	return c
}

// Moves cursor 1 line down.
func (c *cursor) Down() *cursor {
	fmt.Print(ansix364.Down)
	return c
}

// Moves cursor N line(-s) down.
// If N equals zero does nothing.
func (c *cursor) DownN(n int) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.DownN(n))
	return c
}

// Moves cursor to the beginning of new line.
func (c *cursor) NewLine() *cursor {
	fmt.Print(ascii.LineFeed)
	c.x = c.minX
	return c
}

func (c *cursor) repeat(char rune, n int, moveX bool) *cursor {
	if n == 0 {
		return c
	}

	str := strings.Builder{}
	str.Grow(n)

	for range n {
		if _, err := str.WriteRune(char); err != nil {
			// TODO FIX: Replace this after adding a proper log system
			panic(err)
		}
	}
	if moveX {
		c.x += n
	}
	fmt.Print(str.String())

	return c
}

// Moves cursor to the beginning of new line, skipping N - 1 lines.
// If N equals zero does nothing.
func (c *cursor) NewLineN(n int) *cursor {
	return c.repeat(ascii.LineFeed, n, false)
}

// Saves current cursor position.
func (c *cursor) SavePosition() *cursor {
	c.savedX = c.x
	fmt.Print(ansix364.SavePosition)
	return c
}

// Restores saved cursor position.
func (c *cursor) RestorePosition() *cursor {
	if c.savedX == -1 {
		return c
	}
	fmt.Print(ansix364.RestorePosition)
	c.x = c.savedX
	return c
}

// Replaces character under cursor with space (' ') and moves cursor back one time.
func (c *cursor) FlushChar() *cursor {
	return c.Back().Write(" ").Back()
}

// Moves cursor back N times, replacing all characters on the way with spaces (' ').
// If N equals zero does nothing.
func (c *cursor) FlushChars(n int) *cursor {
	if n == 0 {
		return c
	}
	return c.BackN(n).repeat(' ', n, true).BackN(n)
}

// Moves cursor to the beginning of line, removing all characters in it.
func (c *cursor) FlushLine() *cursor {
	c.x = c.minX
	fmt.Print(ansix364.EraseAll + string(ascii.CarriageReturn))
	return c
}
