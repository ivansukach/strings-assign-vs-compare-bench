## BENCH STRING ASSIGNMENT AND COMPARISON

goos: windows
goarch: amd64
pkg: strings-assign-vs-compare-bench
cpu: Intel(R) Core(TM) i3-10110U CPU @ 2.10GHz
BenchmarkAssign
BenchmarkAssign-4                                       148063988
8.344 ns/op
BenchmarkCompare
BenchmarkCompare-4                                      147443259
8.912 ns/op
BenchmarkAssignTwoIdentical
BenchmarkAssignTwoIdentical-4                           157971540
7.234 ns/op
BenchmarkCompareTwoIdentical
BenchmarkCompareTwoIdentical-4                          165897202
7.664 ns/op
BenchmarkAssignTwoIdenticalWithoutBooleanFlag
BenchmarkAssignTwoIdenticalWithoutBooleanFlag-4         166133677
8.156 ns/op
BenchmarkCompareTwoIdenticalWithoutBooleanFlag
BenchmarkCompareTwoIdenticalWithoutBooleanFlag-4        155037798
8.289 ns/op
BenchmarkAssignLong
BenchmarkAssignLong-4                                   144777320
7.309 ns/op
BenchmarkCompareLong
BenchmarkCompareLong-4                                  144897103
7.795 ns/op
BenchmarkAssignLongTwoIdentical
BenchmarkAssignLongTwoIdentical-4                       158898662
7.260 ns/op
BenchmarkCompareLongTwoIdentical
BenchmarkCompareLongTwoIdentical-4                      168596250
7.021 ns/op
PASS
