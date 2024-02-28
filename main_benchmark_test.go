// main_benchmark.go
package main

import (
	"testing"
)

// first version
func BenchmarkGetNthPrime1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNthPrime1(10000)
	}
}

// first optmization
func BenchmarkGetNthPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNthPrime2(10000)
	}
}

// second optmization
func BenchmarkGetNthPrime3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNthPrime3(10000)
	}
}

//last optmization - in this case I didn't get any performance improvement, quite the opposite.
//With more time and a better knowledge of the language's features, I'll be able to make this kind of improvement

func BenchmarkGetNthPrime4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getNthPrime4(10000)
	}
}
