package typeutil

import (
	"errors"
	"go/types"

	"golang.org/x/exp/typeparams"
)

type TypeSet struct {
	Terms []*typeparams.Term
	empty bool
}

func NewTypeSet(typ types.Type) TypeSet {
	terms, err := typeparams.NormalTerms(typ)
	if err != nil {
		if errors.Is(err, typeparams.ErrEmptyTypeSet) {
			return TypeSet{nil, true}
		} else {
			// We couldn't determine the type set. Assume it's all types.
			return TypeSet{nil, false}
		}
	}
	return TypeSet{terms, false}
}

// CoreType returns the type set's core type, or nil if it has none.
func (ts TypeSet) CoreType() types.Type {
	// XXX handle directional channels correctly

	if len(ts.Terms) == 0 {
		// Either the type set is empty, or it isn't constrained. Either way it doesn't have a core type.
		return nil
	}
	typ := ts.Terms[0].Type().Underlying()
	for _, term := range ts.Terms[1:] {
		if !types.Identical(term.Type().Underlying(), typ) {
			return nil
		}
	}
	return typ
}

// CoreType is a wrapper for NewTypeSet(typ).CoreType()
func CoreType(typ types.Type) types.Type {
	return NewTypeSet(typ).CoreType()
}

// All calls fn for each term in the type set and reports whether all invocations returned true.
// If the type set is empty or unconstrained, All immediately returns false.
func (ts TypeSet) All(fn func(*typeparams.Term) bool) bool {
	if len(ts.Terms) == 0 {
		return false
	}
	for _, term := range ts.Terms {
		if !fn(term) {
			return false
		}
	}
	return true
}

// Any calls fn for each term in the type set and reports whether any invocation returned true.
// It stops after the first call that returned true.
func (ts TypeSet) Any(fn func(*typeparams.Term) bool) bool {
	for _, term := range ts.Terms {
		if fn(term) {
			return true
		}
	}
	return false
}

// All is a wrapper for NewTypeSet(typ).All(fn).
func All(typ types.Type, fn func(*typeparams.Term) bool) bool {
	return NewTypeSet(typ).All(fn)
}

// Any is a wrapper for NewTypeSet(typ).Any(fn).
func Any(typ types.Type, fn func(*typeparams.Term) bool) bool {
	return NewTypeSet(typ).Any(fn)
}
