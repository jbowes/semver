//line constraint.rl:1
// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

import (
	"errors"
)

// This file is generated from constraint.rl. DO NOT EDIT.

//line constraint.rl:160

//line zz_generated_constraint.go:20
const constraint_start int = 19
const constraint_error int = 0

const constraint_en_main int = 19

//line constraint.rl:163

var errInvalidConstraint = errors.New("invalid constraint")

type bumper uint8

const (
	bumpUnbound bumper = 1 << iota
	bumpMajor
	bumpMinor

	bumpNone bumper = 0
)

type hasField uint8

const (
	hasMinor hasField = 1 << iota
	hasPatch

	hasNone hasField = 0
)

// BUG(jbowes): 1.x.3 is not a valid constraint, and should not parse.
// BUG(jbowes): hyphen constraint ranges are not supported

// ParseConstraint returns the version constraint(s) parsed from the string
// con. If con is not a valid version constraint, an error is returned.
func ParseConstraint(con string) (*Constraint, error) {
	data := []byte(con)

	// Ragel state
	p := 0          // pointer into data (start)
	pe := len(data) // end pointer
	cs := 0         // current state.
	eof := pe       // end of file

	// Not used by us.
	ts, te, act := 0, 0, 0
	_, _, _ = ts, te, act

	var done bool
	var s int
	var u uint64
	var numSeen bool
	var bump bumper
	var expandRange bool
	var field hasField
	var caret bool

	idx := 0

	// Allocate a bunch of our memory in bulk, and pre-allocate the
	// first clause slice.
	alloc := struct {
		cl  clause
		c   Constraint
		cls [1][]*clause
		// XXX: try pre-allocating 2 clauses here too?
	}{}

	cl := &alloc.cl
	c := &alloc.c
	c.clauses = alloc.cls[:0]

	var partalloc *struct {
		str [2]string
		uin [2]uint64
	}

//line zz_generated_constraint.go:96
	{
		cs = constraint_start
	}

//line constraint.rl:231

//line zz_generated_constraint.go:103
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 19:
			goto st_case_19
		case 0:
			goto st_case_0
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 22:
			goto st_case_22
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 23:
			goto st_case_23
		case 15:
			goto st_case_15
		case 24:
			goto st_case_24
		case 16:
			goto st_case_16
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 17:
			goto st_case_17
		case 27:
			goto st_case_27
		case 18:
			goto st_case_18
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		}
		goto st_out
	st_case_19:
		switch data[p] {
		case 32:
			goto tr57
		case 42:
			goto tr58
		case 48:
			goto tr59
		case 60:
			goto tr61
		case 61:
			goto tr62
		case 62:
			goto tr63
		case 88:
			goto tr58
		case 94:
			goto tr64
		case 120:
			goto tr58
		case 126:
			goto tr65
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr60
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr57:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:157
		cl.op = eq
		expandRange = true
		bump |= bumpUnbound
		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line zz_generated_constraint.go:215
		goto st0
	tr4:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr5:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr12:
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr13:
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr21:
//line constraint.rl:146
		cl.op = lt
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr22:
//line constraint.rl:146
		cl.op = lt
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr25:
//line constraint.rl:148
		cl.op = lt | eq
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr26:
//line constraint.rl:148
		cl.op = lt | eq
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr29:
//line constraint.rl:145
		cl.op = gt
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr30:
//line constraint.rl:145
		cl.op = gt
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr33:
//line constraint.rl:147
		cl.op = gt | eq
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr34:
//line constraint.rl:147
		cl.op = gt | eq
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr36:
//line constraint.rl:150
		cl.op = gt | eq
		caret = true
		expandRange = true
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr37:
//line constraint.rl:150
		cl.op = gt | eq
		caret = true
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr39:
//line constraint.rl:149
		cl.op = gt | eq
		bump = bumpMinor
		expandRange = true
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr40:
//line constraint.rl:149
		cl.op = gt | eq
		bump = bumpMinor
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	tr58:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:125
		bump |= bumpUnbound
		goto st21
	tr59:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line zz_generated_constraint.go:387
		switch data[p] {
		case 9:
			goto tr66
		case 32:
			goto tr67
		case 46:
			goto tr68
		case 124:
			goto tr69
		}
		goto st0
	tr66:
//line constraint.rl:36
		cl.major = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st1
	tr71:
//line constraint.rl:37
		cl.minor = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st1
	tr75:
//line constraint.rl:38
		cl.patch = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st1
	tr80:
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st1
	tr85:
//line constraint.rl:40

		cl.pre = append(cl.pre, string(data[s:p]))
		if numSeen {
			u++ // 0 indicates no number seen.
		}
		cl.preNum = append(cl.preNum, u)

//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line zz_generated_constraint.go:801
		switch data[p] {
		case 9:
			goto st1
		case 32:
			goto st1
		case 124:
			goto st2
		}
		goto st0
	tr69:
//line constraint.rl:36
		cl.major = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st2
	tr74:
//line constraint.rl:37
		cl.minor = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st2
	tr79:
//line constraint.rl:38
		cl.patch = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st2
	tr83:
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st2
	tr90:
//line constraint.rl:40

		cl.pre = append(cl.pre, string(data[s:p]))
		if numSeen {
			u++ // 0 indicates no number seen.
		}
		cl.preNum = append(cl.preNum, u)

//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line zz_generated_constraint.go:1213
		if data[p] == 124 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 9:
			goto st3
		case 32:
			goto st3
		case 42:
			goto tr4
		case 48:
			goto tr5
		case 60:
			goto tr7
		case 61:
			goto tr8
		case 62:
			goto tr9
		case 88:
			goto tr4
		case 94:
			goto tr10
		case 120:
			goto tr4
		case 126:
			goto tr11
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr6
		}
		goto st0
	tr6:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr14:
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr23:
//line constraint.rl:146
		cl.op = lt
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr27:
//line constraint.rl:148
		cl.op = lt | eq
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr31:
//line constraint.rl:145
		cl.op = gt
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr35:
//line constraint.rl:147
		cl.op = gt | eq
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr38:
//line constraint.rl:150
		cl.op = gt | eq
		caret = true
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr41:
//line constraint.rl:149
		cl.op = gt | eq
		bump = bumpMinor
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr60:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:144
		cl.op = eq
		expandRange = true
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	tr70:
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line zz_generated_constraint.go:1368
		switch data[p] {
		case 9:
			goto tr66
		case 32:
			goto tr67
		case 46:
			goto tr68
		case 124:
			goto tr69
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr70
		}
		goto st0
	tr67:
