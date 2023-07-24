package protocol

import (
	"bytes"
	"encoding/binary"
)

type Buffer struct {
	buffer *bytes.Buffer
}

func BufferFrom(data []byte) Buffer {
	return Buffer{
		buffer: bytes.NewBuffer(data),
	}
}

func EmptyBuffer() Buffer {
	return Buffer{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

/* Read instructions */

func (p Buffer) ReadBoolean() bool {
	var b bool
	binary.Read(p.buffer, binary.BigEndian, &b)
	return b
}

func (p Buffer) ReadShort() int {
	var i int16
	binary.Read(p.buffer, binary.BigEndian, &i)
	return int(i)
}

func (p Buffer) ReadInt() int {
	var i int32
	binary.Read(p.buffer, binary.BigEndian, &i)
	return int(i)
}

func (p Buffer) ReadString() string {
	len := binary.BigEndian.Uint16(p.buffer.Next(2))
	return string(p.buffer.Next(int(len)))
}

/* Write instructions */

func (p Buffer) WriteShort(i int) {
	binary.Write(p.buffer, binary.BigEndian, int16(i))
}

func (p Buffer) WriteInt(i int) {
	binary.Write(p.buffer, binary.BigEndian, int32(i))
}

func (p Buffer) WriteString(s string) {
	b := []byte(s)
	p.WriteShort(len(b))
	p.buffer.Write(b)
}

func (p Buffer) WriteBoolean(b bool) {
	binary.Write(p.buffer, binary.BigEndian, b)
}

func (p Buffer) Flush() []byte {
	msg := make([]byte, 4)
	binary.BigEndian.PutUint32(msg, uint32(p.buffer.Len()))
	return append(msg, p.buffer.Bytes()...)
}
