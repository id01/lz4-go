package lz4

/*
#cgo LDFLAGS: -llz4
#cgo CFLAGS: -O3
#include "lz4frame.h"
size_t decompressFrame(void* dst, size_t dst_len, const void* src, size_t src_len) {
	LZ4F_dctx* dctx;
	LZ4F_decompressOptions_t opts;
	opts.stableDst = 1; // Requires at least an additional 64KB of dst buffer
	opts.reserved[0] = opts.reserved[1] = opts.reserved[2] = 0;
	if (LZ4F_isError(LZ4F_createDecompressionContext(&dctx, LZ4F_VERSION))) return -1;
	size_t src_len_read = src_len, dst_len_read = dst_len;
	if (LZ4F_decompress(dctx, dst, &dst_len_read, src, &src_len_read, &opts) != 0
		|| src_len_read != src_len) return -1;
	if (LZ4F_isError(LZ4F_freeDecompressionContext(dctx))) return -1;
	return dst_len_read;
}
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

func LZ4_decompressFrame(in []byte, maxSize int64) (out []byte, err error) {
	out = make([]byte, maxSize+65536) // Need additional 64KB for optimization
	n := C.decompressFrame(unsafe.Pointer(&out[0]), C.ulong(len(out)), unsafe.Pointer(&in[0]), C.ulong(len(in)))
	if n <= 0 || n >= 9223372036854775807 {
		return out, errors.New("Decompression error")
	}
	return out[:n], nil
}