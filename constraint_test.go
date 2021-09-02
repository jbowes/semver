// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver_test

import (
	"fmt"
	"testing"

	"github.com/jbowes/semver"
)

func Example() {
	v, err := semver.Parse("1.2.1")
	if err != nil {
		fmt.Println("invalid semver!")
	}

	c, err := semver.ParseConstraint(">=1.1.9 <2.0.0")
	if err != nil {
		fmt.Println("invalid constraint!")
	}

	if c.Check(v) {
		fmt.Printf("%s satisfies %s\n", v, c)
	} else {
		fmt.Printf("%s does not satisfy %s\n", v, c)
	}

	// Output:
	// 1.2.1 satisfies >=1.1.9 <2.0.0
}

func ExampleParseConstraint() {
	c, err := semver.ParseConstraint(">=1.1.9 <2.0.0")
	if err != nil {
		fmt.Println("invalid constraint!")
	}

	fmt.Println(c)

	// Output:
	// >=1.1.9 <2.0.0
}

// Internally, and when String()ified, Constraints are normalized.
func ExampleParseConstraint_normalized() {
	c, err := semver.ParseConstraint("~1.2.7")
	if err != nil {
		fmt.Println("invalid constraint!")
	}

	fmt.Println(c)

	// Output:
	// >=1.2.7 <1.3.0
}

func ExampleConstraint_Check_wildcard() {
	c, err := semver.ParseConstraint("*")
	if err != nil {
		fmt.Println("invalid constraint!")
	}

	v1, err := semver.Parse("1.2.8")
	if err != nil {
		fmt.Println("invalid version!")
	}

	fmt.Println(c.Check(v1))

	// Output:
	// true
}

// Multiple constraints are ANDed together, and may be explicitly ORed.
func ExampleConstraint_Check_or() {
	c, err := semver.ParseConstraint("~1.2.7 || 2.0.0")
	if err != nil {
		fmt.Println("invalid constraint!")
	}

	v1, err := semver.Parse("1.2.8")
	if err != nil {
		fmt.Println("invalid version!")
	}

	v2, err := semver.Parse("2.0.0")
	if err != nil {
		fmt.Println("invalid version!")
	}

	v3, err := semver.Parse("1.3.0")
	if err != nil {
		fmt.Println("invalid version!")
	}

	fmt.Println(c.Check(v1))
	fmt.Println(c.Check(v2))
	fmt.Println(c.Check(v3))

	// Output:
	// true
	// true
	// false
}

func TestParseConstraint(t *testing.T) {
	tcs := map[string]struct {
		in        string
		valid     bool
		canonical string
	}{
		"wildcard":      {"*", true, ">=0.0.0"},
		"wildcard (eq)": {"=*", true, ">=0.0.0"},
		"unary":         {"=1.0", true, ">=1.0.0 <1.1.0"},
		"unary (no op)": {"3.4", true, ">=3.4.0 <3.5.0"},
		"eq (no value)": {"=", false, ""},

		"less than eq":              {"<=4.5.6", true, "<=4.5.6"},
		"less than eq (partial)":    {"<=4.5", true, "<4.6.0"},
		"less than eq (major only)": {"<=4", true, "<5.0.0"},

		"greater than eq":              {">=4.5.6", true, ">=4.5.6"},
		"greater than eq (partial)":    {">=4.5", true, ">=4.5.0"},
		"greater than eq (major only)": {">=4", true, ">=4.0.0"},

		"greater than":              {">4.5.6", true, ">4.5.6"},
		"greater than (partial)":    {">4.5", true, ">=4.6.0"},
		"greater than (major only)": {">4", true, ">=5.0.0"},

		"tilde":            {"~1.2.3", true, ">=1.2.3 <1.3.0"},
		"tilde (wildcard)": {"~4.5.x", true, ">=4.5.0 <4.6.0"},
		"caret":            {"^4.5.6", true, ">=4.5.6 <5.0.0"},
		"caret (no value)": {"^", false, ""},

		"wildcard (minor)":         {"1.x", true, ">=1.0.0 <2.0.0"},
		"wildcard (two wildcards)": {"1.x.x", true, ">=1.0.0 <2.0.0"},

		"range":       {">=2.1.x <3.1.0", true, ">=2.1.0 <3.1.0"},
		"union":       {"~2.0.0 || =3.1.0", true, ">=2.0.0 <2.1.0 || 3.1.0"},
		"range (pre)": {">=2.0.0-rc.1 <2.0.0", true, ">=2.0.0-rc.1 <2.0.0"},

		"invalid (unary no op leading zero)":   {"02.0", false, ""},
		"invalid (union leading zero)":         {"2.01 || 15", false, ""},
		"invalid (union initial leading zero)": {"01 || 15", false, ""},

		"spacing": {">=  2.1.x < 3.1.0", true, ">=2.1.0 <3.1.0"},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			out, err := semver.ParseConstraint(tc.in)
			if (err == nil) != tc.valid {
				t.Error("got unexpected error value:", err)
			}

			if tc.valid {
				outS := out.String()
				if outS != tc.canonical {
					t.Error("got unexpected output. got:", outS, "want:", tc.canonical)
				}
			}
		})
	}
}