//line constraint.rl:36
		cl.major = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st4
	tr72:
//line constraint.rl:37
		cl.minor = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st4
	tr76:
//line constraint.rl:38
		cl.patch = u
//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st4
	tr81:
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st4
	tr86:
//line constraint.rl:40

		cl.pre = append(cl.pre, string(data[s:p]))
		if numSeen {
			u++ // 0 indicates no number seen.
		}
		cl.preNum = append(cl.preNum, u)

//line constraint.rl:19
		numSeen = false
		u = 0
//line constraint.rl:50

		c.clauses[idx] = append(c.clauses[idx], cl)

		switch {
		case field&hasMinor == 0:
			bump |= bumpMajor
		case field&hasPatch == 0:
			bump |= bumpMinor
		}

		if caret {
			switch {
			case cl.major > 0:
				bump |= bumpMajor
			case cl.minor > 0:
				bump |= bumpMinor
			default:
				cl.op = eq
			}
		}

		if expandRange {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0:
				cl.op = gt | eq
			case bump&bumpMajor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{major: cl.major + 1},
					lt,
				})
			case bump&bumpMinor > 0:
				cl.op = gt | eq
				c.clauses[idx] = append(c.clauses[idx], &clause{
					buildlessVersion{
						major: cl.major,
						minor: cl.minor + 1,
					},
					lt,
				})
			}
		} else {
			switch {
			case bump == bumpNone: // no-op
			case bump&bumpUnbound > 0 && cl.op == gt:
				fallthrough
			case bump&bumpUnbound > 0 && cl.op == lt:
				// tautaulogy.
				cl.op = lt // already major minor and patch of 0.
			case bump&bumpMajor > 0 && cl.op == gt:
				cl.major += 1
				cl.op |= eq
			case bump&bumpMajor > 0 && cl.op == lt|eq:
				cl.major += 1
				cl.op ^= eq
			case bump&bumpMinor > 0 && cl.op == gt:
				cl.minor += 1
				cl.op |= eq
			case bump&bumpMinor > 0 && cl.op == lt|eq:
				cl.minor += 1
				cl.op ^= eq
			}
		}

		bump = bumpNone
		expandRange = false
		field = hasNone
		caret = false

		cl = &clause{}

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line zz_generated_constraint.go:1785
		switch data[p] {
		case 9:
			goto st1
		case 32:
			goto st1
		case 42:
			goto tr12
		case 48:
			goto tr13
		case 60:
			goto st5
		case 61:
			goto st8
		case 62:
			goto st9
		case 88:
			goto tr12
		case 94:
			goto st12
		case 120:
			goto tr12
		case 124:
			goto st2
		case 126:
			goto st13
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
	tr7:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st5
	tr61:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line zz_generated_constraint.go:1831
		switch data[p] {
		case 9:
			goto st6
		case 32:
			goto st6
		case 42:
			goto tr21
		case 48:
			goto tr22
		case 61:
			goto st7
		case 88:
			goto tr21
		case 120:
			goto tr21
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr23
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 9:
			goto st6
		case 32:
			goto st6
		case 42:
			goto tr21
		case 48:
			goto tr22
		case 88:
			goto tr21
		case 120:
			goto tr21
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr23
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 9:
			goto st7
		case 32:
			goto st7
		case 42:
			goto tr25
		case 48:
			goto tr26
		case 88:
			goto tr25
		case 120:
			goto tr25
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr27
		}
		goto st0
	tr8:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st8
	tr62:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line zz_generated_constraint.go:1913
		switch data[p] {
		case 9:
			goto st8
		case 32:
			goto st8
		case 42:
			goto tr12
		case 48:
			goto tr13
		case 88:
			goto tr12
		case 120:
			goto tr12
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
	tr9:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st9
	tr63:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line zz_generated_constraint.go:1947
		switch data[p] {
		case 9:
			goto st10
		case 32:
			goto st10
		case 42:
			goto tr29
		case 48:
			goto tr30
		case 61:
			goto st11
		case 88:
			goto tr29
		case 120:
			goto tr29
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr31
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 9:
			goto st10
		case 32:
			goto st10
		case 42:
			goto tr29
		case 48:
			goto tr30
		case 88:
			goto tr29
		case 120:
			goto tr29
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr31
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 9:
			goto st11
		case 32:
			goto st11
		case 42:
			goto tr33
		case 48:
			goto tr34
		case 88:
			goto tr33
		case 120:
			goto tr33
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr35
		}
		goto st0
	tr10:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st12
	tr64:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line zz_generated_constraint.go:2029
		switch data[p] {
		case 9:
			goto st12
		case 32:
			goto st12
		case 42:
			goto tr36
		case 48:
			goto tr37
		case 88:
			goto tr36
		case 120:
			goto tr36
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr38
		}
		goto st0
	tr11:
//line constraint.rl:49
		idx++
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st13
	tr65:
//line constraint.rl:48
		c.clauses = append(c.clauses, make([]*clause, 0, 2))
		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line zz_generated_constraint.go:2063
		switch data[p] {
		case 9:
			goto st13
		case 32:
			goto st13
		case 42:
			goto tr39
		case 48:
			goto tr40
		case 88:
			goto tr39
		case 120:
			goto tr39
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr41
		}
		goto st0
	tr68:
//line constraint.rl:36
		cl.major = u
//line constraint.rl:19
		numSeen = false
		u = 0
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line zz_generated_constraint.go:2093
		switch data[p] {
		case 42:
			goto tr42
		case 48:
			goto tr43
		case 88:
			goto tr42
		case 120:
			goto tr42
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr44
		}
		goto st0
	tr42:
//line constraint.rl:140
		field |= hasMinor
//line constraint.rl:126
		bump |= bumpMajor
		goto st23
	tr43:
//line constraint.rl:140
		field |= hasMinor
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line zz_generated_constraint.go:2130
		switch data[p] {
		case 9:
			goto tr71
		case 32:
			goto tr72
		case 46:
			goto tr73
		case 124:
			goto tr74
		}
		goto st0
	tr73:
//line constraint.rl:37
		cl.minor = u
//line constraint.rl:19
		numSeen = false
		u = 0
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line zz_generated_constraint.go:2153
		switch data[p] {
		case 42:
			goto tr45
		case 48:
			goto tr46
		case 88:
			goto tr45
		case 120:
			goto tr45
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr47
		}
		goto st0
	tr45:
//line constraint.rl:140
		field |= hasPatch
//line constraint.rl:127
		bump |= bumpMinor
		goto st24
	tr46:
//line constraint.rl:140
		field |= hasPatch
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line zz_generated_constraint.go:2190
		switch data[p] {
		case 9:
			goto tr75
		case 32:
			goto tr76
		case 43:
			goto tr77
		case 45:
			goto tr78
		case 124:
			goto tr79
		}
		goto st0
	tr77:
//line constraint.rl:38
		cl.patch = u
//line constraint.rl:19
		numSeen = false
		u = 0
		goto st16
	tr87:
//line constraint.rl:40

		cl.pre = append(cl.pre, string(data[s:p]))
		if numSeen {
			u++ // 0 indicates no number seen.
		}
		cl.preNum = append(cl.preNum, u)

//line constraint.rl:19
		numSeen = false
		u = 0
		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line zz_generated_constraint.go:2227
		switch data[p] {
		case 45:
			goto st25
		case 48:
			goto tr49
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr50
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st25
			}
		default:
			goto st25
		}
		goto st0
	tr49:
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line zz_generated_constraint.go:2261
		switch data[p] {
		case 9:
			goto tr80
		case 32:
			goto tr81
		case 45:
			goto st25
		case 46:
			goto st16
		case 124:
			goto tr83
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st25
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st25
			}
		default:
			goto st25
		}
		goto st0
	tr50:
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st26
	tr84:
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line zz_generated_constraint.go:2308
		switch data[p] {
		case 9:
			goto tr80
		case 32:
			goto tr81
		case 45:
			goto st25
		case 46:
			goto st16
		case 124:
			goto tr83
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr84
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st25
			}
		default:
			goto st25
		}
		goto st0
	tr78:
