# nomad-anytrust-mix-sim

Executable research model for Nomad's **anytrust batch-mixing invariant**.

The simulator models a batch mixer that:

1. refuses batches below a configured anonymity threshold,
2. applies a secret unbiased permutation,
3. replaces every cell representation with a fresh re-randomized representation,
4. exposes no application-selected timing channel.

The intended security assumption is **anytrust**: privacy should survive when at least one critical mix stage behaves honestly. This repository tests statistical consequences of that assumption; it does not claim to implement a formally verified deployable mixnet cryptosystem.

## Build and test

```bash
go test ./...
go test -race ./...
go vet ./...
go run ./cmd/mix-sim -batch 5000 -targets 16 -rounds 10000
```

The Monte Carlo output should converge on the random-guess baseline.

## Security boundaries

- `Tag` is test-only ground truth and never encoded in the simulated cell.
- `ReRandomize` models a property; it is intentionally not presented as production cryptography.
- A real system would need a reviewed re-randomizable encryption or verifiable shuffle construction, robust batch integrity, active-tagging defenses, and formal proofs.
