Compared to [Masterminds/semver](https://github.com/Masterminds/semver):
```
go run golang.org/x/perf/cmd/benchstat -geomean masterminds.txt ours.txt
name                        old time/op    new time/op    delta
Parse/simple-8                 604ns ± 3%      61ns ± 6%  -89.92%  (p=0.000 n=10+9)
Parse/pre-8                    859ns ± 2%     140ns ± 2%  -83.74%  (p=0.000 n=9+10)
Parse/build-8                  851ns ± 2%      96ns ± 3%  -88.66%  (p=0.000 n=10+10)
Parse/pre+build-8             1.37µs ± 5%    0.19µs ± 2%  -85.81%  (p=0.000 n=9+9)
ParseConstraint/equal-8       3.42µs ± 2%    0.21µs ± 3%  -93.95%  (p=0.000 n=9+10)
ParseConstraint/tilde-8       5.19µs ± 4%    0.21µs ± 2%  -96.01%  (p=0.000 n=9+9)
ParseConstraint/caret-8       5.14µs ± 2%    0.21µs ± 2%  -95.96%  (p=0.000 n=8+10)
ParseConstraint/wildcard-8    3.41µs ± 2%    0.20µs ± 2%  -94.21%  (p=0.000 n=9+10)
ParseConstraint/range-8       9.95µs ± 2%    0.22µs ± 1%  -97.78%  (p=0.000 n=10+10)
ParseConstraint/union-8       11.0µs ± 1%     0.4µs ± 1%  -96.36%  (p=0.000 n=8+8)
[Geo mean]                    2.70µs         0.17µs       -93.57%

name                        old alloc/op   new alloc/op   delta
Parse/simple-8                  402B ± 0%       96B ± 0%  -76.12%  (p=0.000 n=10+10)
Parse/pre-8                     418B ± 0%      148B ± 0%  -64.62%  (p=0.000 n=10+10)
Parse/build-8                   418B ± 0%      101B ± 0%  -75.85%  (p=0.000 n=10+10)
Parse/pre+build-8               467B ± 0%      160B ± 0%  -65.74%  (p=0.002 n=8+10)
ParseConstraint/equal-8       1.30kB ± 0%    0.30kB ± 0%  -76.68%  (p=0.000 n=10+10)
ParseConstraint/tilde-8       1.26kB ± 0%    0.30kB ± 0%  -75.96%  (p=0.000 n=10+10)
ParseConstraint/caret-8       1.26kB ± 0%    0.30kB ± 0%  -75.94%  (p=0.000 n=10+10)
ParseConstraint/wildcard-8    1.29kB ± 0%    0.30kB ± 0%  -76.38%  (p=0.000 n=10+10)
ParseConstraint/range-8       2.34kB ± 0%    0.30kB ± 0%  -87.02%  (p=0.000 n=10+10)
ParseConstraint/union-8       2.50kB ± 0%    0.45kB ± 0%  -82.11%  (p=0.000 n=6+10)
[Geo mean]                      936B           220B       -76.48%

name                        old allocs/op  new allocs/op  delta
Parse/simple-8                  3.00 ± 0%      1.00 ± 0%  -66.67%  (p=0.000 n=10+10)
Parse/pre-8                     4.00 ± 0%      3.00 ± 0%  -25.00%  (p=0.000 n=10+10)
Parse/build-8                   4.00 ± 0%      2.00 ± 0%  -50.00%  (p=0.000 n=10+10)
Parse/pre+build-8               5.00 ± 0%      4.00 ± 0%  -20.00%  (p=0.000 n=10+10)
ParseConstraint/equal-8         15.0 ± 0%       4.0 ± 0%  -73.33%  (p=0.000 n=10+10)
ParseConstraint/tilde-8         12.0 ± 0%       4.0 ± 0%  -66.67%  (p=0.000 n=10+10)
ParseConstraint/caret-8         12.0 ± 0%       4.0 ± 0%  -66.67%  (p=0.000 n=10+10)
ParseConstraint/wildcard-8      14.0 ± 0%       4.0 ± 0%  -71.43%  (p=0.000 n=10+10)
ParseConstraint/range-8         22.0 ± 0%       4.0 ± 0%  -81.82%  (p=0.000 n=10+10)
ParseConstraint/union-8         21.0 ± 0%       7.0 ± 0%  -66.67%  (p=0.000 n=10+10)
[Geo mean]                      8.96           3.34       -62.76%
```

Compared to [hashicorp/go-version](https://github.com/hashicorp-go-version) (note that `go-version` does not have feature parity):
```
go run golang.org/x/perf/cmd/benchstat -geomean hashicorp.txt ours.txt
name                     old time/op    new time/op    delta
Parse/simple-8             1.10µs ± 1%    0.06µs ± 6%   -94.47%  (p=0.000 n=8+9)
Parse/pre-8                1.26µs ± 3%    0.14µs ± 2%   -88.90%  (p=0.000 n=9+10)
Parse/build-8              1.32µs ± 6%    0.10µs ± 3%   -92.71%  (p=0.000 n=10+10)
Parse/pre+build-8          1.59µs ± 3%    0.19µs ± 2%   -87.79%  (p=0.000 n=9+9)
ParseConstraint/equal-8    2.03µs ± 4%    0.21µs ± 3%   -89.78%  (p=0.000 n=8+10)
ParseConstraint/tilde-8    2.48µs ± 6%    0.21µs ± 2%   -91.66%  (p=0.000 n=9+9)
ParseConstraint/range-8    1.26µs ± 8%    0.22µs ± 1%   -82.39%  (p=0.000 n=9+10)
[Geo mean]                 1.52µs         0.17µs        -88.55%

name                     old alloc/op   new alloc/op   delta
Parse/simple-8               539B ± 0%       96B ± 0%   -82.19%  (p=0.000 n=10+10)
Parse/pre-8                  539B ± 0%      148B ± 0%   -72.54%  (p=0.000 n=9+10)
Parse/build-8                539B ± 0%      101B ± 0%   -81.25%  (p=0.000 n=10+10)
Parse/pre+build-8            539B ± 0%      160B ± 0%   -70.32%  (p=0.002 n=8+10)
ParseConstraint/equal-8    1.05kB ± 0%    0.30kB ± 0%   -71.18%  (p=0.000 n=10+10)
ParseConstraint/tilde-8    1.05kB ± 0%    0.30kB ± 0%   -70.95%  (p=0.000 n=10+10)
ParseConstraint/range-8      112B ± 0%      304B ± 0%  +171.43%  (p=0.000 n=8+10)
[Geo mean]                   521B           220B        -57.76%

name                     old allocs/op  new allocs/op  delta
Parse/simple-8               5.00 ± 0%      1.00 ± 0%   -80.00%  (p=0.000 n=10+10)
Parse/pre-8                  5.00 ± 0%      3.00 ± 0%   -40.00%  (p=0.000 n=10+10)
Parse/build-8                5.00 ± 0%      2.00 ± 0%   -60.00%  (p=0.000 n=10+10)
Parse/pre+build-8            5.00 ± 0%      4.00 ± 0%   -20.00%  (p=0.000 n=10+10)
ParseConstraint/equal-8      11.0 ± 0%       4.0 ± 0%   -63.64%  (p=0.000 n=10+10)
ParseConstraint/tilde-8      10.0 ± 0%       4.0 ± 0%   -60.00%  (p=0.000 n=10+10)
ParseConstraint/range-8      5.00 ± 0%      4.00 ± 0%   -20.00%  (p=0.000 n=10+10)
[Geo mean]                   6.18           3.34        -45.97%
```