package paasio

import (
	"io"
	"sync"
)

type counter struct {
	callCounter int
	length      int64
	sync.Mutex
}

type writerCounter struct {
	buf io.Writer
	counter
}

type readerCounter struct {
	buf io.Reader
	counter
}

type rwCounter struct {
	ReadCounter
	WriteCounter
}

func (c *counter) addBytes(l int) {
	c.Lock()
	defer c.Unlock()

	c.callCounter++
	c.length += int64(l)
}

func (c *counter) count() (int64, int) {
	c.Lock()
	defer c.Unlock()

	return c.length, c.callCounter
}

func (wc *writerCounter) Write(s []byte) (int, error) {
	l, err := wc.buf.Write(s)

	if err != nil {
		return l, err
	}

	wc.addBytes(l)
	return l, err
}

func (wc *writerCounter) WriteCount() (int64, int) {
	return wc.count()
}

func (rc *readerCounter) Read(s []byte) (int, error) {
	l, err := rc.buf.Read(s)

	rc.addBytes(l)

	return l, err
}

func (rc *readerCounter) ReadCount() (int64, int) {
	return rc.count()
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writerCounter{buf: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readerCounter{buf: reader}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{
		NewReadCounter(rw),
		NewWriteCounter(rw),
	}
}
