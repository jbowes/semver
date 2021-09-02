// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package semver provides semantic version and constraint parsing, comparison,
// and testing.
//
// There are many semver packages. This one aims to be the fastest at parsing
// and comparing values, with low memory usage.
//
// Versions can be compared with one another to determine which is newer.
// Constraints specify inclusions and exclusions of semver ranges that a given
// version must satisfy. Typically, constraints are used when expressing a
// dependency.
//
// Parsing of semver values, and comparison of them, strictly follows the rules
// outlined in Semantic Versioning 2.0.0 (https://semver.org). Constraint
// parsing and testing follows the syntax and semantics of npm version specifiers
// (https://semver.npmjs.com/). This behaviour is only what exists now; feature
// requests or contributions are welcome to expand on it!
//
// Normalized Constraint forms
//
// Internally during parsing, and when returning string representations,
// Constraints are normalized in their representation.
//
// Rules for normalization:
//  - replace the empty match-any value "" with *
//  - replace any whitespace with a single space
//  - surround `||` with a single space on each side, if missing
//  - remove '='.
//  - replace partial semver (2.0) with wildcard version (2.0.x)
//  - replace partial wildcard (2.x) with full wildcard (2.x.x)
//    special case: leave full wildcard (*, x, X) as-is, but convert to *.
//  - for wildcard with leading op (including tilde and semver), replace
//    wildcards with 0.
//  - replace remaining wildcard semver (including previous partial) with
//    >= and < clauses (eg 2.x.x becomes >=2.0.0 < 3.0.0)
//  - replace tilde with >= and < clauses (eg ~2.1.1 becomes >=2.1.1 < 2.2.0)
//  - replace caret with >= and < clauses (eg ^2.1.1 becomes >=2.1.1 < 3.0.0)
//  - replace hyphenated ranges with >= and <= clauses
//    (eg 2.1.0 - 2.2.0 becomes >=2.1.0 <=2.2.0)
//
// Normalized constraints are not canonical. Two Constraints can have different
// normalized forms but still be equivalent.
package semver
