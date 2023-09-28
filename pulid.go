// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pulid implements the pulid type.
// A pulid is an identifier that is a two-byte prefixed ULIDs, with the first two bytes encoding the type of the entity.
package pulid

import (
	"crypto/rand"
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

func init() {
	// Seed the default entropy source.
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

// ID implements a PULID - a prefixed ULID.
type ID string

func (id ID) Parts() (prefix, ulid string) {
	s := string(id)
	index := strings.Index(s, ":")
	if index == -1 {
		return "", s
	}
	return s[:index], s[index+1:]
}

func (id ID) Compare(other ID) int {
	return strings.Compare(string(id), string(other))
}

type ULIDGenerator interface {
	newULID(string, time.Time) ulid.ULID
}

type ulidGenerator struct{}

func (ulidGenerator) newULID(prefix string, t time.Time) ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(t), defaultEntropySource)
}

// The default entropy source.
var defaultEntropySource *ulid.MonotonicEntropy
var defaultULIDGenerator ULIDGenerator = ulidGenerator{}

func SetULIDGenerator(g ULIDGenerator) {
	defaultULIDGenerator = g
}

// MustNew returns a new PULID for time.Now() with the defaultULIDGenerator, given
// given the prefix.
func MustNew(prefix string) ID {
	return ID(prefix + ":" + defaultULIDGenerator.newULID(prefix, time.Now()).String())
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}

// Scan implements the Scanner interface.
func (u *ID) Scan(src interface{}) error {
	if src == nil {
		// Permit empty pulids.
		return nil
	}

	switch v := src.(type) {
	case ID:
		*u = v
	case string:
		*u = ID(v)
	default:
		return fmt.Errorf("pulid: scan error, unexpected type %T", src)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}

func (u ID) String() string {
	return string(u)
}
