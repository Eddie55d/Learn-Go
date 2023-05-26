package serverGo

import (
	"io"
	"log"
	"net/http"
)

func ServerGO() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World! I am Go\n")
	}

	http.HandleFunc("/hello-world", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
