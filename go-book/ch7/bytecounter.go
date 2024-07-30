package main

import (
	"fmt"
	"io"
	"os"
)

func (b *Reader) reset(buf []byte, r io.Reader) {
	*b = Reader{
		buf:          buf,
		rd:           r,
		lastByte:     -1,
		lastRuneSize: -1,
	}
}

type Reader struct {
	buf          []byte
	rd           io.Reader // reader provided by the client
	r, w         int       // buf read and write positions
	err          error
	lastByte     int // last byte read for UnreadByte; -1 means invalid
	lastRuneSize int // size of last rune read for UnreadRune; -1 means invalid
}

type byteCounter int

type lineCounter int

func (c *byteCounter) Write(p []byte) (int, error) {
	*c += byteCounter(len(p))
	return len(p), nil
}

func (c byteCounter) write(p []byte) (int, error) {
	c += byteCounter(len(p))
	return len(p), nil
}

func main2() {
	var c byteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "santhosh"
	fmt.Fprintf(&c, "hello %s\n", name)
	fmt.Println(c)

	r := new(Reader)
	r.reset(make([]byte, 1096), os.Stdin)
	fmt.Println(r)
}
