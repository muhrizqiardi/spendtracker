package testutil

import "github.com/google/go-cmp/cmp"

func Compare(exp, got interface{}, opts ...cmp.Option) (string, bool) {
	diff := cmp.Diff(exp, got, opts...)
	equal := cmp.Equal(exp, got, opts...)

	return diff, equal
}
