// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver_test

import (
	"fmt"
	"testing"

	"github.com/jbowes/semver"
)

func ExampleParse() {
	v, err := semver.Parse("1.2.0-rc.0+happysunshine")
	if err != nil {
		fmt.Println("invalid semver!")
	}

	fmt.Println(v)
	// Output:
	// 1.2.0-rc.0+happysunshine
}

func ExampleVersion_Compare() {
	v1, err := semver.Parse("1.2.0-rc.0+happysunshine")
	if err != nil {
		fmt.Println("invalid semver!")
	}

	v2, err := semver.Parse("1.4.0")
	if err != nil {
		fmt.Println("invalid semver!")
	}

	fmt.Println(v1.Compare(v2))
	// Output:
	// -1
}

// In accordance with the semver spec, build fields are ignored for determining
// version precedence.
func ExampleVersion_Compare_buildField() {
	v1, err := semver.Parse("1.0.0+somebuild")
	if err != nil {
		fmt.Println("invalid semver!")
	}

	v2, err := semver.Parse("1.0.0+builds.are.ignored.for.comparison")
	if err != nil {
		fmt.Println("invalid semver!")
	}

	fmt.Println(v1.Compare(v2))
	// Output:
	// 0
}

func TestParse(t *testing.T) {
	tcs := map[string]struct {
		in    string
		valid bool
	}{
		"simple":      {"1.0.0", true},
		"multi digit": {"10.22.345", true},

		"pre":              {"1.0.0-alpha", true},
		"pre (multi-part)": {"1.0.0-alpha.0.2", true},

		"build":               {"1.0.0+build", true},
		"pre+build":           {"1.0.0-alpha.1+build.data", true},
		"pre+build (hyphens)": {"1.0.0-alp-ha.1+me-ta.data", true},

		"invalid (leading zero)":       {"1.0.02", false},
		"invalid (many leading zeros)": {"01.02.03", false},
		"invalid (double semver)":      {"1.0.01.0.1", false},
		"invalid (slash)":              {"1.0.1-r\\c.0", false},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			out, err := semver.Parse(tc.in)
			if (err == nil) != tc.valid {
				t.Error("got unexpected error:", err)
			}

			if tc.valid {
				outS := out.String()
				if outS != tc.in {
					t.Error("got unexpected output. got:", outS, "want:", tc.in)
				}
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tcs := map[string]struct {
		v1  string
		v2  string
		out int
	}{
		"basic equality": {"1.2.3", "1.2.3", 0},

		"less than (major)": {"1.2.3", "2.2.3", -1},
		"less than (minor)": {"1.2.3", "1.3.3", -1},
		"less than (patch)": {"1.2.3", "1.2.4", -1},

		"build ignored":        {"1.2.3+foo", "1.2.3", 0},
		"build ignored (both)": {"1.2.3+foo", "1.2.3+bar", 0},

		"pre equality":              {"1.2.3-rc", "1.2.3-rc", 0},
		"pre equality (multi part)": {"1.2.3-rc.0", "1.2.3-rc.0", 0},

		"pre build ignored":        {"1.2.3-rc.0+foo", "1.2.3-rc.0", 0},
		"pre build ignored (both)": {"1.2.3-rc.0+foo", "1.2.3-rc.0+bar", 0},

		"less than (same pre)": {"1.2.3-rc", "2.2.3-rc", -1},

		"pre less than no pre": {"1.2.3-rc", "1.2.3", -1},

		"lexicographic pre comparison":                     {"1.2.3-beta", "1.2.3-alpha", 1},
		"lexicographic pre comparison (multi-part)":        {"1.2.3-alpha.beta", "1.2.3-beta.beta", -1},
		"lexicographic pre comparison (uneven multi-part)": {"1.2.3-alpha.beta", "1.2.3-beta", -1},

		"more pre parts wins": {"1.2.3-alpha.first", "1.2.3-alpha", 1},

		"pre digit comparison":               {"1.2.3-rc.0", "1.2.3-rc.1", -1},
		"pre digit comparison (multi digit)": {"1.2.3-rc.2", "1.2.3-rc.10", -1},
		"pre digit comparison (multi part)":  {"1.2.3-rc.0.0", "1.2.3-rc.0.1", -1},
		"pre leading zeros aren't digits":    {"1.2.3-rc.01", "1.2.3-rc.1", -1},

		"pre comparison (mixed alpha and digit)": {"1.2.3-rc", "1.2.3-rc.2", -1},

		// According to semver, `-` is a non-digit and so eg
		// 3 is less than -2 for pre. This is not what masterminds/semver does.
		"pre comparison (negatives aren't numeric)": {"1.2.3-rc.3", "1.2.3-rc.-2", 1},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			v1, err := semver.Parse(tc.v1)
			if err != nil {
				t.Error("couldn't parse v1:", err)
			}

			v2, err := semver.Parse(tc.v2)
			if err != nil {
				t.Error("couldn't parse v2:", err)
			}

			cmp := v1.Compare(v2)
			if cmp != tc.out {
				t.Errorf("comparison failed. expected: %d, got: %d", tc.out, cmp)
			}

			cmp = v2.Compare(v1)
			if cmp != -1*tc.out {
				t.Errorf("inverse comparison failed. expected: %d, got: %d", -1*tc.out, cmp)

			}
		})
	}
}
