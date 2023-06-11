package main

import (
	"testing"
)

func TestCosine(t *testing.T) {
	v1 := []float64{1, 2, 3}
	v2 := []float64{1, 2, 3}
	result := calculateCosineSimilarity(v1, v2)

	if result != 1.0 {
		t.Errorf("Expected cosine similarity of 1.0, but got %f", result)
	}
}
