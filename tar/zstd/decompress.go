package zstd

import (
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/rymdport/archive/tar"
)

// DecompressArchive takes a tar.zst source to decompress from and a target to decompress to.
func DecompressArchive(source io.Reader, target string) error {
	zstd, err := zstd.NewReader(source)
	if err != nil {
		return err
	}

	defer zstd.Close() // Does not return any error value.

	return tar.ExtractFromReader(zstd, target)
}
