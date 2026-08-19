package main

import (
	"flag"
	"fmt"
	"math/rand/v2"

	"github.com/Jtensetti/nomad-anytrust-mix-sim/mix"
)

func main() {
	batch := flag.Int("batch", 5000, "cells per batch")
	targets := flag.Int("targets", 16, "target cells")
	rounds := flag.Int("rounds", 10000, "Monte Carlo rounds")
	flag.Parse()
	if *targets > *batch || *targets <= 0 {
		panic("invalid target count")
	}

	hits := 0
	for r := 0; r < *rounds; r++ {
		// After a perfect secret permutation, output positions are exchangeable.
		// Simulate an adversary that chooses targetCount candidate positions.
		truth := rand.Perm(*batch)[:*targets]
		guess := rand.Perm(*batch)[:*targets]
		set := make(map[int]struct{}, *targets)
		for _, x := range truth {
			set[x] = struct{}{}
		}
		for _, x := range guess {
			if _, ok := set[x]; ok {
				hits++
			}
		}
	}
	recall := float64(hits) / float64(*rounds**targets)
	fmt.Printf("batch=%d targets=%d rounds=%d\n", *batch, *targets, *rounds)
	fmt.Printf("observed random-guess recall: %.6f\n", recall)
	fmt.Printf("expected baseline:           %.6f\n", mix.RandomGuessRecall(*batch, *targets))
}
