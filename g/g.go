package g

import (
	"reflect"
	"unsafe"
)

//go:nosplit
func getg() unsafe.Pointer

//go:nosplit
//go:linkname getgp runtime.getgp
func getgp() unsafe.Pointer {
	return getg()
}

//go:nosplit
//go:linkname getgt runtime.getgt
func getgt() reflect.Type {
	return find("runtime.g")
}
