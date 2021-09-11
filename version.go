// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

import (
	"fmt"
	"strings"
)

// Version is a valid semver version as defined by https://semver.org
type Version struct {
	buildlessVersion
	build string
}

// buildlessVersion is version without the build info, as that's not used
// for comparison. This lets us reuse the struct and logic.
type buildlessVersion struct {
	major uint64
	minor uint64
	patch uint64

	pre []string
	// 0 indicates it wasn't a number, everything else + 1
	preNum []uint64
}

// String returns the string representation of the Version.
func (v *Version) String() string {
	if v == nil {
		return ""
	}

	b := strings.Builder{}
	v.buildlessVersion.write(&b)

	if len(v.build) != 0 {
		b.WriteRune('+')
		b.WriteString(v.build)
	}

	return b.String()
}

// write is pulled out here so we can call it from Constraint.
func (v *buildlessVersion) write(b *strings.Builder) {
	fmt.Fprintf(b, "%d.%d.%d", v.major, v.minor, v.patch)
	if len(v.pre) != 0 {
		b.WriteRune('-')
		for i, p := range v.pre {
			b.WriteString(p)
			if i != len(v.pre)-1 {
				b.WriteRune('.')
			}
		}
	}
}

// Compare is broken in to two steps, so that we may easily re-use
// compareNonPre (major, minor, patch) comparison in constraints,
// and honor particulars for ranges and pre.

// Compare returns an integer comparing semver Version v with other,
// following the rules outlined in the semver spec at
// https://semver.org/#spec-item-11
//
// If v has higher precedence than other, 1 is returned. If other has higher
// precedence, -1 is returned. If the two versions are equal, zero is returned.
// A nil Version has lower precedence than a non-nil version. Two nil versions
// have the same precedence.
func (v *Version) Compare(other *Version) int {
	switch {
	case v == nil && other == nil:
		return 0
	case v == nil:
		return -1
	case other == nil:
		return 1
	}

	cmp := v.compareNonPre(other)
	if cmp == 0 {
		cmp = v.comparePre(other)
	}

	return cmp
}

// compareNonPre compares the major.minor.patch part of two semver versions,
// (that is, the non-prerelease parts).
func (v *buildlessVersion) compareNonPre(other *Version) int {
	switch {
	case v.major < other.major:
		return -1
	case v.major > other.major:
		return 1
	case v.minor < other.minor:
		return -1
	case v.minor > other.minor:
		return 1
	case v.patch < other.patch:
		return -1
	case v.patch > other.patch:
		return 1
	case len(v.pre) > 0 && len(other.pre) == 0:
		return -1
	case len(v.pre) == 0 && len(other.pre) > 0:
		return 1
	default:
		return 0
	}
}

// comparePre compares two prelease parts of a semver version,
// assuming their major.minor.patch parts have already been compared.
func (v *buildlessVersion) comparePre(other *Version) int {
	for i := range v.pre {
		if i == len(other.pre) {
			return 1
		}

		if v.preNum[i] > 0 && other.preNum[i] > 0 {
			if v.preNum[i] < other.preNum[i] {
				return -1
			} else if v.preNum[i] > other.preNum[i] {
				return 1
			}

		} else {
			o := strings.Compare(v.pre[i], other.pre[i])
			if o != 0 {
				return o
			}
		}
	}

	if len(v.pre) < len(other.pre) {
		return -1
	}

	return 0

	// build is not used for semver comparison.
}

// Prerelease returns the prerelease portion of a Version. If the version had
// no prerelease or is nil, the empty string is returned.
//
// Note: This method is considered experimental, and may be removed or changed
// before v1.0.0
func (v *Version) Prerelease() string {
	if v == nil {
		return ""
	}

	return strings.Join(v.pre, ".")
}
