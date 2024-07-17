package pulid

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		ulid    string
		wantErr bool
	}{
		{
			name:    "AA",
			prefix:  "AA",
			ulid:    "01FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: false,
		},
		{
			name:    "B",
			prefix:  "B",
			ulid:    "01FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: false,
		},
		{
			name:    "Empty Prefix",
			prefix:  "",
			ulid:    "01FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: true,
		},
		{
			name:    "Overflow",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: true,
		},
		{
			name:    "Short (25)",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GE",
			wantErr: true,
		},
		{
			name:    "Long (27)",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GEKK",
			wantErr: true,
		},
		{
			name:    "Invalid Char 'O'",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GEO",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix, gotULID, err := Parse(tt.prefix + ":" + tt.ulid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if gotPrefix != tt.prefix {
				t.Errorf("Parse() gotPrefix = %v, want %v", gotPrefix, tt.prefix)
			}
			if gotULID.String() != tt.ulid {
				t.Errorf("Parse() goULID = %v, want %v", gotULID, tt.ulid)
			}
		})
	}
}

func TestParseStrict(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		ulid    string
		wantErr bool
	}{
		{
			name:    "AA",
			prefix:  "AA",
			ulid:    "01FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: false,
		},
		{
			name:    "B",
			prefix:  "B",
			ulid:    "01FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: false,
		},
		{
			name:    "Empty Prefix",
			prefix:  "",
			ulid:    "01FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: true,
		},
		{
			name:    "Overflow",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GEK",
			wantErr: true,
		},
		{
			name:    "Short (25)",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GE",
			wantErr: true,
		},
		{
			name:    "Long (27)",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GEKK",
			wantErr: true,
		},
		{
			name:    "Invalid Char 'O'",
			prefix:  "CC",
			ulid:    "81FD7SJ7J006AFVGQT5ZYC0GEO",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefix, gotULID, err := ParseStrict(tt.prefix + ":" + tt.ulid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if gotPrefix != tt.prefix {
				t.Errorf("Parse() gotPrefix = %v, want %v", gotPrefix, tt.prefix)
			}
			if gotULID.String() != tt.ulid {
				t.Errorf("Parse() goULID = %v, want %v", gotULID, tt.ulid)
			}
		})
	}
}
