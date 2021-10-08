package presenter

import "net/http"

func MountResponse(
	w http.ResponseWriter,
	status int,
	response []byte,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
