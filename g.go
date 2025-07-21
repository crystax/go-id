package main

import (
	"reflect"
	"unsafe"
	_ "test/g"
)

type g struct {}

//go:norace
func (g *g) goid() uint64 {
	return get[uint64](g, offsetGoid)
}

func getg() *g {
	p := getgp()
	if p == nil {
		panic("can't obtain the address of 'g'")
	}

	return p
}

//go:linkname getgp runtime.getgp
func getgp() *g

//go:linkname getgt runtime.getgt
func getgt() reflect.Type

// A helper to obtain a field of type F from the object 'obj'.
// The field is located at the 'offset' within the object 'obj'.
func get[F any, T any](obj *T, offset uintptr) F {
	return *(*F)(unsafe.Pointer(uintptr(unsafe.Pointer(obj)) + offset))
}
