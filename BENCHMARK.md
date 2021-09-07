Compared to [Masterminds/semver](https://github.com/Masterminds/semver):
```
go run golang.org/x/perf/cmd/benchstat -geomean masterminds.txt ours.txt
name                        old time/op    new time/op    delta
Parse/simple-8                 665ns ± 2%      64ns ± 3%  -90.44%  (p=0.000 n=9+9)
Parse/pre-8                    968ns ±15%     180ns ± 1%  -81.45%  (p=0.000 n=9+9)
Parse/build-8                  937ns ± 5%      97ns ± 5%  -89.65%  (p=0.000 n=8+8)
Parse/pre+build-8             1.49µs ± 4%    0.23µs ± 2%  -84.35%  (p=0.000 n=9+8)
ParseConstraint/equal-8       3.52µs ± 2%    0.21µs ±11%  -93.92%  (p=0.000 n=9+10)
ParseConstraint/tilde-8       5.28µs ± 2%    0.21µs ± 2%  -96.03%  (p=0.000 n=9+8)
ParseConstraint/caret-8       5.40µs ± 9%    0.21µs ± 2%  -96.12%  (p=0.000 n=9+9)
ParseConstraint/wildcard-8    3.53µs ± 4%    0.20µs ± 2%  -94.40%  (p=0.000 n=9+9)
ParseConstraint/range-8       10.2µs ± 5%     0.2µs ± 2%  -97.73%  (p=0.000 n=10+8)
ParseConstraint/union-8       11.3µs ± 2%     0.4µs ± 3%  -96.36%  (p=0.000 n=9+9)
[Geo mean]                    2.86µs         0.18µs       -93.55%

name                        old alloc/op   new alloc/op   delta
Parse/simple-8                  402B ± 0%       96B ± 0%  -76.12%  (p=0.000 n=10+10)
Parse/pre-8                     418B ± 0%      148B ± 0%  -64.62%  (p=0.000 n=10+10)
Parse/build-8                   418B ± 0%      101B ± 0%  -75.87%  (p=0.000 n=10+10)
Parse/pre+build-8               467B ± 0%      160B ± 0%  -65.74%  (p=0.000 n=9+10)
ParseConstraint/equal-8       1.30kB ± 0%    0.30kB ± 0%  -76.68%  (p=0.000 n=10+10)
ParseConstraint/tilde-8       1.26kB ± 0%    0.30kB ± 0%  -75.96%  (p=0.000 n=10+10)
ParseConstraint/caret-8       1.26kB ± 0%    0.30kB ± 0%  -75.95%  (p=0.000 n=10+10)
ParseConstraint/wildcard-8    1.29kB ± 0%    0.30kB ± 0%  -76.38%  (p=0.000 n=10+10)
ParseConstraint/range-8       2.34kB ± 0%    0.30kB ± 0%  -87.03%  (p=0.000 n=9+10)
ParseConstraint/union-8       2.50kB ± 0%    0.45kB ± 0%  -82.11%  (p=0.000 n=9+10)
[Geo mean]                      936B           220B       -76.49%

name                        old allocs/op  new allocs/op  delta
Parse/simple-8                  3.00 ± 0%      1.00 ± 0%  -66.67%  (p=0.000 n=10+10)
Parse/pre-8                     4.00 ± 0%      4.00 ± 0%     ~     (all equal)
Parse/build-8                   4.00 ± 0%      2.00 ± 0%  -50.00%  (p=0.000 n=10+10)
Parse/pre+build-8               5.00 ± 0%      5.00 ± 0%     ~     (all equal)
ParseConstraint/equal-8         15.0 ± 0%       4.0 ± 0%  -73.33%  (p=0.000 n=10+10)
ParseConstraint/tilde-8         12.0 ± 0%       4.0 ± 0%  -66.67%  (p=0.000 n=10+10)
ParseConstraint/caret-8         12.0 ± 0%       4.0 ± 0%  -66.67%  (p=0.000 n=10+10)
ParseConstraint/wildcard-8      14.0 ± 0%       4.0 ± 0%  -71.43%  (p=0.000 n=10+10)
ParseConstraint/range-8         22.0 ± 0%       4.0 ± 0%  -81.82%  (p=0.000 n=10+10)
ParseConstraint/union-8         21.0 ± 0%       7.0 ± 0%  -66.67%  (p=0.000 n=10+10)
[Geo mean]                      8.96           3.51       -60.81%
```

Compared to [hashicorp/go-version](https://github.com/hashicorp-go-version) (note that `go-version` does not have feature parity):
```
go run golang.org/x/perf/cmd/benchstat -geomean hashicorp.txt ours.txt
name                     old time/op    new time/op    delta
Parse/simple-8             1.10µs ± 2%    0.06µs ± 7%   -94.09%  (p=0.000 n=10+9)
Parse/pre-8                1.26µs ± 3%    0.18µs ± 4%   -85.57%  (p=0.000 n=10+9)
Parse/build-8              1.35µs ±13%    0.10µs ± 6%   -92.63%  (p=0.000 n=10+9)
Parse/pre+build-8          1.59µs ± 3%    0.24µs ±12%   -84.61%  (p=0.000 n=9+9)
ParseConstraint/equal-8    2.14µs ± 2%    0.22µs ± 9%   -89.94%  (p=0.000 n=9+10)
ParseConstraint/tilde-8    2.64µs ± 6%    0.21µs ± 3%   -91.94%  (p=0.000 n=10+8)
ParseConstraint/range-8    1.28µs ± 4%    0.25µs ±13%   -80.75%  (p=0.000 n=10+10)
[Geo mean]                 1.55µs         0.19µs        -87.67%

name                     old alloc/op   new alloc/op   delta
Parse/simple-8               539B ± 0%       96B ± 0%   -82.19%  (p=0.002 n=8+10)
Parse/pre-8                  539B ± 0%      148B ± 0%   -72.54%  (p=0.002 n=8+10)
Parse/build-8                539B ± 0%      101B ± 0%   -81.26%  (p=0.002 n=8+10)
Parse/pre+build-8            539B ± 0%      160B ± 0%   -70.30%  (p=0.000 n=10+10)
ParseConstraint/equal-8    1.06kB ± 0%    0.30kB ± 0%   -71.19%  (p=0.000 n=10+10)
ParseConstraint/tilde-8    1.05kB ± 0%    0.30kB ± 0%   -70.95%  (p=0.000 n=10+10)
ParseConstraint/range-8      113B ± 1%      304B ± 0%  +169.98%  (p=0.000 n=10+10)
[Geo mean]                   521B           220B        -57.79%

name                     old allocs/op  new allocs/op  delta
Parse/simple-8               5.00 ± 0%      1.00 ± 0%   -80.00%  (p=0.000 n=10+10)
Parse/pre-8                  5.00 ± 0%      4.00 ± 0%   -20.00%  (p=0.000 n=10+10)
Parse/build-8                5.00 ± 0%      2.00 ± 0%   -60.00%  (p=0.000 n=10+10)
Parse/pre+build-8            5.00 ± 0%      5.00 ± 0%      ~     (all equal)
ParseConstraint/equal-8      11.0 ± 0%       4.0 ± 0%   -63.64%  (p=0.000 n=10+10)
ParseConstraint/tilde-8      10.0 ± 0%       4.0 ± 0%   -60.00%  (p=0.000 n=10+10)
ParseConstraint/range-8      5.00 ± 0%      4.00 ± 0%   -20.00%  (p=0.000 n=10+10)
[Geo mean]                   6.18           3.51        -43.13%
```