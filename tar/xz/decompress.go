package xz

import (
	"io"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar"
	"github.com/ulikunitz/xz"
)

// Decompress takes a tar.xz source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) error {
	xz, err := xz.NewReader(source)
	if err != nil {
		fyne.LogError("Could not create a gzip reader", err)
		return err
	}

	err = tar.Unarchive(xz, target)
	if err != nil {
		fyne.LogError("Could not unarchive using tar", err)
		return err
	}

	return nil
}
