package webapi

import "net/http"

// ReplayStatus API
func ReplayStatus(r *http.Request, w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

// ReceiveTask API, receive task from server
func ReceiveTask(r *http.Request, w http.ResponseWriter) {

}
