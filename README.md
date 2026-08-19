# Nomad verifiable batch-mix prototype

This repository carries payloads through independently randomized, verifiable batch shuffles. It uses Kyber's implementation of Andrew Neff's verifiable shuffle of ElGamal pairs. Nomad code does not implement the shuffle proof itself.

The v0.1 profile encodes each 504-byte clear cell as eighteen ElGamal pairs. A mixed representation occupies 1152 bytes and receives 48 bytes of fresh padding to form one 1200-byte wire cell. All chunks of a cell follow the same secret permutation.

Implemented:

- payload-preserving ElGamal re-randomization;
- non-interactive correctness proofs for every shuffle round;
- exact batch-size preservation and fail-closed proof verification;
- an anytrust committee chain: one honest permutation is sufficient to hide the ordering from the other mixers;
- strict 1200-byte serialization for the constant-rate test profile.

Run go test -race ./..., go vet ./... and go run ./cmd/mix-sim.

## Security boundary

This is a research integration of an established construction, not an audited Nomad mixnet. v0.1 uses one test decryption key. Distributed key generation, threshold decryption, mixer identity signatures, replay handling, network-level delay/drop accountability and independent review are release gates for production use. Kyber also requires an application-specific security review before security-critical deployment.
