package main

import (
	"fmt"
	"net/http"

	"github.com/rajatjindal/wasi-go-sdk/pkg/wasihttp"
)

func init() {
	wasihttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello Wasi !!")
	})
}

func main() {}
