package shared

// TODO: error code
// TODO: log exit reason
// TODO: exit gracefully?
func CheckError(e error) {
	if e != nil {
		switch e.Error() {
		case "encoding/hex: invalid byte: U+0000":
			return
		default:
			panic(e)
		}
	}
}
