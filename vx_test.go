package vx_test

import (
	"testing"

	"github.com/ianfoo/vx"
)

func TestProceed(t *testing.T) {
	tt := []struct {
		in             string
		proceed, valid bool
	}{
		{"k", true, true},
		{"kk", true, true},
		{" k ", true, true},
		{"kkk", false, false},
		{"yeh", true, true},
		{" yeh ", true, true},
		{"yeah", false, false},
		{"w/e", false, true},
		{"\tw/e ", false, true},
		{"naw", false, true},
		{"naw dawg", false, true},
		{"naw\tdawg", false, true},
		{"naw   dawg", false, true},
		{"ur face", false, true},
		{"ur  face", false, true},
		{"gtfo", false, true},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			gotProceed, gotValid := vx.Proceed(tc.in)
			if want, got := tc.proceed, gotProceed; want != got {
				t.Errorf("%q: wanted proceed value = %t, got %t", tc.in, want, got)
			}
			if want, got := tc.valid, gotValid; want != got {
				t.Errorf("%q: wanted valid value = %t, got %t", tc.in, want, got)
			}
		})
	}
}
