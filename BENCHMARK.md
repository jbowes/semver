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