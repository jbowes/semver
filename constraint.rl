// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

import (
    "errors"
)

// This file is generated from constraint.rl. DO NOT EDIT.
%%{
    # (except when you are actually in the file here)

    machine constraint;

    action start { s = p }
    action numstart { numSeen = true }
    action resetnum { numSeen = false; u = 0}
     action partstart {
        // optimize for eg "-rc.0"
        cl.pre = make([]string, 0, 2)
        cl.preNum = make([]uint64, 0, 2)
    }

    action num {
        u *= 10
        switch data[p] {
        case '0':
		case '1':
			u += 1
		case '2':
			u += 2
		case '3':
			u += 3
		case '4':
			u += 4
		case '5':
			u += 5
		case '6':
			u += 6
		case '7':
			u += 7
		case '8':
			u += 8
		case '9':
			u += 9
        }
    }

    action major { cl.major = u }
    action minor { cl.minor  = u }
    action patch { cl.patch = u }

     action part {
        cl.pre = append(cl.pre, string(data[s:p]))
        if (numSeen) {
            u++ // 0 indicates no number seen.
        }
        cl.preNum = append(cl.preNum, u)
    }

    action range { c.clauses = append(c.clauses, make([]*clause, 0, 2)) }
    action endRange { idx++ }
    action simple {
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
                cl.op = eq;
            }
        }

        if expandRange {
            switch {
            case bump == bumpNone: // no-op
            case bump&bumpUnbound > 0:
                cl.op = gt | eq;
            case bump&bumpMajor > 0:
                cl.op = gt | eq;
                c.clauses[idx] = append(c.clauses[idx], &clause{
                    buildlessVersion{ major: cl.major + 1 },
                    lt,
                })
            case bump&bumpMinor > 0:
                cl.op = gt | eq;
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
    }

    num = ('0' | [1-9] [0-9]*) >numstart @num; # TODO: @num might be wrong
    wildcard = [xX\*]; # TODO: is this right?
    major = (num | wildcard >{ bump |= bumpUnbound }) %major %resetnum;
    minor = (num | wildcard >{ bump |= bumpMajor }) %minor %resetnum;
    patch = (num | wildcard >{ bump |= bumpMinor }) %patch %resetnum;

    part  = num | [\-0-9A-Za-z]+ >start %part;
    parts = part ('.' part)*;
    pre   = parts >start >partstart %resetnum;

    # Build isn't used in precedence, so we discard it on constraints
    buildPart = num | [\-0-9A-Za-z]+;
    buildParts = buildPart ('.' buildPart)*;
    build = buildParts;

    qualifier = ('-' pre)? ('+' build)?;

    partial = major ('.' minor >{ field |= hasMinor } ('.' patch >{ field |= hasPatch } qualifier )? )?;

    sp = [ \t]*;

    eq = ('=' sp)? %{ cl.op = eq; expandRange = true };
    gt = ('>' sp) %{ cl.op = gt };
    lt = ('<' sp) %{ cl.op = lt };
    gteq = ('>=' sp) %{ cl.op = gt | eq };
    lteq = ('<=' sp) %{ cl.op = lt | eq };
    tilde = ('~' sp) %{ cl.op = gt | eq; bump = bumpMinor; expandRange = true };
    caret = ('^' sp) %{ cl.op = gt | eq; caret = true; expandRange = true };

    primitive = (eq | gt | lt | gteq | lteq | tilde | caret) partial;
    simple = primitive %simple;
    hyphen = partial ' - ' partial;
    range = (simple ( ' ' simple )* ) >range;
    logicalOr = sp '||' sp;
    nothing = ' '? >range >{ cl.op = eq; expandRange = true; bump |= bumpUnbound } %simple;
    rangeSet = (range (logicalOr range >endRange )*) | nothing;
    main := rangeSet %{ done = true };
}%%

%% write data nofinal;

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
	p := 0  // pointer into data (start)
	pe := len(data) // end pointer
    cs := 0 // current state.
    eof := pe // end of file

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
    alloc := struct{
        cl clause
        c Constraint
        cls [1][]*clause
    // XXX: try pre-allocating 2 clauses here too?
    }{}

    cl := &alloc.cl
    c := &alloc.c
    c.clauses = alloc.cls[:0]

    %%write init;
    %%write exec;

    if p != eof || !done {
        return nil, errInvalidConstraint
    }

    return c, nil
}