func TestConstraint_Check(t *testing.T) {
	tests := map[string]struct {
		constraint string
		version    string
		check      bool
	}{

		"equals":                     {"=2.3.4", "2.3.4", true},
		"equals (no match)":          {"=2.3.4", "1.2.3", false},
		"equals (no op)":             {"2.3.4", "2.3.4", true},
		"equals (partial)":           {"=2.0", "2.0.0", true},
		"equals (partial no match)":  {"=2.0", "1.4.7", false},
		"equals (partial non-exact)": {"=2.0", "2.0.1", true},
		"equals (partial no op)":     {"5.8", "5.8.0", true},

		"wildcard":                    {"*", "4.5.6", true},
		"wildcard (pre)":              {"*", "1.2.3-rc.0", false},
		"wildcard (empty string)":     {"", "4.5.6", true},
		"wildcard (empty string pre)": {"", "1.2.3-rc.0", false},

		"partial wildcard":                  {"2.*", "2.1.5", true},
		"partial wildcard (no match)":       {"1.*", "2.1.5", false},
		"partial wildcard (patch)":          {"2.1.*", "2.1.5", true},
		"partial wildcard (patch no match)": {"2.2.*", "2.1.5", false},
		"partial (no splat)":                {"2", "2.1.5", true},
		"partial (no splat no match)":       {"2", "3.1.5", false},
		"partial (no splat patch)":          {"1.2", "1.2.3", true},
		"partial (no splat patch no match)": {"1.2", "1.3.4", false},

		"less than":            {"<21", "0.1.8", true},
		"less than (no match)": {"<21", "21.7.0", false},

		"less than (no patch)":                        {"<4.5", "0.1.8", true},
		"less than (no patch same minor)":             {"<4.5", "4.5.0", false},
		"less than (no patch same minor large patch)": {"<4.5", "4.5.1", false},

		"less than eq":                         {"<=21", "0.1.8", true},
		"less than eq (same major)":            {"<=21", "21.0.0", true},
		"less than eq (same major with minor)": {"<=21", "21.1.8", true},
		"less than eq (no match)":              {"<=21", "22.0.0", false},

		"less than eq (no patch)":                        {"<=4.5", "0.1.8", true},
		"less than eq (no patch same minor)":             {"<=4.5", "4.5.0", true},
		"less than eq (no patch same minor large patch)": {"<=4.5", "4.5.1", true},
		"less than eq (no patch over)":                   {"<=4.5", "4.6.0", false},

		"greater than":            {">21", "22.1.8", true},
		"greater than (no match)": {">21", "21.7.0", false},

		"greater than (no patch)":                        {">4.5", "5.1.8", true},
		"greater than (no patch same minor)":             {">4.5", "4.5.0", false},
		"greater than (no patch same minor large patch)": {">4.5", "4.5.1", false},

		"greater than eq":                         {">=21", "22.1.8", true},
		"greater than eq (same major)":            {">=21", "21.0.0", true},
		"greater than eq (same major with minor)": {">=21", "21.1.8", true},
		"greater than eq (no match)":              {">=21", "20.0.0", false},

		"greater than eq (no patch)":                        {">=4.5", "4.6.8", true},
		"greater than eq (no patch same minor)":             {">=4.5", "4.5.0", true},
		"greater than eq (no patch same minor large patch)": {">=4.5", "4.5.1", true},
		"greater than eq (no patch under)":                  {">=4.5", "4.4.0", false},

		"greater than (pre)":                     {">0", "0.0.1-rc", false},
		"greater than (pre no match)":            {">0.0", "0.0.1-rc", false},
		"greater than (pre same patch no match)": {">0", "0.0.0-rc", false},
		"greater than (pre is greater)":          {">0.0.0-0", "0.0.0-rc", true},
		"greater than (pre numeric is greater)":  {">1.2.3-rc.1", "1.2.3-rc.2", true},
		"greater than (pre different minor)":     {">1.2.3-rc.1", "1.3.3-rc.2", false}, // npm doesn't compare pre across versions

		"greater than eq (pre)":                            {">=0", "0.0.1-rc", false},
		"greater than eq (pre with minor)":                 {">=0.0", "0.0.1-rc", false},
		"greater than eq (pre no patch)":                   {">=0", "0.0.0-rc", false},
		"greater than eq (constraint pre)":                 {">=0.0.0-0", "0.0.0-rc", true},
		"greater than eq (constraint pre no match)":        {">=0.0.0-0", "3.4.5-rc.1", false},
		"greater than eq (constraint pre, no version pre)": {">=0.0.0-0", "1.2.3", true},

		"tilde":                       {"~1", "1.2.4", true},
		"tilde (no match)":            {"~1", "2.3.4", false},
		"tilde (with minor)":          {"~1.2", "1.2.4", true},
		"tilde (with minor no match)": {"~1.2", "1.3.4", false},
		"tilde (with patch)":          {"~1.2.3", "1.2.4", true},
		"tilde (with patch no match)": {"~1.2.3", "1.3.4", false},
		"tilde (zero major)":          {"~0.2.3", "0.2.5", true},
		"tilde (zero major no match)": {"~0.2.3", "0.3.5", false},
		"tilde (larger rc)":           {"~1.2.3-rc.2", "1.2.3-rc.4", true},

		"caret":                       {"^1.2.3", "1.7.3", true},
		"caret (too large)":           {"^1.2.3", "2.7.3", false},
		"caret (too small)":           {"^1.2.3", "1.2.1", false},
		"caret (no patch)":            {"^1.2", "1.7.3", true},
		"caret (no patch no match)":   {"^1.2", "2.7.3", false},
		"caret (major only)":          {"^1", "1.7.3", true},
		"caret (major only no match)": {"^1", "2.7.3", false},

		"caret (leading zero)":          {"^0.2.3", "0.2.5", true},
		"caret (leading zero no match)": {"^0.2.3", "0.5.6", false},

		"caret (leading zero no patch)":          {"^0.2", "0.2.3", true},
		"caret (leading zero no patch no match)": {"^0.2", "0.4.9", false},

		"caret (patch only)":          {"^0.0.3", "0.0.3", true},
		"caret (patch only no match)": {"^0.0.3", "0.0.4", false},

		"caret (zeros no patch)":                {"^0.0", "0.0.3", true},
		"caret (zeros no patch no match)":       {"^0.0", "0.1.3", false},
		"caret (zeros no patch no match major)": {"^0.0", "1.0.4", false},
		"caret (major zero only)":               {"^0", "0.2.3", true},
		"caret (major zero only no match )":     {"^0", "1.2.3", false},

		"caret (right pre)":                            {"^1.2.0", "1.2.1-rc.1", false},
		"caret (pre equality)":                         {"^0.2.3-rc.2", "0.2.3-rc.2", true},
		"caret (pre same major minor patch)":           {"^0.2.3-rc.2", "0.2.3-rc.4", true},
		"caret (less than pre same major minor patch)": {"^0.2.3-rc.2", "0.2.3-rc.1", false},

		// npm specific: pre-release don't compare across versions.
		"npm pre (diff minor)":           {"^1.2.0-rc.0", "1.2.1-rc.0", false},
		"npm pre (diff minor, diff pre)": {"^1.2.0-rc.0", "1.2.1-rc.1", false},
		"npm pre (gt diff pre)":          {">0.2.3-rc.2", "0.2.4-rc.2", false},
		"npm pre (left pre larger)":      {"^1.2.0-rc.2", "1.2.0-rc.1", false},

		"pre range (less than)":            {">=1.0.0-rc.1 <2.0.0-rc.5", "1.0.0-rc.0", false},
		"pre range (equals left)":          {">=1.0.0-rc.1 <2.0.0-rc.5", "1.0.0-rc.1", true},
		"pre range (larger pre than left)": {">=1.0.0-rc.1 <2.0.0-rc.5", "1.0.0-rc.2", true},
		"pre range (at range no pre)":      {">=1.0.0-rc.1 <2.0.0-rc.5", "1.0.0", true},
		"pre range (within range no pre)":  {">=1.0.0-rc.1 <2.0.0-rc.5", "1.2.0", true},
		// "pre range (within right pre)": {">=1.0.0-rc.1 <2.0.0-rc.5", "2.0.0-rc.4", true},
		"pre range (outside right)":        {">=1.0.0-rc.1 <2.0.0-rc.5", "2.0.0-rc.5", false},
		"pre range (outside right no pre)": {">=1.0.0-rc.1 <2.0.0-rc.5", "2.0.0", false},
		// "pre range (multi within right pre)": {">=1.0.0-rc.1 =1.2.0-rc.2 <2.0.0-rc.5", "1.2.0-rc.3", true},

	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := semver.ParseConstraint(tc.constraint)
			if err != nil {
				t.Error("got unexpected error value:", err)
				return
			}

			v, err := semver.Parse(tc.version)
			if err != nil {
				t.Error("got unexpected error value:", err)
				return
			}

			res := c.Check(v)
			if res != tc.check {
				t.Error("got unexpected result. wanted:", tc.check, "got:", res)
			}
		})
	}
}
