package main

import "testing"

func TestFrequencySort(t *testing.T) {
	// Case

	s := "abbcccddddeeeee"
	w := "eeeeeddddcccbba"

	res := FrequencySort(s)

	assert(t, s, w, res)

	// Case

	s = "tree"
	w = "eert"

	res = FrequencySort(s)

	assert(t, s, w, res)

	// Case

	s = "cccaaa"
	w = "aaaccc"

	res = FrequencySort(s)

	assert(t, s, w, res)

	// Case

	s = "Aabb"
	w = "bbaA"

	res = FrequencySort(s)

	assert(t, s, w, res)
}

func assert(t *testing.T, s string, w string, r string) {
	if r != w {
		t.Fatalf("expected %v, got %v", w, r)
	}
}
