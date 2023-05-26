package serverFast

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func ServerFast() {

	listenAddr := "127.0.0.1:8083"

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/hello-world":
			fmt.Fprintf(ctx, "Hello, world! I am Fast!")
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	if err := fasthttp.ListenAndServe(listenAddr, requestHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %v", err)
	}
}
