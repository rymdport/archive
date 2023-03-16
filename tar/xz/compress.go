package xz

import (
	"io"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar/internal/common"
	"github.com/ulikunitz/xz"
)

// Compress takes a source to compress and a target to compress and archive to.
func Compress(source string, target io.Writer) (err error) {
	xz, err := xz.NewWriter(target)
	if err != nil {
		fyne.LogError("Could not create a xz writer", err)
		return err
	}

	defer func() {
		if cerr := xz.Close(); cerr != nil {
			fyne.LogError("Could not close the xz writer", err)
			err = cerr
		}
	}()

	err = common.ArchiveAndCompress(source, xz)
	if err != nil {
		fyne.LogError("Could not archive and compress", err)
		return err
	}

	return
}
