package main_test

//nolint
import (
	"testing"

	testmain "treuzedev/geheim"
	"treuzedev/geheim/packages/testhelpers"
)

func BenchmarkGeheimEncryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksEncryption()

		b.StartTimer()
		testmain.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimDecryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksDecryption()

		b.StartTimer()
		testmain.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimCheckDecrypted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksCheckDecrypted()

		b.StartTimer()
		testmain.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimCheckEncrypted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksCheckEncrypted()

		b.StartTimer()
		testmain.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}
