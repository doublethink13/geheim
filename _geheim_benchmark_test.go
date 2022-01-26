package geheim_test

//nolint
import (
	"testing"

	"treuzedev/geheim"
	"treuzedev/geheim/packages/testhelpers"
)

func BenchmarkGeheimEncryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksEncryption()

		b.StartTimer()
		geheim.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimDecryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksDecryption()

		b.StartTimer()
		geheim.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimCheckDecrypted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksCheckDecrypted()

		b.StartTimer()
		geheim.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimCheckEncrypted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := testhelpers.SetupBenchmarksCheckEncrypted()

		b.StartTimer()
		geheim.Geheim(benchmarkConfig)
		b.StopTimer()

		testhelpers.CleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}
