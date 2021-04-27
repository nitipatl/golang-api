svz@Nitipats-MBP utils % go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/nitipatl/golang-api/mvc/utils
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkBubbleSort10-8         176059380                6.781 ns/op
BenchmarkBubbleSort10000-8        260623              4423 ns/op
BenchmarkBubbleSort50000-8             1        2250343758 ns/op
BenchmarkBubbleSort100000-8            1        9104066140 ns/op
BenchmarkNativeSort10-8         12202651                94.65 ns/op
BenchmarkNativeSort10000-8          2377            464423 ns/op
BenchmarkNativeSort50000-8           441           2667571 ns/op
BenchmarkNativeSort100000-8          208           5850144 ns/op
PASS
ok      github.com/nitipatl/golang-api/mvc/utils        21.905s