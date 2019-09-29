package testpkg1

import "testing"

// BenchmarkGetPrimes test for GetPrimes
func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}
