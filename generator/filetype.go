package main

import (
	"bytes"
	"io"
	"os"
)

var fileSignatures = []struct {
	name   string
	marker []byte
}{
	{"cpio", []byte{'0', '7', '0', '7', '0', '1'}}, // "new" cpio format
	{"xz", []byte{0xfd, '7', 'z', 'X', 'Z', 0x00}},
	{"lz4", []byte{0x02, 0x21, 0x4c, 0x18}}, // legacy format used by linux loader
	{"zstd", []byte{0x28, 0xb5, 0x2f, 0xfd}},
	{"gzip", []byte{0x1f, 0x8b}},
}

func filetype(r *os.File) (string, error) {
	loc, err := r.Seek(0, io.SeekCurrent)
	if err != nil {
		return "", err
	}
	defer r.Seek(loc, io.SeekStart)

	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	var header [6]byte
	n, err := io.ReadFull(r, header[:])
	if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
		return "", err
	}

	for _, s := range fileSignatures {
		if n >= len(s.marker) && bytes.Equal(header[:len(s.marker)], s.marker) {
			return s.name, nil
		}
	}

	return "", nil
}
