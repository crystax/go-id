// SPDX-License-Identifier: MIT
//go:build go1.23
package g

import (
	"reflect"
	"unsafe"
)

type iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

// According to the comments at https://github.com/golang/go/blob/go1.23.11/src/runtime/runtime1.go#L619-L643,
// it's safe to use reflect.typelinks() since its signature and behavior is not going to be changed despite
// the fact it's an internal function.
//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)

// According to the comments at https://github.com/golang/go/blob/go1.23.11/src/runtime/runtime1.go#L660-L676,
// it's safe to use reflect.resolveTypeOff() since its signature and behavior is not going to be changed despite
// the fact it's an internal function.
//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

func newType(section unsafe.Pointer, off int32) reflect.Type {
	typ := reflect.TypeOf(0)
	(*iface)(unsafe.Pointer(&typ)).data = resolveTypeOff(section, off)
	return typ
}

func find(typename string) reflect.Type {
	if typename == "" {
		return nil
	}

	ptr := typename
	if typename[0] != '*' {
		ptr = "*" + ptr
	}

	sections, offset := typelinks()
	for idx, offs := range offset {
		section := sections[idx]

		findType := func() reflect.Type {
			for _, off := range offs {
				typ := newType(section, off)
				if typ.String() != ptr {
					continue
				}

				return typ
			}

			return nil
		}

		typ := findType()
		if typ == nil {
			continue
		}

		if typ.Kind() != reflect.Ptr {
			continue
		}

		if typ.String() == typename {
			return typ
		}

		elem := typ.Elem()
		if elem.String() == typename {
			return elem
		}
	}

	return nil
}
