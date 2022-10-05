// Copyright (c) 2021-2022 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

import (
    "errors"
)

// This file is generated from semver.rl. DO NOT EDIT.
%%{
    # (except when you are actually in the file here)

    machine semver;

    action start { s = p }
    action numstart { numSeen = true }
    action resetnum {numSeen = false; u = 0}
    action partstart {
        partalloc = new(struct{
            str [2]string
            uin [2]uint64
        })

        // optimize for eg "-rc.0"
        v.pre = partalloc.str[:0]
        v.preNum = partalloc.uin[:0]
    }

    action num {
        u *= 10
        u += uint64(data[p] - 48)
    }

    action major { v.major = u }
    action minor { v.minor  = u }
    action patch { v.patch = u }

    action part {
        v.pre = append(v.pre, string(data[s:p]))
        if (numSeen) {
            u++ // 0 indicates no number seen.
        }
        v.preNum = append(v.preNum, u)
    }

    action build { v.build = string(data[s:p]) }

    num = ('0' | [1-9] [0-9]*) >numstart @num;
    major = num %major %resetnum;
    minor = num %minor %resetnum;
    patch = num %patch %resetnum;

    part  = num | [\-0-9A-Za-z]+ >start %part;
    parts = part ('.' part)*;
    pre   = parts >partstart;
    buildPart = num | [\-0-9A-Za-z]+;
    buildParts = buildPart ('.' buildPart)*;
    build = buildParts >start %build;

    qualifier = ('-' pre)? ('+' build)?;

    semver = major '.' minor '.' patch qualifier;
    main := semver %{ done = true };
}%%

%% write data nofinal;

var errInvalidVersion = errors.New("invalid version")

// Parse returns the semver version parsed from the string ver. If ver is not
// a valid semver version, an error is returned.
func Parse(ver string) (*Version, error) {
    data := []byte(ver)

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

    var partalloc *struct{
        str [2]string
        uin [2]uint64
    }

    v := &Version{}

    %%write init;
    %%write exec;

    if p != eof || !done {
        return nil, errInvalidVersion
    }

    return v, nil
}