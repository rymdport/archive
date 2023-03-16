package gzip

import (
	"io"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar/internal/common"
	"github.com/klauspost/pgzip"
)

// Compress takes a source to compress and a target to compress and archive to.
func Compress(source string, target io.Writer) (err error) {
	gz := pgzip.NewWriter(target)

	defer func() {
		if cerr := gz.Close(); cerr != nil {
			fyne.LogError("Could not close the pgzip writer", err)
			err = cerr
		}
	}()

	err = common.ArchiveAndCompress(source, gz)
	if err != nil {
		fyne.LogError("Could not archive and compress", err)
		return err
	}

	return
}
