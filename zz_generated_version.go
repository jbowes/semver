//line version.rl:1
// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

import (
	"errors"
)

// This file is generated from semver.rl. DO NOT EDIT.

//line version.rl:81

//line zz_generated_version.go:20
const semver_start int = 1
const semver_error int = 0

const semver_en_main int = 1

//line version.rl:84

var errInvalidVersion = errors.New("invalid version")

// Parse returns the semver version parsed from the string ver. If ver is not
// a valid semver version, an error is returned.
func Parse(ver string) (*Version, error) {
	data := []byte(ver)

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

	v := &Version{}

//line zz_generated_version.go:54
	{
		cs = semver_start
	}

//line version.rl:110

//line zz_generated_version.go:61
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 12:
			goto st_case_12
		case 6:
			goto st_case_6
		case 13:
			goto st_case_13
		case 7:
			goto st_case_7
		case 14:
			goto st_case_14
		case 8:
			goto st_case_8
		case 15:
			goto st_case_15
		case 9:
			goto st_case_9
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		}
		goto st_out
	st_case_1:
		if data[p] == 48 {
			goto tr0
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr0:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line zz_generated_version.go:151
		if data[p] == 46 {
			goto tr3
		}
		goto st0
	tr3:
//line version.rl:51
		v.major = u
//line version.rl:19
		numSeen = false
		u = 0
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line zz_generated_version.go:167
		if data[p] == 48 {
			goto tr4
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr5
		}
		goto st0
	tr4:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line zz_generated_version.go:209
		if data[p] == 46 {
			goto tr6
		}
		goto st0
	tr6:
//line version.rl:52
		v.minor = u
//line version.rl:19
		numSeen = false
		u = 0
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line zz_generated_version.go:225
		if data[p] == 48 {
			goto tr7
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto tr8
		}
		goto st0
	tr7:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line zz_generated_version.go:267
		switch data[p] {
		case 43:
			goto tr23
		case 45:
			goto tr24
		}
		goto st0
	tr23:
//line version.rl:53
		v.patch = u
//line version.rl:19
		numSeen = false
		u = 0
		goto st6
	tr27:
//line version.rl:55

		v.pre = append(v.pre, string(data[s:p]))
		if numSeen {
			u++ // 0 indicates no number seen.
		}
		v.preNum = append(v.preNum, u)

		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line zz_generated_version.go:296
		switch data[p] {
		case 45:
			goto tr9
		case 48:
			goto tr10
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr9
			}
		default:
			goto tr9
		}
		goto st0
	tr13:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st13
	tr9:
//line version.rl:17
		s = p
		goto st13
	tr10:
//line version.rl:17
		s = p
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line zz_generated_version.go:385
		switch data[p] {
		case 45:
			goto st13
		case 46:
			goto st7
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 45:
			goto st13
		case 48:
			goto tr13
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	tr14:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st14
	tr11:
//line version.rl:17
		s = p
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st14
	tr26:
//line version.rl:26

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

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line zz_generated_version.go:521
		switch data[p] {
		case 45:
			goto st13
		case 46:
			goto st7
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	tr24:
//line version.rl:53
		v.patch = u
//line version.rl:19
		numSeen = false
		u = 0
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line zz_generated_version.go:552
		switch data[p] {
		case 45:
			goto tr15
		case 48:
			goto tr16
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr17
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
	tr18:
//line version.rl:17
		s = p
		goto st15
	tr15:
//line version.rl:20

		// optimize for eg "-rc.0"
		v.pre = make([]string, 0, 2)
		v.preNum = make([]uint64, 0, 2)

//line version.rl:17
		s = p
		goto st15
	tr16:
//line version.rl:20

		// optimize for eg "-rc.0"
		v.pre = make([]string, 0, 2)
		v.preNum = make([]uint64, 0, 2)

//line version.rl:18
		numSeen = true
//line version.rl:26

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

//line version.rl:17
		s = p
		goto st15
	tr19:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

//line version.rl:17
		s = p
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line zz_generated_version.go:659
		switch data[p] {
		case 43:
			goto tr27
		case 45:
			goto st15
		case 46:
			goto tr29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	tr29:
//line version.rl:55

		v.pre = append(v.pre, string(data[s:p]))
		if numSeen {
			u++ // 0 indicates no number seen.
		}
		v.preNum = append(v.preNum, u)

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line zz_generated_version.go:696
		switch data[p] {
		case 45:
			goto tr18
		case 48:
			goto tr19
		}
		switch {
		case data[p] < 65:
			if 49 <= data[p] && data[p] <= 57 {
				goto tr20
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr18
			}
		default:
			goto tr18
		}
		goto st0
	tr17:
//line version.rl:20

		// optimize for eg "-rc.0"
		v.pre = make([]string, 0, 2)
		v.preNum = make([]uint64, 0, 2)

//line version.rl:18
		numSeen = true
//line version.rl:26

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

//line version.rl:17
		s = p
		goto st16
	tr20:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

//line version.rl:17
		s = p
		goto st16
	tr30:
//line version.rl:26

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

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line zz_generated_version.go:816
		switch data[p] {
		case 43:
			goto tr27
		case 45:
			goto st15
		case 46:
			goto tr29
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr30
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		goto st0
	tr8:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st17
	tr31:
//line version.rl:26

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

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line zz_generated_version.go:899
		switch data[p] {
		case 43:
			goto tr23
		case 45:
			goto tr24
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr31
		}
		goto st0
	tr5:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st10
	tr21:
//line version.rl:26

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

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line zz_generated_version.go:971
		if data[p] == 46 {
			goto tr6
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr21
		}
		goto st0
	tr2:
//line version.rl:18
		numSeen = true
//line version.rl:26

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

		goto st11
	tr22:
//line version.rl:26

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

		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line zz_generated_version.go:1040
		if data[p] == 46 {
			goto tr3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr22
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 15, 16:
//line version.rl:55

				v.pre = append(v.pre, string(data[s:p]))
				if numSeen {
					u++ // 0 indicates no number seen.
				}
				v.preNum = append(v.preNum, u)

//line version.rl:80
				done = true
			case 13, 14:
//line version.rl:63
				v.build = string(data[s:p])
//line version.rl:80
				done = true
			case 12, 17:
//line version.rl:53
				v.patch = u
//line version.rl:19
				numSeen = false
				u = 0
//line version.rl:80
				done = true
//line zz_generated_version.go:1092
			}
		}

	_out:
		{
		}
	}

//line version.rl:111

	if p != eof || !done {
		return nil, errInvalidVersion
	}

	return v, nil
}