//line constraint.rl:38
		cl.patch = u
//line constraint.rl:19
		numSeen = false
		u = 0
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line zz_generated_constraint.go:2345
		switch data[p] {
		case 45:
			goto tr51
		case 48:
			goto tr52
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr53
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr51
			}
		default:
			goto tr51
		}
		goto st0
	tr51:
//line constraint.rl:17
		s = p
//line constraint.rl:20

		partalloc = new(struct {
			str [2]string
			uin [2]uint64
		})

		// optimize for eg "-rc.0"
		cl.pre = partalloc.str[:0]
		cl.preNum = partalloc.uin[:0]

		goto st27
	tr52:
//line constraint.rl:17
		s = p
//line constraint.rl:20

		partalloc = new(struct {
			str [2]string
			uin [2]uint64
		})

		// optimize for eg "-rc.0"
		cl.pre = partalloc.str[:0]
		cl.preNum = partalloc.uin[:0]

//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st27
	tr54:
//line constraint.rl:17
		s = p
		goto st27
	tr55:
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

//line constraint.rl:17
		s = p
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line zz_generated_constraint.go:2422
		switch data[p] {
		case 9:
			goto tr85
		case 32:
			goto tr86
		case 43:
			goto tr87
		case 45:
			goto st27
		case 46:
			goto tr89
		case 124:
			goto tr90
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st27
			}
		default:
			goto st27
		}
		goto st0
	tr89:
