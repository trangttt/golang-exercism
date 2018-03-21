package circular

import "errors"

type Buffer struct {
	data []byte
	count int
	index int
}
func NewBuffer(size int) *Buffer {
	data := make([]byte, size)
	return &Buffer{data: data, count: 0, index: 0}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("Empty buffer")
	}
	ret := b.data[b.index]
	b.count--
	b.index = (b.index + 1) % len(b.data)
	return ret, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.count == len(b.data) {
		return errors.New("Full buffer")
	}
	next := (b.index + b.count) % len(b.data)
	b.data[next] = c
	b.count++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.count < len(b.data) {
		b.WriteByte(c)
		return
	}

	b.data[b.index] = c
	b.index = (b.index + 1) % len(b.data)

}
func (b *Buffer) Reset() {
	b.count = 0
	b.index = 0
}


