package main

import (
	"testing"
	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/testhelpers"
)

func BenchmarkGeheimEncryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := setupBenchmarksEncryption()

		b.StartTimer()
		geheim(benchmarkConfig)
		b.StopTimer()

		cleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimDecryption(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := setupBenchmarksDecryption()

		b.StartTimer()
		geheim(benchmarkConfig)
		b.StopTimer()

		cleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimCheckDecrypted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := setupBenchmarksCheckDecrypted()

		b.StartTimer()
		geheim(benchmarkConfig)
		b.StopTimer()

		cleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func BenchmarkGeheimCheckEncrypted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkConfig := setupBenchmarksCheckEncrypted()

		b.StartTimer()
		geheim(benchmarkConfig)
		b.StopTimer()

		cleanUpBenchmarkTesfiles(benchmarkConfig)
	}
}

func setupBenchmarksEncryption() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     testhelpers.Testfile0D,
		},
		{
			filepath: "file1",
			data:     testhelpers.Testfile1D,
		},
		{
			filepath: "file2",
			data:     testhelpers.Testfile2D,
		},
	}
	filepaths := []string{}

	for _, file := range files {
		testhelpers.GenerateTestFiles(file.data, file.filepath)
		filepaths = append(filepaths, file.filepath)
	}

	return config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   true,
		Decrypt:   false,
		Files:     filepaths,
	}
}

func cleanUpBenchmarkTesfiles(benchmarkConfig config.Config) {
	for _, filepath := range benchmarkConfig.Files {
		testhelpers.RemoveTestFile(filepath)
	}
}

func setupBenchmarksDecryption() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     testhelpers.Testfile0E,
		},
		{
			filepath: "file1",
			data:     testhelpers.Testfile1E,
		},
		{
			filepath: "file2",
			data:     testhelpers.Testfile2E,
		},
	}
	filepaths := []string{}

	for _, file := range files {
		testhelpers.GenerateTestFiles(file.data, file.filepath)
		filepaths = append(filepaths, file.filepath)
	}

	return config.Config{
		Check:     "",
		SecretKey: "imsosecret",
		Encrypt:   false,
		Decrypt:   true,
		Files:     filepaths,
	}
}

func setupBenchmarksCheckDecrypted() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     testhelpers.Testfile0D,
		},
		{
			filepath: "file1",
			data:     testhelpers.Testfile1D,
		},
		{
			filepath: "file2",
			data:     testhelpers.Testfile2D,
		},
	}
	filepaths := []string{}

	for _, file := range files {
		testhelpers.GenerateTestFiles(file.data, file.filepath)
		filepaths = append(filepaths, file.filepath)
	}

	return config.Config{
		Check:     "d",
		SecretKey: "",
		Encrypt:   true,
		Decrypt:   false,
		Files:     filepaths,
	}
}

func setupBenchmarksCheckEncrypted() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     testhelpers.Testfile0E,
		},
		{
			filepath: "file1",
			data:     testhelpers.Testfile1E,
		},
		{
			filepath: "file2",
			data:     testhelpers.Testfile2E,
		},
	}
	filepaths := []string{}

	for _, file := range files {
		testhelpers.GenerateTestFiles(file.data, file.filepath)
		filepaths = append(filepaths, file.filepath)
	}

	return config.Config{
		Check:     "e",
		SecretKey: "",
		Encrypt:   true,
		Decrypt:   false,
		Files:     filepaths,
	}
}
