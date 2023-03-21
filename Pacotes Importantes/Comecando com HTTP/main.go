package main

import "net/http"

func main() {
	http.HandleFunc("/", SearchCEP)
	http.ListenAndServe(":8080", nil)
}

func SearchCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Teste"))
}