//line constraint.rl:40

		cl.pre = append(cl.pre, string(data[s:p]))
		if numSeen {
			u++ // 0 indicates no number seen.
		}
		cl.preNum = append(cl.preNum, u)

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line zz_generated_constraint.go:2465
		switch data[p] {
		case 45:
			goto tr54
		case 48:
			goto tr55
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr56
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr54
			}
		default:
			goto tr54
		}
		goto st0
	tr53:
//line constraint.rl:17
		s = p
//line constraint.rl:20

		partalloc = new(struct {
			str [2]string
			uin [2]uint64
		})

		// optimize for eg "-rc.0"
		cl.pre = partalloc.str[:0]
		cl.preNum = partalloc.uin[:0]

//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st28
	tr56:
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

//line constraint.rl:17
		s = p
		goto st28
	tr91:
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line zz_generated_constraint.go:2530
		switch data[p] {
		case 9:
			goto tr85
		case 32:
			goto tr86
		case 43:
			goto tr87
		case 45:
			goto st27
		case 46:
			goto tr89
		case 124:
			goto tr90
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr91
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st27
			}
		default:
			goto st27
		}
		goto st0
	tr47:
//line constraint.rl:140
		field |= hasPatch
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st29
	tr92:
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line zz_generated_constraint.go:2581
		switch data[p] {
		case 9:
			goto tr75
		case 32:
			goto tr76
		case 43:
			goto tr77
		case 45:
			goto tr78
		case 124:
			goto tr79
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr92
		}
		goto st0
	tr44:
