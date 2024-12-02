package main

import (
	"fmt"
	"net/http"

	"github.com/rajatjindal/wasi-go-sdk/pkg/wasihttp"
)

func init() {
	wasihttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("x-wasi-rocks", "true")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
		fmt.Fprint(w, "Hello Wasi !!")
	})
}

func main() {}
