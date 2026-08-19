package mix

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
)

type Config struct {
	MinBatch int
}

func (c Config) Validate() error {
	if c.MinBatch < 2 {
		return errors.New("min batch must be >= 2")
	}
	return nil
}

// HonestMix performs a cryptographically seeded Fisher-Yates permutation and
// re-randomizes every output. Ground-truth tags are returned separately only
// to enable statistical tests; they never influence the permutation.
func HonestMix(cfg Config, in []TaggedCell) ([]TaggedCell, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if len(in) < cfg.MinBatch {
		return nil, ErrBatchTooSmall
	}

	out := append([]TaggedCell(nil), in...)
	for i := len(out) - 1; i > 0; i-- {
		j, err := uniform(i + 1)
		if err != nil {
			return nil, err
		}
		out[i], out[j] = out[j], out[i]
	}
	for i := range out {
		nonce, err := freshNonce()
		if err != nil {
			return nil, err
		}
		out[i].Cell = ReRandomize(out[i].Cell, nonce)
	}
	return out, nil
}

func uniform(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}
	// rejection sampling avoids modulo bias
	max := ^uint64(0) - (^uint64(0) % uint64(n))
	for {
		var b [8]byte
		if _, err := rand.Read(b[:]); err != nil {
			return 0, err
		}
		v := binary.BigEndian.Uint64(b[:])
		if v < max {
			return int(v % uint64(n)), nil
		}
	}
}
