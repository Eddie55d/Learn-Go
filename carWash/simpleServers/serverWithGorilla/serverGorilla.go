package serverGorilla

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerGorilla() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World! I am Gorilla!\n")
	}

	r := mux.NewRouter()
	r.HandleFunc("/hello-world", helloHandler)

	log.Fatal(http.ListenAndServe(":8081", r))
}
