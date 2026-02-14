package vbantxt_test

import (
	"bufio"
	"bytes"
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onyx-and-iris/vbantxt"
)

//go:embed testdata/vm.txt
var vm []byte

//go:embed testdata/matrix.txt
var matrix []byte

func run(t *testing.T, client *vbantxt.VbanTxt, script []byte) {
	t.Helper()

	r := bytes.NewReader(script)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := client.Send(scanner.Text())
		require.NoError(t, err)
	}
}

func TestSendVm(t *testing.T) {
	client, err := vbantxt.New(os.Getenv("VBANTXT_HOST"), 6980, os.Getenv("VBANTXT_STREAMNAME"))
	require.NoError(t, err)

	run(t, client, vm)
}

func TestSendMatrix(t *testing.T) {
	client, err := vbantxt.New(os.Getenv("VBANTXT_HOST"), 6990, os.Getenv("VBANTXT_STREAMNAME"))
	require.NoError(t, err)

	run(t, client, matrix)
}
