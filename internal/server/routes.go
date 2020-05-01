package server

import (
	"fmt"
	"net/http"

	httpctrl "github.com/nydan/glean/internal/controller/http"
)

func router(c *httpctrl.Controller) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	}))
	return router
}