//line constraint.rl:140
		field |= hasMinor
//line constraint.rl:18
		numSeen = true
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st30
	tr93:
//line constraint.rl:31

		u *= 10
		u += uint64(data[p] - 48)

		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line zz_generated_constraint.go:2621
		switch data[p] {
		case 9:
			goto tr71
		case 32:
			goto tr72
		case 46:
			goto tr73
		case 124:
			goto tr74
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr93
		}
		goto st0
	st_out:
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 20, 25, 26:
//line constraint.rl:50

				c.clauses[idx] = append(c.clauses[idx], cl)

				switch {
				case field&hasMinor == 0:
					bump |= bumpMajor
				case field&hasPatch == 0:
					bump |= bumpMinor
				}

				if caret {
					switch {
					case cl.major > 0:
						bump |= bumpMajor
					case cl.minor > 0:
						bump |= bumpMinor
					default:
						cl.op = eq
					}
				}

				if expandRange {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0:
						cl.op = gt | eq
					case bump&bumpMajor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{major: cl.major + 1},
							lt,
						})
					case bump&bumpMinor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{
								major: cl.major,
								minor: cl.minor + 1,
							},
							lt,
						})
					}
				} else {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0 && cl.op == gt:
						fallthrough
					case bump&bumpUnbound > 0 && cl.op == lt:
						// tautaulogy.
						cl.op = lt // already major minor and patch of 0.
					case bump&bumpMajor > 0 && cl.op == gt:
						cl.major += 1
						cl.op |= eq
					case bump&bumpMajor > 0 && cl.op == lt|eq:
						cl.major += 1
						cl.op ^= eq
					case bump&bumpMinor > 0 && cl.op == gt:
						cl.minor += 1
						cl.op |= eq
					case bump&bumpMinor > 0 && cl.op == lt|eq:
						cl.minor += 1
						cl.op ^= eq
					}
				}

				bump = bumpNone
				expandRange = false
				field = hasNone
				caret = false

				cl = &clause{}

//line constraint.rl:159
				done = true
			case 21, 22:
//line constraint.rl:36
				cl.major = u
//line constraint.rl:19
				numSeen = false
				u = 0
