// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

//go:generate -command rgl ragel -Z -G2
//go:generate -command fmt gofmt -s -w
//go:generate rgl version.rl -o zz_generated_version.go
//go:generate rgl constraint.rl -o zz_generated_constraint.go
//go:generate fmt zz_generated_version.go
//go:generate fmt zz_generated_constraint.go
