// cmd/server/main.go
package main

import (
	"net/http"
	"smoothie/internal/api"
)

func main() {
	http.HandleFunc("/workflows/run", api.RunWorkflowHandler)
	http.ListenAndServe(":8080", nil)
}
