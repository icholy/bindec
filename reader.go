package binread

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Reader struct {
	R     *bufio.Reader
	order binary.ByteOrder
}

func NewReader(r io.Reader, order binary.ByteOrder) *Reader {
	return &Reader{R: bufio.NewReader(r), order: order}
}

func (r Reader) Read(data []byte) (n int, err error) {
	return r.R.Read(data)
}

func (r Reader) ReadData(data interface{}) error {
	return binary.Read(r.R, r.order, data)
}

func (r Reader) ReadBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(r.R, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (r Reader) ReadUint16() (uint16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(r.R, buf[:]); err != nil {
		return 0, err
	}
	return r.order.Uint16(buf[:]), nil
}

func (r Reader) ReadInt16() (int16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(r.R, buf[:]); err != nil {
		return 0, err
	}
	return int16(r.order.Uint16(buf[:])), nil
}

func (r Reader) ReadUint32() (uint32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(r.R, buf[:]); err != nil {
		return 0, err
	}
	return r.order.Uint32(buf[:]), nil
}

func (r Reader) ReadInt32() (int32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(r.R, buf[:]); err != nil {
		return 0, err
	}
	return int32(r.order.Uint32(buf[:])), nil
}

func (r Reader) ReadUint64() (uint64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(r.R, buf[:]); err != nil {
		return 0, err
	}
	return r.order.Uint64(buf[:]), nil
}

func (r Reader) ReadInt64() (int64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(r.R, buf[:]); err != nil {
		return 0, err
	}
	return int64(r.order.Uint64(buf[:])), nil
}
