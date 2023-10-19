package hash

import (
	"testing"
)

func TestHashSolution(t *testing.T) {
	solution, err := HashSolutionBySecretKey("test", "secret")

	if err != nil {
		t.Errorf("Failed to hash solution: %v", err)
	}

	if solution == "" {
		t.Errorf("Failed to hash solution: %v", solution)
	}
}

func TestUnhashSolution(t *testing.T) {
	solution, err := HashSolutionBySecretKey("test", "secret")

	if err != nil {
		t.Errorf("Failed to hash solution: %v", err)
	}

	unhashedSolution, err := UnhashSolution(solution, "secret")

	if err != nil {
		t.Errorf("Failed to unhash solution: %v", err)
	}

	if unhashedSolution != "test" {
		t.Errorf("Failed to unhash solution: %v", unhashedSolution)
	}
}
