// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

import (
	"strings"
)

// Constraint is a constraint specifier on a set of versions.
// Constraint specifiers follow the syntax (for now) of npm:
// https://semver.npmjs.com/
type Constraint struct {
	// Disjunctive normal form; ors of ands
	clauses [][]*clause // XXX: check benchmark with non-pointer
}

type clause struct {
	buildlessVersion
	op op
}

type op byte

const (
	eq op = 1 << iota
	lt
	gt
)

// String returns the normalized string representation of a Constraint.
func (c *Constraint) String() string {
	if c == nil {
		return ""
	}

	b := strings.Builder{}

	for i, andC := range c.clauses {
		for j, clause := range andC {
			clause.write(&b)
			if j != len(andC)-1 {
				b.WriteRune(' ')
			}
		}

		if i != len(c.clauses)-1 {
			b.WriteString(" || ")
		}
	}

	return b.String()
}

func (c *clause) write(b *strings.Builder) {
	switch c.op {
	case eq: // nothing.
	case lt:
		b.WriteRune('<')
	case gt:
		b.WriteRune('>')
	case lt | eq:
		b.WriteString("<=")
	case gt | eq:
		b.WriteString(">=")
	}

	c.buildlessVersion.write(b)
}

// BUG(jbowes): Multiple prerelease versions are not compared within a constraint, only the first.

// Check returns a bool indicating if Version v satisfies Constraint c.
func (c *Constraint) Check(v *Version) bool {
	for _, andC := range c.clauses {
		i := 0

		// if we have a direct match on one of the clauses with a pre, then we
		// ignore the pre for the other clauses (at least for npm)

		preMatch := false
		for ; i < len(andC); i++ {
			clause := andC[i]

			// If the constraint doesn't have pre, *no* pre will match
			if len(clause.pre) == 0 && len(v.pre) != 0 && !preMatch {
				break
			}

			cmp := clause.compareNonPre(v)

			if !preMatch {
				// TODO: npm doesn't allow cmp of pre across versions
				if cmp == 0 { // same major.minor.patch
					cmp = clause.comparePre(v)
				} else if len(v.pre) > 0 {
					break // !ok
				}
			}

			var ok bool
			switch cmp {
			case -1:
				ok = clause.op&gt > 0
			case 0:
				ok = clause.op&eq > 0
			case 1:
				ok = clause.op&lt > 0
			}

			if ok && len(v.pre) > 0 {
				preMatch = true
			}

			if !ok {
				break
			}
		}

		if i == len(andC) {
			return true
		}
	}

	return false
}
