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

func LZ4_compressFrame(in []byte) (outArr []byte, err error) {
	out := make([]byte, (len(in)<<1)+128)
	n := C.LZ4F_compressFrame(unsafe.Pointer(&in[0]), C.ulong(len(in)), unsafe.Pointer(out), C.ulong(len(out)), nil)
	if n <= 0 {
		return out, errors.New("Compression error")
	}
	return out, nil
}