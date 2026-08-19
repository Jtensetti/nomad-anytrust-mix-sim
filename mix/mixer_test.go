package mix

import (
	"bytes"
	"testing"
)

func TestHonestMixChangesRepresentation(t *testing.T) {
	in := make([]TaggedCell, 256)
	for i := range in {
		c, err := NewTagged(uint64(i))
		if err != nil {
			t.Fatal(err)
		}
		in[i] = c
	}
	out, err := HonestMix(Config{MinBatch: 64}, in)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != len(in) {
		t.Fatal("batch size changed")
	}
	unchanged := 0
	for i := range out {
		if bytes.Equal(out[i].Cell[:], in[i].Cell[:]) {
			unchanged++
		}
	}
	if unchanged != 0 {
		t.Fatalf("%d cells retained representation", unchanged)
	}
}

func TestFailClosedBelowThreshold(t *testing.T) {
	in := make([]TaggedCell, 7)
	if _, err := HonestMix(Config{MinBatch: 8}, in); err != ErrBatchTooSmall {
		t.Fatalf("expected ErrBatchTooSmall, got %v", err)
	}
}

func TestRandomGuessBaseline(t *testing.T) {
	got := RandomGuessRecall(5000, 16)
	want := 16.0 / 5000.0
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}
