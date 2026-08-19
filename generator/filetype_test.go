package main

import (
	"os"
	"testing"

	"github.com/cavaliergopher/cpio"
	"github.com/stretchr/testify/require"
)

func TestFileType(t *testing.T) {
	dir := t.TempDir()
	check := func(compression, expectedType string) {
		fileName := dir + "/" + compression
		img, err := NewImage(fileName, compression, false)
		require.NoError(t, err)

		require.NoError(t, img.AppendEntry("foo.txt", cpio.TypeReg, []byte("hello, world!")))
		require.NoError(t, img.Close())

		f, err := os.Open(fileName)
		require.NoError(t, err)

		kind, err := filetype(f)
		require.NoError(t, err)

		require.Equal(t, expectedType, kind)
	}

	check("zstd", "zstd")
	check("gzip", "gzip")
	check("xz", "xz")
	check("lz4", "lz4")
	check("none", "cpio")

	// Test empty file
	emptyFile, err := os.CreateTemp(dir, "empty")
	require.NoError(t, err)
	defer emptyFile.Close()
	kind, err := filetype(emptyFile)
	require.NoError(t, err)
	require.Equal(t, "", kind)

	// Test unknown small file
	unknownFile, err := os.CreateTemp(dir, "unknown")
	require.NoError(t, err)
	defer unknownFile.Close()
	_, err = unknownFile.Write([]byte("random"))
	require.NoError(t, err)
	kind, err = filetype(unknownFile)
	require.NoError(t, err)
	require.Equal(t, "", kind)
}
