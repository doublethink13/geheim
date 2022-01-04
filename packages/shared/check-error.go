package shared

import "treuzedev/geheim/packages/logging"

// TODO: error code
// TODO: log exit reason
// TODO: exit gracefully?
func CheckError(e error) {
	if e != nil {
		logging.Log(logging.Error, e.Error())
	}
}
