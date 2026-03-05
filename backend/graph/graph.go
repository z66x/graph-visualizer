package graph

type Edge struct {
	To     string  `json:"to"`
	Weight float64 `json:"weight"`
}

type Graph struct {
	Nodes    []string          `json:"nodes"`
	Edges    map[string][]Edge `json:"edges"`
	Directed bool              `json:"directed"`
}

type Step struct {
	Visited   []string          `json:"visited"`
	Current   string            `json:"current"`
	Frontier  []string          `json:"frontier"`
	EdgeTaken map[string]string `json:"edgeTaken"`
	Message   string            `json:"message"`
}
