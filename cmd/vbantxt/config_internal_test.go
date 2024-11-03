package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig_Success(t *testing.T) {
	conn, err := loadConfig("testdata/config.toml")
	require.NoError(t, err)

	assert.Equal(t, conn.Host, "localhost")
	assert.Equal(t, conn.Port, 7000)
	assert.Equal(t, conn.Streamname, "vbantxt")
}

func TestLoadConfig_Errors(t *testing.T) {
	tt := map[string]struct {
		input string
		err   string
	}{
		"no such file": {
			input: "/no/such/dir/config.toml",
			err:   "no such file or directory",
		},
	}

	for name, tc := range tt {
		_, err := loadConfig("/no/such/dir/config.toml")

		t.Run(name, func(t *testing.T) {
			assert.Error(t, err)
			assert.ErrorContains(t, err, tc.err)
		})
	}

}
