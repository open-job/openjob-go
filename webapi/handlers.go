package webapi

import "net/http"

// ReplayStatus API
func ReplayStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

// ReceiveTask API, receive task from server
func ReceiveTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
