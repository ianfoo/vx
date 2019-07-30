package vx_test

import (
	"testing"

	"github.com/ianfoo/vx"
	"github.com/matryer/is"
)

func TestProceed(t *testing.T) {
	tt := []struct {
		in                     string
		wantProceed, wantValid bool
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
			is := is.New(t)
			gotProceed, gotValid := vx.Proceed(tc.in)
			is.Equal(tc.wantProceed, gotProceed) // proceed return value
			is.Equal(tc.wantValid, gotValid)     // valid input return value
		})
	}
}
