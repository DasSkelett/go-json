package json

import (
	"unsafe"
)

type ptrDecoder struct {
	dec decoder
	typ *rtype
}

func newPtrDecoder(dec decoder, typ *rtype) *ptrDecoder {
	return &ptrDecoder{dec: dec, typ: typ}
}

//go:linkname unsafe_New reflect.unsafe_New
func unsafe_New(*rtype) uintptr

func (d *ptrDecoder) decode(buf []byte, cursor int, p uintptr) (int, error) {
	newptr := unsafe_New(d.typ)
	c, err := d.dec.decode(buf, cursor, newptr)
	if err != nil {
		return 0, err
	}
	cursor = c
	*(*uintptr)(unsafe.Pointer(p)) = newptr
	return cursor, nil
}