package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, h *http.Request) {
		w.Write([]byte("hhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
