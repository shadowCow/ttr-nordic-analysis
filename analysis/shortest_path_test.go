package analysis_test

import (
	"testing"

	"github.com/shadowcow/ttrn_analysis/analysis"
)

func TestShortestPath(t *testing.T) {
	// tickets have the shortest path cost between two cities
	// so those are easy test cases
	tickets := analysis.Tickets()

	for _, ticket := range tickets {
		graph := analysis.NewGraph(analysis.Routes())
		shortestPaths := analysis.NewShortestPaths(graph)
		gotCost, gotPath := shortestPaths.ShortestPath(ticket.From, ticket.To)
		wantCost := ticket.Value

		t.Logf("shortestPath %s to %s: %v cost: %d", ticket.From, ticket.To, gotPath, gotCost)
		if gotCost != wantCost {
			t.Errorf("[FAIL] %s to %s: got %d, want %d", ticket.From, ticket.To, gotCost, wantCost)
		}
	}
}
