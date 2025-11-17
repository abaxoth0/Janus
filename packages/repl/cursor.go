package repl

import (
	"fmt"
	"strconv"
	"strings"
)

type cursor struct {

}

func (c *cursor) WriteChar(char rune) *cursor {
	fmt.Printf("%c", char)
	return c
}

func (c *cursor) Write(s string) *cursor {
	print(s)
	return c
}

func (c *cursor) Rewind() *cursor {
	print("\r")
	return c
}

func (c *cursor) Back() *cursor {
	print("\033[1D")
	return c
}

func (c *cursor) BackN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	print("\033["+strconv.Itoa(int(n))+"D")
	return c
}

func (c *cursor) Forward() *cursor {
	print("\033[1C")
	return c
}

func (c *cursor) ForwardN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	print("\033["+strconv.Itoa(int(n))+"C")
	return c
}

func (c *cursor) Up() *cursor {
	print("\033[1A")
	return c
}

func (c *cursor) UpN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	print("\033["+strconv.Itoa(int(n))+"A")
	return c
}

func (c *cursor) Down() *cursor {
	print("\033[1B")
	return c
}

func (c *cursor) DownN(n uint8) *cursor {
	if n == 0 {
		return c
	}
	print("\033["+strconv.Itoa(int(n))+"B")
	return c
}

func (c *cursor) NewLine() *cursor {
	print("\n")
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

	print(str.String())

	return c
}

func (c *cursor) NewLineN(n uint8) *cursor {
	return c.repeat('\n', n)
}

func (c *cursor) SavePosition() *cursor {
	print("\033[s")
	return c
}

func (c *cursor) RestorePosition() *cursor {
	print("\033[u")
	return c
}

func (c *cursor) FlushLine() *cursor {
	print("\033[2K")
	return c
}
