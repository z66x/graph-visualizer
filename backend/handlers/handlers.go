package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/z66x/graph-visualizer/graph"
)

type AlgoRequest struct {
	Graph graph.Graph `json:"graph"`
	Start string      `json:"start"`
}

type AlgoResponse struct {
	Steps []graph.Step `json:"steps"`
}

func BFS(w http.ResponseWriter, r *http.Request) {
	req, err := decode(r)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	steps := graph.BFSSteps(req.Graph, req.Start)
	encode(w, AlgoResponse{Steps: steps})
}

func DFS(w http.ResponseWriter, r *http.Request) {
	req, err := decode(r)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	steps := graph.DFSSteps(req.Graph, req.Start)
	encode(w, AlgoResponse{Steps: steps})
}

func Dijkstra(w http.ResponseWriter, r *http.Request) {
    req, err := decode(r)
    if err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }
    found := false
    for _, n := range req.Graph.Nodes {
        if n == req.Start { found = true; break }
    }
    if !found {
        http.Error(w, "start node not in graph", http.StatusBadRequest)
        return
    }
    steps := graph.DijkstraSteps(req.Graph, req.Start)
    encode(w, AlgoResponse{Steps: steps})
}

func decode(r *http.Request) (AlgoRequest, error) {
	var req AlgoRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func encode(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
