package bzip2

import (
	"compress/bzip2"
	"io"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar"
)

// Decompress takes a tar.bz2 source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) error {
	err := tar.Unarchive(bzip2.NewReader(source), target)
	if err != nil {
		fyne.LogError("Could not unarchive using tar", err)
		return err
	}

	return nil
}
