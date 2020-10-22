package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func readLen(buf *bytes.Buffer) (res int, err error) {
	ln := binary.LittleEndian.Uint16(buf.Next(2))
	if uint16(buf.Len()) < ln {
		err = errors.New("Not enough bytes to read lenght value")
		return
	}

	res = int(ln)
	return
}

func readUint16(buf *bytes.Buffer) (res uint16, err error) {
	if buf.Len() < 2 {
		err = errors.New("Not enough bytes to read uint16 value")
		return
	}

	res = binary.LittleEndian.Uint16(buf.Next(2))

	return
}

func readString(buf *bytes.Buffer) (res string, err error) {
	ln, err := readLen(buf)
	if err != nil {
		return
	}

	res = string(buf.Next(int(ln)))
	return
}

func writeGeneric(buf *bytes.Buffer, v interface{}) (err error) {
	err = binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return
	}

	return
}

func writeString(buf *bytes.Buffer, str string) (err error) {
	err = writeGeneric(buf, uint16(len(str)))
	if err != nil {
		return
	}

	err = binary.Write(buf, binary.LittleEndian, str)
	if err != nil {
		return
	}

	return
}
