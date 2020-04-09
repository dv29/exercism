package circular

import (
	"errors"
)

type Buffer struct {
	buf  []interface{}
	head int
	tail int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		buf:  make([]interface{}, size),
		head: -1,
		tail: -1,
	}
}

func (b *Buffer) isFull() bool {
	if b.head == 0 && b.tail == len(b.buf)-1 || b.head == b.tail+1 {
		return true
	}

	return false
}

func (b *Buffer) isEmpty() bool {
	if b.head == -1 {
		return true
	}
	return false
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.isEmpty() {
		return 0, errors.New("empty buffer")
	}
	val := b.buf[b.head]
	if b.head == b.tail {
		b.head = -1
		b.tail = -1
	} else {
		b.head = (b.head + 1) % len(b.buf)
	}

	return val.(byte), nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.isFull() {
		return errors.New("buffer full")
	}

	if b.head == -1 {
		b.head = 0
	}

	b.tail = (b.tail + 1) % len(b.buf)
	b.buf[b.tail] = c

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.isFull() {
		b.tail = b.head
		b.buf[b.head] = c
		b.head = (b.head + 1) % len(b.buf)
	} else {
		b.WriteByte(c)
	}
}

func (b *Buffer) Reset() {
	b.head = -1
	b.tail = -1
}
