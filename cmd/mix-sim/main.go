package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/Jtensetti/nomad-anytrust-mix-sim/mix"
)

func main() {
	batch := flag.Int("batch", 1000, "cells per batch")
	targets := flag.Int("targets", 16, "target cells")
	rounds := flag.Int("rounds", 100, "Monte Carlo rounds")
	flag.Parse()
	if *batch < 2 || *targets <= 0 || *targets > *batch || *rounds <= 0 {
		log.Fatal("invalid parameters")
	}

	hits := 0
	for round := 0; round < *rounds; round++ {
		input := make([]mix.TaggedCell, *batch)
		for i := range input {
			input[i].Tag = uint64(i)
		}
		output, err := mix.ShuffleModel(mix.Config{MinBatch: 2}, input)
		if err != nil {
			log.Fatal(err)
		}

		truth := make(map[int]struct{}, *targets)
		for position, cell := range output {
			if int(cell.Tag) < *targets {
				truth[position] = struct{}{}
			}
		}
		for _, guess := range rand.Perm(*batch)[:*targets] {
			if _, ok := truth[guess]; ok {
				hits++
			}
		}
	}

	recall := float64(hits) / float64(*rounds**targets)
	fmt.Printf("batch=%d targets=%d rounds=%d\n", *batch, *targets, *rounds)
	fmt.Printf("random-position adversary recall after ShuffleModel: %.6f\n", recall)
	fmt.Printf("expected random baseline:                        %.6f\n", mix.RandomGuessRecall(*batch, *targets))
	fmt.Println("note: tags are test-harness ground truth and are not part of the modeled wire cell")
}
