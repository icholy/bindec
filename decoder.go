package bindec

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Decoder struct {
	R     *bufio.Reader
	order binary.ByteOrder
}

func NewDecoder(r io.Reader, order binary.ByteOrder) *Decoder {
	return &Decoder{R: bufio.NewReader(r), order: order}
}

func (d Decoder) Bytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(d.R, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (d Decoder) Uint16(order binary.ByteOrder) (uint16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	if order == nil {
		order = d.order
	}
	return order.Uint16(buf[:]), nil
}

func (d Decoder) Int16(order binary.ByteOrder) (int16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	if order == nil {
		order = d.order
	}
	return int16(order.Uint16(buf[:])), nil
}

func (d Decoder) Uint32(order binary.ByteOrder) (uint32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	if order == nil {
		order = d.order
	}
	return order.Uint32(buf[:]), nil
}

func (d Decoder) Int32(order binary.ByteOrder) (int32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	if order == nil {
		order = d.order
	}
	return int32(order.Uint32(buf[:])), nil
}

func (d Decoder) Uint64(order binary.ByteOrder) (uint64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	if order == nil {
		order = d.order
	}
	return order.Uint64(buf[:]), nil
}

func (d Decoder) Int64(order binary.ByteOrder) (int64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(d.R, buf[:]); err != nil {
		return 0, err
	}
	if order == nil {
		order = d.order
	}
	return int64(order.Uint64(buf[:])), nil
}
