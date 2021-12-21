package shared

// TODO: error code
// TODO: log exit reason
// TODO: exit gracefully?
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
