// +build cgo

package ciris

// #cgo CFLAGS: -I${SRCDIR}/include
// #cgo LDFLAGS: -L${SRCDIR}/lib
// #cgo darwin,amd64 LDFLAGS: -lzebra_darwin_amd64
// #cgo darwin,arm64 LDFLAGS: -lzebra_darwin_aarch64
// #cgo linux,amd64 LDFLAGS: -lzebra_linux_amd64
// #include "zebra.h"
import "C"
import "unsafe"

func Gen(frame string) string {
	frame_c := C.CString(frame)
	defer C.free(unsafe.Pointer(frame_c))
	token_c := C.zebra_gen_token(frame_c)
	defer C.zebra_free_token(token_c)
	bytes := *(*[40]byte)(unsafe.Pointer(&token_c.buf))
	return string(bytes[0:token_c.len])
}

func Verify(frame string, token string) bool {
	frame_c := C.CString(frame)
	defer C.free(unsafe.Pointer(frame_c))
	token_c := C.CString(token)
	defer C.free(unsafe.Pointer(token_c))
	return C.zebra_verify_token(frame_c, token_c) == 1
}
