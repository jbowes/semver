// Copyright (c) 2021-2023 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashicorp_test

import (
	"testing"

	"github.com/hashicorp/go-version"
)

// Benchmarks are defined as slices so they run in the same order every time,
// for easier visual comparison.

func BenchmarkParse(b *testing.B) {
	bcs := []struct {
		name string
		in   string
	}{
		{"simple", "1.2.3"},
		{"pre", "1.2.3-beta"},
		{"build", "1.2.3+build"},
		{"pre+build", "1.2.3-rc.1+build.info"},
	}

	for _, bc := range bcs {
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = version.NewVersion(bc.in)
			}
		})
	}
}

func BenchmarkParseConstraint(b *testing.B) {
	bcs := []struct {
		name string
		in   string
	}{
		{"equal", "=2.2"},
		{"tilde", "~>2.2.0"},
		{"range", ">=2.2.x, <3.2.0"},

		// unsupported by go-version
		// {"wildcard", "1.x"},
		// {"caret", "^2.2.0"},
		// {"union", "~2.2.0 || =3.3.0"},
	}

	for _, bc := range bcs {
		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = version.NewConstraint(bc.in)
			}
		})
	}
}
