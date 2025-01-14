package mysql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		version  string
		want     string
		wantRest string
	}{
		{
			version:  "8.0.27",
			want:     "8.0.27",
			wantRest: "",
		},
		{
			version:  "5.7.22-log",
			want:     "5.7.22",
			wantRest: "-log",
		},
		{
			version:  "5.6.29_ddm_3.0.1.7",
			want:     "5.6.29",
			wantRest: "_ddm_3.0.1.7",
		},
		{
			version:  "10.4.7-MariaDB",
			want:     "10.4.7",
			wantRest: "-MariaDB",
		},
	}

	a := require.New(t)
	for _, tc := range tests {
		version, rest, err := parseVersion(tc.version)
		a.NoError(err)
		a.Equal(tc.want, version)
		a.Equal(tc.wantRest, rest)
	}
}

func TestIsNonTransactionStatement(t *testing.T) {
	tests := []struct {
		stmt string
		want bool
	}{
		{
			`CREATE DATABASE "hello" ENCODING "UTF8";`,
			false,
		},
		{
			`CREATE table hello(id integer);`,
			false,
		},
		{
			`CREATE INDEX c1 ON t1 (c1);`,
			true,
		},
		{
			`CREATE UNIQUE INDEX c1 ON t1 (c1);`,
			true,
		},
		{
			`CREATE UNIQUE INDEX c1 ON t1 (c1) INVISIBLE;`,
			true,
		},
	}

	for _, test := range tests {
		got := isNonTransactionStatement(test.stmt)
		require.Equal(t, test.want, got)
	}
}
