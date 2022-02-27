package typeutil

import (
	"go/types"

	"golang.org/x/exp/typeparams"
)

func All(terms []*typeparams.Term, fn func(*typeparams.Term) bool) bool {
	if len(terms) == 0 {
		return fn(nil)
	}
	for _, term := range terms {
		if !fn(term) {
			return false
		}
	}
	return true
}

func Any(terms []*typeparams.Term, fn func(*typeparams.Term) bool) bool {
	if len(terms) == 0 {
		return fn(nil)
	}
	for _, term := range terms {
		if fn(term) {
			return true
		}
	}
	return false
}

func AllAndAny(terms []*typeparams.Term, fn func(*typeparams.Term) bool) bool {
	return All(terms, func(term *typeparams.Term) bool {
		if term == nil {
			return false
		}
		return fn(term)
	})
}

// XXX handle directional channels correctly
func CoreType(t types.Type) types.Type {
	if t, ok := t.(*typeparams.TypeParam); ok {
		terms, err := typeparams.NormalTerms(t)
		if err != nil || len(terms) == 0 {
			return nil
		}
		typ := terms[0].Type().Underlying()
		for _, term := range terms[1:] {
			if !types.Identical(term.Type().Underlying(), typ) {
				return nil
			}
		}
		return typ
	}
	return t
}
