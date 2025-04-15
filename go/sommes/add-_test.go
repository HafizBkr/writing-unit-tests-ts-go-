package main

import(
	"testing"
)
//test unitaire classique
func TestAdd(t *testing.T) {
	Obtenue:=add(12,12)
	Attendue:=24

	if Obtenue != Attendue{
		t.Errorf("Obtenue: %v, attendue: %v", Obtenue, Attendue)
	}

}

//test de benchmarking
func BenchmarkAdd(b *testing.B) {

	for i := 0; i < b.N; i++ {
		add(12, 12) 
	}
}

//capturer le résultat du benchmark
// func BenchmarkAddResult(b *testing.B) testing.BenchmarkResult {
// 	return testing.Benchmark(BenchmarkAdd)
// }