package mix

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type Config struct {
	MinBatch int
}

func (c Config) Validate() error {
	if c.MinBatch < 2 {
		return errors.New("minimum batch must be at least 2")
	}
	return nil
}

// ShuffleModel applies a secret random permutation and replaces every test
// representation. It is a simulation primitive only: modelReRandomize destroys
// the payload and is not a re-randomizable encryption scheme.
func ShuffleModel(cfg Config, in []TaggedCell) ([]TaggedCell, error) {
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
		out[i].Cell = modelReRandomize(out[i].Cell, nonce)
	}
	return out, nil
}

func uniform(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return int(v.Int64()), nil
}
