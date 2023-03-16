package zstd

import (
	"io"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar/internal/common"
	"github.com/klauspost/compress/zstd"
)

// Compress takes a source to compress and a target to compress and archive to.
func Compress(source string, target io.Writer) (err error) {
	zstd, err := zstd.NewWriter(target)
	if err != nil {
		fyne.LogError("Could not create a zstd writer", err)
		return err
	}

	defer func() {
		if cerr := zstd.Close(); cerr != nil {
			fyne.LogError("Could not close the zstd writer", err)
			err = cerr
		}
	}()

	err = common.ArchiveAndCompress(source, zstd)
	if err != nil {
		fyne.LogError("Could not archive and compress", err)
		return err
	}

	return
}
