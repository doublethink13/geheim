package testhelpers

import "treuzedev/geheim/packages/config"

func CleanUpBenchmarkTesfiles(benchmarkConfig config.Config) {
	for _, filepath := range benchmarkConfig.Files {
		RemoveTestFile(filepath)
	}
}

func SetupBenchmarksEncryption() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     GetTestfile0D(),
		},
		{
			filepath: "file1",
			data:     GetTestfile1D(),
		},
		{
			filepath: "file2",
			data:     GetTestfile2D(),
		},
	}
	filepaths := []string{}

	for _, file := range files {
		GenerateTestFiles(file.data, file.filepath)
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

func SetupBenchmarksDecryption() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     GetTestfile0E(),
		},
		{
			filepath: "file1",
			data:     GetTestfile1E(),
		},
		{
			filepath: "file2",
			data:     GetTestfile2E(),
		},
	}
	filepaths := []string{}

	for _, file := range files {
		GenerateTestFiles(file.data, file.filepath)
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

func SetupBenchmarksCheckDecrypted() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     GetTestfile0D(),
		},
		{
			filepath: "file1",
			data:     GetTestfile1D(),
		},
		{
			filepath: "file2",
			data:     GetTestfile2D(),
		},
	}
	filepaths := []string{}

	for _, file := range files {
		GenerateTestFiles(file.data, file.filepath)
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

func SetupBenchmarksCheckEncrypted() (benchmarkConfig config.Config) {
	files := []struct {
		filepath string
		data     string
	}{
		{
			filepath: "file0",
			data:     GetTestfile0E(),
		},
		{
			filepath: "file1",
			data:     GetTestfile1E(),
		},
		{
			filepath: "file2",
			data:     GetTestfile2E(),
		},
	}
	filepaths := []string{}

	for _, file := range files {
		GenerateTestFiles(file.data, file.filepath)
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
