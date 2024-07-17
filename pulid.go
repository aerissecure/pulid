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
	"errors"
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
func (id *ID) UnmarshalGQL(v interface{}) error {
	return id.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (id ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(id)))
}

// Scan implements the Scanner interface.
func (id *ID) Scan(src interface{}) error {
	if src == nil {
		// Permit empty pulids.
		return nil
	}

	switch v := src.(type) {
	case ID:
		*id = v
	case string:
		*id = ID(v)
	default:
		return fmt.Errorf("pulid: scan error, unexpected type %T", src)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (id ID) Value() (driver.Value, error) {
	return string(id), nil
}

func (id ID) String() string {
	return string(id)
}

// Parse parses an ID, validating the underlying ULID and returning
// an error in case of failure.
//
// ulid.ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings produce undefined ULIDs. For a version that
// returns an error instead, see ParseStrict.
func Parse(id string) (prefix string, u ulid.ULID, err error) {
	var ul string
	prefix, ul = ID(id).Parts()

	if len(prefix) == 0 {
		return prefix, u, errors.New("pulid: id has no prefix")
	}

	u, err = ulid.Parse(ul)
	return prefix, u, err
}

// ParseStrict parses an ID, validating the underlying ULID and returning
// an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists
// only of valid base32 characters. It is slightly slower than Parse.
//
// ulid.ErrDataSize is returned if the len(ulid) is different from an encoded
// ULID's length. Invalid encodings return ErrInvalidCharacters.
func ParseStrict(id string) (prefix string, u ulid.ULID, err error) {
	var ul string
	prefix, ul = ID(id).Parts()

	if len(prefix) == 0 {
		return prefix, u, errors.New("publid: id has no prefix")
	}

	u, err = ulid.ParseStrict(ul)
	return prefix, u, err
}
