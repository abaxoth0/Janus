package repl

import (
	"fmt"
	"strings"

	ansix364 "github.com/abaxoth0/Janus/packages/ansix3.64"
	"github.com/abaxoth0/Janus/packages/ascii"
)

type cursor struct {

}

func (c *cursor) WriteChar(char rune) *cursor {
	fmt.Printf("%c", char)
	return c
}

func (c *cursor) Write(s string) *cursor {
	fmt.Print(s)
	return c
}

func (c *cursor) Rewind() *cursor {
	fmt.Print(ascii.CarriageReturn)
	return c
}

func (c *cursor) Back() *cursor {
	fmt.Print(ansix364.Left)
	return c
}

func (c *cursor) BackN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.LeftN(int(n)))
	return c
}

func (c *cursor) Forward() *cursor {
	fmt.Print(ansix364.Right)
	return c
}

func (c *cursor) ForwardN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.RightN(int(n)))
	return c
}

func (c *cursor) Up() *cursor {
	fmt.Print(ansix364.Up)
	return c
}

func (c *cursor) UpN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.UpN(int(n)))
	return c
}

func (c *cursor) Down() *cursor {
	fmt.Print(ansix364.Down)
	return c
}

func (c *cursor) DownN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	fmt.Print(ansix364.DownN(int(n)))
	return c
}

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

func (c *cursor) NewLineN(n uint8) *cursor {
	return c.repeat(ascii.LineFeed, n)
}

func (c *cursor) SavePosition() *cursor {
	fmt.Print(ansix364.SavePosition)
	return c
}

func (c *cursor) RestorePosition() *cursor {
	fmt.Print(ansix364.RestorePosition)
	return c
}

func (c *cursor) FlushLine() *cursor {
	fmt.Print(ansix364.EraseAll+string(ascii.CarriageReturn))
	return c
}
