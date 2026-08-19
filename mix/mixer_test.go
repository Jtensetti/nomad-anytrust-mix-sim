package mix

import (
	"bytes"
	"testing"
)

func TestShuffleModelPreservesTagsAndChangesEachTaggedRepresentation(t *testing.T) {
	in := make([]TaggedCell, 256)
	original := make(map[uint64]Cell, len(in))
	for i := range in {
		cell, err := NewTagged(uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		in[i] = cell
		original[cell.Tag] = cell.Cell
	}
	out, err := ShuffleModel(Config{MinBatch: 64}, in)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != len(in) {
		t.Fatalf("batch size changed from %d to %d", len(in), len(out))
	}
	seen := make(map[uint64]bool, len(out))
	for _, mixed := range out {
		if seen[mixed.Tag] {
			t.Fatalf("duplicate tag %d", mixed.Tag)
		}
		seen[mixed.Tag] = true
		before, ok := original[mixed.Tag]
		if !ok {
			t.Fatalf("unknown tag %d", mixed.Tag)
		}
		if bytes.Equal(before[:], mixed.Cell[:]) {
			t.Fatalf("tag %d retained its original representation", mixed.Tag)
		}
	}
	if len(seen) != len(in) {
		t.Fatalf("saw %d unique tags, want %d", len(seen), len(in))
	}
}

func TestFailClosedBelowThreshold(t *testing.T) {
	in := make([]TaggedCell, 7)
	if _, err := ShuffleModel(Config{MinBatch: 8}, in); err != ErrBatchTooSmall {
		t.Fatalf("expected ErrBatchTooSmall, got %v", err)
	}
}

func TestRandomGuessBaselineFormula(t *testing.T) {
	got := RandomGuessRecall(5000, 16)
	want := 16.0 / 5000.0
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
