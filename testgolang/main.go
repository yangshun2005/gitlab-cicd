package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello chinaase willim ,I running in docker-container and buit by gitlab-ci"))
	})
	http.ListenAndServe(":8001", nil)
}