//line constraint.rl:50

				c.clauses[idx] = append(c.clauses[idx], cl)

				switch {
				case field&hasMinor == 0:
					bump |= bumpMajor
				case field&hasPatch == 0:
					bump |= bumpMinor
				}

				if caret {
					switch {
					case cl.major > 0:
						bump |= bumpMajor
					case cl.minor > 0:
						bump |= bumpMinor
					default:
						cl.op = eq
					}
				}

				if expandRange {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0:
						cl.op = gt | eq
					case bump&bumpMajor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{major: cl.major + 1},
							lt,
						})
					case bump&bumpMinor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{
								major: cl.major,
								minor: cl.minor + 1,
							},
							lt,
						})
					}
				} else {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0 && cl.op == gt:
						fallthrough
					case bump&bumpUnbound > 0 && cl.op == lt:
						// tautaulogy.
						cl.op = lt // already major minor and patch of 0.
					case bump&bumpMajor > 0 && cl.op == gt:
						cl.major += 1
						cl.op |= eq
					case bump&bumpMajor > 0 && cl.op == lt|eq:
						cl.major += 1
						cl.op ^= eq
					case bump&bumpMinor > 0 && cl.op == gt:
						cl.minor += 1
						cl.op |= eq
					case bump&bumpMinor > 0 && cl.op == lt|eq:
						cl.minor += 1
						cl.op ^= eq
					}
				}

				bump = bumpNone
				expandRange = false
				field = hasNone
				caret = false

				cl = &clause{}

//line constraint.rl:159
				done = true
			case 23, 30:
//line constraint.rl:37
				cl.minor = u
//line constraint.rl:19
				numSeen = false
				u = 0
//line constraint.rl:50

				c.clauses[idx] = append(c.clauses[idx], cl)

				switch {
				case field&hasMinor == 0:
					bump |= bumpMajor
				case field&hasPatch == 0:
					bump |= bumpMinor
				}

				if caret {
					switch {
					case cl.major > 0:
						bump |= bumpMajor
					case cl.minor > 0:
						bump |= bumpMinor
					default:
						cl.op = eq
					}
				}

				if expandRange {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0:
						cl.op = gt | eq
					case bump&bumpMajor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{major: cl.major + 1},
							lt,
						})
					case bump&bumpMinor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{
								major: cl.major,
								minor: cl.minor + 1,
							},
							lt,
						})
					}
				} else {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0 && cl.op == gt:
						fallthrough
					case bump&bumpUnbound > 0 && cl.op == lt:
						// tautaulogy.
						cl.op = lt // already major minor and patch of 0.
					case bump&bumpMajor > 0 && cl.op == gt:
						cl.major += 1
						cl.op |= eq
					case bump&bumpMajor > 0 && cl.op == lt|eq:
						cl.major += 1
						cl.op ^= eq
					case bump&bumpMinor > 0 && cl.op == gt:
						cl.minor += 1
						cl.op |= eq
					case bump&bumpMinor > 0 && cl.op == lt|eq:
						cl.minor += 1
						cl.op ^= eq
					}
				}

				bump = bumpNone
				expandRange = false
				field = hasNone
				caret = false

				cl = &clause{}

//line constraint.rl:159
				done = true
			case 24, 29:
//line constraint.rl:38
				cl.patch = u
//line constraint.rl:19
				numSeen = false
				u = 0
//line constraint.rl:50

				c.clauses[idx] = append(c.clauses[idx], cl)

				switch {
				case field&hasMinor == 0:
					bump |= bumpMajor
				case field&hasPatch == 0:
					bump |= bumpMinor
				}

				if caret {
					switch {
					case cl.major > 0:
						bump |= bumpMajor
					case cl.minor > 0:
						bump |= bumpMinor
					default:
						cl.op = eq
					}
				}

				if expandRange {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0:
						cl.op = gt | eq
					case bump&bumpMajor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{major: cl.major + 1},
							lt,
						})
					case bump&bumpMinor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{
								major: cl.major,
								minor: cl.minor + 1,
							},
							lt,
						})
					}
				} else {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0 && cl.op == gt:
						fallthrough
					case bump&bumpUnbound > 0 && cl.op == lt:
						// tautaulogy.
						cl.op = lt // already major minor and patch of 0.
					case bump&bumpMajor > 0 && cl.op == gt:
						cl.major += 1
						cl.op |= eq
					case bump&bumpMajor > 0 && cl.op == lt|eq:
						cl.major += 1
						cl.op ^= eq
					case bump&bumpMinor > 0 && cl.op == gt:
						cl.minor += 1
						cl.op |= eq
					case bump&bumpMinor > 0 && cl.op == lt|eq:
						cl.minor += 1
						cl.op ^= eq
					}
				}

				bump = bumpNone
				expandRange = false
				field = hasNone
				caret = false

				cl = &clause{}

