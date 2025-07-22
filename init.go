// SPDX-License-Identifier: MIT
package main

var offsetGoid uintptr

func init() {
	gt := getgt()
	if gt == nil {
		panic("can't obtain the 'g' type")
	}

	field, found := gt.FieldByName("goid")
	if !found {
		panic("can't find the 'goid' field in the 'g' struct")
	}

	offsetGoid = field.Offset
}
