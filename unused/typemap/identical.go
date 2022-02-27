package typemap

import (
	"go/types"

	"honnef.co/go/tools/go/types/typeutil"
)

// XXX Go 1.18 makes it really annoying to vendor a modified copy of types.Identical... we'll need to fix this eventually

func Identical(x, y types.Type) bool {
	xi, okx := x.(*typeutil.Iterator)
	yi, oky := y.(*typeutil.Iterator)
	if okx != oky {
		return false
	}
	if okx {
		return Identical(xi.Elem(), yi.Elem())
	} else {
		return types.Identical(x, y)
	}
}

// // Identical reports whether x and y are identical types.
// // Unlike types.Identical, receivers of Signature types are not ignored.
// // Unlike types.Identical, interfaces are compared via pointer equality (except for the empty interface, which gets deduplicated).
// // Unlike types.Identical, structs are compared via pointer equality.
// func Identical(x, y types.Type) (ret bool) {
// 	return identical0(x, y)
// }
