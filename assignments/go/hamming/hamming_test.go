package hamming

import (
	"testing"
)

var testCases = []struct {
	expected         int
	strandA, strandB string
	description      string
}{
	{0, "", "", "no difference between empty strands"},
	{2, "AG", "CT", "complete hamming distance for small strands"},
	{0, "A", "A", "no difference between identical strands"},
	{1, "A", "G", "complete distance for single nucleotide strands"},
	{1, "AT", "CT", "small hamming distance"},
	{1, "GGACG", "GGTCG", "small hamming distance in longer strands"},
	{0, "AAAG", "AAA", "ignores extra length on first strand when longer"},
	{0, "AAA", "AAAG", "ignores extra length on second strand when longer"},
	{4, "GATACA", "GCATAA", "large hamming distance"},
	{9, "GGACGGATTCTG", "AGGACGGATTCT", "hamming distance in very long strands"},
}

func TestHamming(t *testing.T) {
	for _, tc := range testCases {

		observed := Distance(tc.strandA, tc.strandB)

		if tc.expected != observed {
			t.Fatalf(`%s:
{%v,%v}
expected: %v
observed: %v`,
				tc.description,
				tc.strandA,
				tc.strandB,
				tc.expected,
				observed,
			)
		}
	}
}
