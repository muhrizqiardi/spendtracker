package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Compare(exp, got interface{}, opts ...cmp.Option) (string, bool) {
	diff := cmp.Diff(exp, got, opts...)
	equal := cmp.Equal(exp, got, opts...)

	return diff, equal
}

func CompareAndAssert(t *testing.T, exp, got interface{}, opts ...cmp.Option) {
	if !cmp.Equal(exp, got, opts...) {
		t.Errorf("mismatch (-exp, +got): %s", cmp.Diff(exp, got, opts...))
	}
}
