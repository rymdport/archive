package zstd

import (
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/rymdport/archives/tar"
)

// Decompress takes a tar.zst source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) error {
	zstd, err := zstd.NewReader(source)
	if err != nil {
		return err
	}

	defer zstd.Close() // Does not return any error value.

	err = tar.Unarchive(zstd, target)
	if err != nil {
		return err
	}

	return nil
}
