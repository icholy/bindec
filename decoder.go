package bindec

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Decoder struct {
	R     *bufio.Reader
	Order binary.ByteOrder
}

func NewDecoder(r io.Reader, order binary.ByteOrder) *Decoder {
	return &Decoder{R: bufio.NewReader(r), Order: order}
}

func (d Decoder) Bytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(d.R, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (d Decoder) Uint16() (uint16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	return d.Order.Uint16(buf[:]), nil
}

func (d Decoder) Int16() (int16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	return int16(d.Order.Uint16(buf[:])), nil
}

func (d Decoder) Uint32() (uint32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	return d.Order.Uint32(buf[:]), nil
}

func (d Decoder) Int32() (int32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	return int32(d.Order.Uint32(buf[:])), nil
}

func (d Decoder) Uint64() (uint64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	return d.Order.Uint64(buf[:]), nil
}

func (d Decoder) Int64() (int64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	return int64(d.Order.Uint64(buf[:])), nil
}
