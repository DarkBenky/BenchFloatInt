package benchmarktest

import (
	"testing"
)

// Integer operations benchmarks
func BenchmarkIntAdd(b *testing.B) {
	var result int32
	for i := 0; i < b.N; i++ {
		result = 100 + 200
	}
	_ = result
}

func BenchmarkIntSub(b *testing.B) {
	var result int32
	for i := 0; i < b.N; i++ {
		result = 300 - 100
	}
	_ = result
}

func BenchmarkIntMul(b *testing.B) {
	var result int32
	for i := 0; i < b.N; i++ {
		result = 20 * 30
	}
	_ = result
}

func BenchmarkIntDiv(b *testing.B) {
	var result int32
	for i := 0; i < b.N; i++ {
		result = 400 / 10
	}
	_ = result
}

// Float operations benchmarks
func BenchmarkFloatAdd(b *testing.B) {
	var result float32
	for i := 0; i < b.N; i++ {
		result = 100.5131 + 200.3764
	}
	_ = result
}

func BenchmarkFloatSub(b *testing.B) {
	var result float32
	for i := 0; i < b.N; i++ {
		result = 300.4131 - 100.2131
	}
	_ = result
}

func BenchmarkFloatMul(b *testing.B) {
	var result float32
	for i := 0; i < b.N; i++ {
		result = 20.6164 * 30.7133
	}
	_ = result
}

func BenchmarkFloatDiv(b *testing.B) {
	var result float32
	for i := 0; i < b.N; i++ {
		result = 400.8454 / 10.1140
	}
	_ = result
}
