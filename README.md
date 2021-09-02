<!--
  Attractive html formatting for rendering in github. sorry text editor
  readers! Besides the header and section links, everything should be clean and
  readable.
-->
<h1 align="center">semver</h1>
<p align="center"><i>🦔 semver and constraint parsing with a focus on performance</i></p>

<div align="center">
  <a href="https://pkg.go.dev/github.com/jbowes/semver"><img src="https://pkg.go.dev/badge/github.com/jbowes/semver.svg" alt="Go Reference"></a>
  <a href="https://github.com/jbowes/semver/releases/latest"><img alt="GitHub tag" src="https://img.shields.io/github/tag/jbowes/semver.svg"></a>
  <a href="https://github.com/jbowes/semver/actions/workflows/go.yml"><img alt="Build Status" src="https://github.com/jbowes/semver/actions/workflows/go.yml/badge.svg?branch=main"></a>
  <a href="./LICENSE"><img alt="BSD license" src="https://img.shields.io/badge/license-BSD-blue.svg"></a>
  <a href="https://codecov.io/gh/jbowes/semver"><img alt="codecov" src="https://img.shields.io/codecov/c/github/jbowes/semver.svg"></a>
  <a href="https://goreportcard.com/report/github.com/jbowes/semver"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/jbowes/semver"></a>
</div><br /><br />

---

[`semver`][godoc] provides semantic version and constraint parsing, comparison,
and testing.

There are many semver packages. This one aims to be the fastest at parsing
and comparing values, with low memory usage. On average, this package is over
ten times faster at parsing versions and constraints than the popular
`Masterminds/semver` package. View more stats in the [benchmarks][bench].

Versions can be compared with one another to determine which is newer.
Constraints specify inclusions and exclusions of semver ranges that a given
version must satisfy. Typically, constraints are used when expressing a
dependency.

## Quick start

```go
import (
  "log"

  "github.com/jbowes/semver"
)

func main() {
  // Parse a version. Two versions can also be Compare()ed
  ver, err := semver.Parse("1.0.2")
  if err != nil {
    log.Fatal("invalid semver")
  }

  // Parse a Constraint, typically used to express dependency
  constr, err := semver.ParseConstraint(">=1.0.0")
  if err != nil {
    log.Fatalln("invalid constraint")
  }

  // Check if a Version satisfies a Constraint
  if constr.Check(ver) {
    log.Printf("%s satisfies constraint %s\n", ver, constr)
  } else {
    log.Printf("%s does not satisfy constraint %s\n", ver, constr)
  }
}
```

For more usage and examples, see the [GoDoc Reference][godoc]

## Contributing

I would love your help!

`semver` is still a work in progress. You can help by:

- Opening a pull request to resolve an [open issue][issues].
- Adding a feature or enhancement of your own! If it might be big, please
  [open an issue][enhancement] first so we can discuss it.
- Improving this `README` or adding other documentation to `semver`.
- Letting [me] know if you're using `semver`.

## Links

- [Semantic Versioning 2.0.0][semver]
- [npm semver calculator][calc]
- [Benchmarks][bench]

[semver]: https://semver.org
[calc]: https://semver.npmjs.com/
[bench]: https://github.com/jbowes/semver/blob/main/BENCHMARK.md

[godoc]: https://pkg.go.dev/github.com/jbowes/semver

[issues]: ./issues
[bug]: ./issues/new?labels=bug
[enhancement]: ./issues/new?labels=enhancement

[me]: https://twitter.com/jrbowes