//line constraint.rl:159
				done = true
			case 27, 28:
//line constraint.rl:40

				cl.pre = append(cl.pre, string(data[s:p]))
				if numSeen {
					u++ // 0 indicates no number seen.
				}
				cl.preNum = append(cl.preNum, u)

//line constraint.rl:19
				numSeen = false
				u = 0
//line constraint.rl:50

				c.clauses[idx] = append(c.clauses[idx], cl)

				switch {
				case field&hasMinor == 0:
					bump |= bumpMajor
				case field&hasPatch == 0:
					bump |= bumpMinor
				}

				if caret {
					switch {
					case cl.major > 0:
						bump |= bumpMajor
					case cl.minor > 0:
						bump |= bumpMinor
					default:
						cl.op = eq
					}
				}

				if expandRange {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0:
						cl.op = gt | eq
					case bump&bumpMajor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{major: cl.major + 1},
							lt,
						})
					case bump&bumpMinor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{
								major: cl.major,
								minor: cl.minor + 1,
							},
							lt,
						})
					}
				} else {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0 && cl.op == gt:
						fallthrough
					case bump&bumpUnbound > 0 && cl.op == lt:
						// tautaulogy.
						cl.op = lt // already major minor and patch of 0.
					case bump&bumpMajor > 0 && cl.op == gt:
						cl.major += 1
						cl.op |= eq
					case bump&bumpMajor > 0 && cl.op == lt|eq:
						cl.major += 1
						cl.op ^= eq
					case bump&bumpMinor > 0 && cl.op == gt:
						cl.minor += 1
						cl.op |= eq
					case bump&bumpMinor > 0 && cl.op == lt|eq:
						cl.minor += 1
						cl.op ^= eq
					}
				}

				bump = bumpNone
				expandRange = false
				field = hasNone
				caret = false

				cl = &clause{}

//line constraint.rl:159
				done = true
			case 19:
//line constraint.rl:48
				c.clauses = append(c.clauses, make([]*clause, 0, 2))
//line constraint.rl:157
				cl.op = eq
				expandRange = true
				bump |= bumpUnbound
//line constraint.rl:50

				c.clauses[idx] = append(c.clauses[idx], cl)

				switch {
				case field&hasMinor == 0:
					bump |= bumpMajor
				case field&hasPatch == 0:
					bump |= bumpMinor
				}

				if caret {
					switch {
					case cl.major > 0:
						bump |= bumpMajor
					case cl.minor > 0:
						bump |= bumpMinor
					default:
						cl.op = eq
					}
				}

				if expandRange {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0:
						cl.op = gt | eq
					case bump&bumpMajor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{major: cl.major + 1},
							lt,
						})
					case bump&bumpMinor > 0:
						cl.op = gt | eq
						c.clauses[idx] = append(c.clauses[idx], &clause{
							buildlessVersion{
								major: cl.major,
								minor: cl.minor + 1,
							},
							lt,
						})
					}
				} else {
					switch {
					case bump == bumpNone: // no-op
					case bump&bumpUnbound > 0 && cl.op == gt:
						fallthrough
					case bump&bumpUnbound > 0 && cl.op == lt:
						// tautaulogy.
						cl.op = lt // already major minor and patch of 0.
					case bump&bumpMajor > 0 && cl.op == gt:
						cl.major += 1
						cl.op |= eq
					case bump&bumpMajor > 0 && cl.op == lt|eq:
						cl.major += 1
						cl.op ^= eq
					case bump&bumpMinor > 0 && cl.op == gt:
						cl.minor += 1
						cl.op |= eq
					case bump&bumpMinor > 0 && cl.op == lt|eq:
						cl.minor += 1
						cl.op ^= eq
					}
				}

				bump = bumpNone
				expandRange = false
				field = hasNone
				caret = false

				cl = &clause{}

//line constraint.rl:159
				done = true
//line zz_generated_constraint.go:3152
			}
		}

	_out:
		{
		}
	}

//line constraint.rl:232

	if p != eof || !done {
		return nil, errInvalidConstraint
	}

	return c, nil
}
