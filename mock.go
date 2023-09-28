package pulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type MockULIDGenerator struct {
	entropy map[string]*rand.Rand
}

// newULID returns a new ULID for a fixed date using a mock entropy source
// with a fixed seed. A separate predicatable entropy source is used for
// each unique prefix to keep IDs as fixed as possible.
func (g *MockULIDGenerator) newULID(prefix string, t time.Time) ulid.ULID {
	if g.entropy == nil {
		g.entropy = make(map[string]*rand.Rand)
	}

	entropy, ok := g.entropy[prefix]
	if !ok {
		entropy = rand.New(rand.NewSource(0))
		g.entropy[prefix] = entropy
	}

	// Overwrite input time with fixed time.
	t = time.Date(2021, time.August, 16, 15, 34, 0, 0, time.UTC)

	return ulid.MustNew(ulid.Timestamp(t), entropy)
}
