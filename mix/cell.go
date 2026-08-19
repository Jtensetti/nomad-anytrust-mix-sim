package mix

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"
)

const CellSize = 1200

type Cell [CellSize]byte

func RandomCell() (Cell, error) {
	var c Cell
	_, err := io.ReadFull(rand.Reader, c[:])
	return c, err
}

// TaggedCell exists only for adversarial simulation. Tag is ground truth and
// is never included in the mixed wire representation.
type TaggedCell struct {
	Tag  uint64
	Cell Cell
}

func NewTagged(tag uint64) (TaggedCell, error) {
	c, err := RandomCell()
	return TaggedCell{Tag: tag, Cell: c}, err
}

// modelReRandomize transforms a cell into a fresh pseudorandom-looking cell.
// It deliberately destroys the payload and exists only to model representation
// replacement across a shuffle boundary.
func modelReRandomize(in Cell, nonce [32]byte) Cell {
	var out Cell
	for block, off := uint64(0), 0; off < len(out); block++ {
		h := sha256.New()
		_, _ = h.Write([]byte("nomad-mix-rerandomize-v1"))
		_, _ = h.Write(nonce[:])
		_, _ = h.Write(in[:])
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], block)
		_, _ = h.Write(b[:])
		sum := h.Sum(nil)
		off += copy(out[off:], sum)
	}
	return out
}

func freshNonce() ([32]byte, error) {
	var n [32]byte
	_, err := io.ReadFull(rand.Reader, n[:])
	return n, err
}

var ErrBatchTooSmall = errors.New("batch smaller than configured minimum")
