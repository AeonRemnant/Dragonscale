package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aeonremnant/dragonscale/internal/docker"
)

func StartRESTServer() {
	fmt.Println("Initialized REST API")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		containers, err := docker.ListRunningContainers()
		if err != nil {
			http.Error(w, "Error Fetching Containers", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(containers)
	})

	http.ListenAndServe(":8080", nil)
}
