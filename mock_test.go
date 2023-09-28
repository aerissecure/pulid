package pulid

import (
	"strings"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

func Test_mockULIDGenerator_newULID(t *testing.T) {
	type args struct {
		prefix string
		t      time.Time
	}
	tests := []struct {
		name string
		args args
		want ulid.ULID
	}{
		{
			name: "AA#1",
			args: args{prefix: "AA", t: time.Now()},
			want: ulid.MustParseStrict("01FD7SJ7J006AFVGQT5ZYC0GEK"),
		},
		{
			name: "AA#2",
			args: args{prefix: "AA", t: time.Now()},
			want: ulid.MustParseStrict("01FD7SJ7J0ZW908PVKS1Q4ZYAZ"),
		},
		{
			name: "AA#3",
			args: args{prefix: "AA", t: time.Now()},
			want: ulid.MustParseStrict("01FD7SJ7J0YSHABVQ85AYZ8JHD"),
		},
		{
			name: "BB#1",
			args: args{prefix: "BB", t: time.Now()},
			want: ulid.MustParseStrict("01FD7SJ7J006AFVGQT5ZYC0GEK"),
		},
		{
			name: "BB#1",
			args: args{prefix: "BB", t: time.Now()},
			want: ulid.MustParseStrict("01FD7SJ7J0ZW908PVKS1Q4ZYAZ"),
		},
		{
			name: "BB#3",
			args: args{prefix: "BB", t: time.Now()},
			want: ulid.MustParseStrict("01FD7SJ7J0YSHABVQ85AYZ8JHD"),
		},
	}
	g := &MockULIDGenerator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := g.newULID(tt.args.prefix, tt.args.t); got.Compare(tt.want) != 0 {
				t.Errorf("MockULIDGenerator.newULID() = %v, want %v, %v", got, tt.want, got.Compare(tt.want))
			}
		})
	}

	g = &MockULIDGenerator{}
	SetULIDGenerator(g)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pu := MustNew(tt.args.prefix)
			got := ulid.MustParseStrict(strings.Split(pu.String(), ":")[1])

			if got.Compare(tt.want) != 0 {
				t.Errorf("MustNew() = %v, want %v, %v", got, tt.want, got.Compare(tt.want))
			}
		})
	}
}
