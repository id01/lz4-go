package lz4

/*
#cgo LDFLAGS: -llz4
#cgo CFLAGS: -O3
#include "lz4frame.h"
*/
import "C"

import (
	"unsafe"
	"errors"
)

func LZ4_compressFrame(in []byte) (out []byte, err error) {
	out = make([]byte, (len(in)<<1)+128) // Overestimate
	n := C.LZ4F_compressFrame(unsafe.Pointer(&out[0]), C.ulong(len(out)), unsafe.Pointer(&in[0]), C.ulong(len(in)), nil)
	if n <= 0 || n >= 9223372036854775807 {
		return out, errors.New("Compression error")
	}
	return out[:n], nil
}