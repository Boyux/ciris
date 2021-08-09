// +build linux

package ciris

import "syscall"

func Gettid() uint64 {
	return uint64(syscall.Gettid())
}
