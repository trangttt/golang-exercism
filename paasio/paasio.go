package paasio

import (
	"io"
	"sync/atomic"
)

type MyReadCounter struct {
	ByteCount *int64
	CallCount *int32
	Reader    io.Reader
	Channel   chan string
}

type MyWriteCounter struct {
	ByteCount *int64
	CallCount *int32
	Writer    io.Writer
	Channel   chan string
}

func (rc MyReadCounter) Read(p []byte) (n int, err error) {
	n, err = rc.Reader.Read(p)
	if err == nil {
		<-rc.Channel
		atomic.AddInt64(rc.ByteCount, int64(n))
		atomic.AddInt32(rc.CallCount, int32(1))
		rc.Channel <- "OK"
	}
	return n, err
}

func (rc MyReadCounter) ReadCount() (n int64, nops int) {
	<-rc.Channel
	n, nops = atomic.LoadInt64(rc.ByteCount), int(atomic.LoadInt32(rc.CallCount))
	rc.Channel <- "OK"
	return n, nops
}

func (wc MyWriteCounter) Write(p []byte) (n int, err error) {
	n, err = wc.Writer.Write(p)
	if err == nil {
		<-wc.Channel
		atomic.AddInt64(wc.ByteCount, int64(n))
		atomic.AddInt32(wc.CallCount, int32(1))
		wc.Channel <- "OK"
	}
	return n, err
}

func (wc MyWriteCounter) WriteCount() (n int64, nops int) {
	<-wc.Channel
	n, nops = atomic.LoadInt64(wc.ByteCount), int(atomic.LoadInt32(wc.CallCount))
	wc.Channel <- "OK"
	return n, nops
}

func NewReadCounter(reader io.Reader) ReadCounter {
	var bc int64
	var cc int32
	channel := make(chan string, 1)
	channel <- "OK"
	mrc := MyReadCounter{
		ByteCount: &bc,
		CallCount: &cc,
		Reader:    reader,
		Channel:   channel,
	}
	return mrc
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	var bc int64
	var cc int32
	channel := make(chan string, 1)
	channel <- "OK"
	mwc := MyWriteCounter{
		ByteCount: &bc,
		CallCount: &cc,
		Writer:    writer,
		Channel:   channel,
	}
	return mwc
}

type MyReadWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewReadWriteCounter(readWriter io.ReadWriter) ReadWriteCounter {
	reader, _ := readWriter.(io.Reader)
	writer, _ := readWriter.(io.Writer)
	return MyReadWriteCounter{ReadCounter: NewReadCounter(reader),
		WriteCounter: NewWriteCounter(writer)}
}
