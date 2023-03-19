package bzip2

import (
	"compress/bzip2"
	"io"

	"github.com/rymdport/archive/tar"
)

// DecompressArchive takes a tar.bz2 source to decompress from and a target to decompress to.
func DecompressArchive(source io.Reader, target string) error {
	return tar.ExtractFromReader(bzip2.NewReader(source), target)
}
