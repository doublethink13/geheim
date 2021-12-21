package shared

func CompareStringSlices(a, b []string) bool {
	switch {
	case len(a) != len(b):
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CompareByteSlices(a, b []byte) bool {
	switch {
	case len(a) != len(b):
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
