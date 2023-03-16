package zstd

import (
	"io"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar"
	"github.com/klauspost/compress/zstd"
)

// Decompress takes a tar.zst source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) error {
	zstd, err := zstd.NewReader(source)
	if err != nil {
		fyne.LogError("Could not create a gzip reader", err)
		return err
	}

	defer zstd.Close() // Does not return any error value.

	err = tar.Unarchive(zstd, target)
	if err != nil {
		fyne.LogError("Could not unarchive using tar", err)
		return err
	}

	return nil
}
