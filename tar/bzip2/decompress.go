package bzip2

import (
	"compress/bzip2"
	"io"

	"github.com/rymdport/archives/tar"
)

// Decompress takes a tar.bz2 source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) error {
	return tar.Unarchive(bzip2.NewReader(source), target)
}
