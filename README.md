# nomad-anytrust-mix-sim

A simulation of one batch-shuffle property needed by the Nomad experiments.

Despite the repository name, this code does **not** implement an anytrust committee protocol or a mixnet cryptosystem. It models a single honest shuffle stage so surrounding code can test batch thresholds and ground-truth bookkeeping without inventing cryptography.

`ShuffleModel`:

1. rejects batches below a configured minimum,
2. applies a CSPRNG-driven Fisher–Yates permutation,
3. replaces every test representation with a fresh model representation.

`TaggedCell.Tag` exists only as test-harness ground truth. It is not encoded in `Cell` and must never be treated as a wire field.

## What is meaningful here

- configured minimum-batch behavior,
- permutation preserves the cohort exactly once,
- the model replacement changes each tagged representation,
- a random-position adversary lands at the combinatorial baseline.

## What is deliberately absent

`modelReRandomize` destroys payloads. There is no payload-preserving re-randomizable encryption, verifiable shuffle, threshold key protocol, malicious-mixer accountability, active-tagging defense or availability protocol. A real mix layer must replace this repository's model boundary with a reviewed construction.

```bash
go test -race ./...
go vet ./...
go run ./cmd/mix-sim -batch 1000 -targets 16 -rounds 100
```
