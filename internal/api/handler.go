// internal/api/handler.go
package api

import (
	"encoding/json"
	"net/http"
	"smoothie/internal/workflow"
)

func RunWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	var wf workflow.Workflow
	if err := json.NewDecoder(r.Body).Decode(&wf); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go wf.Run() // async execution

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Workflow started"))
}
