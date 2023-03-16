package gzip

import (
	"io"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar"
	"github.com/klauspost/pgzip"
)

// Decompress takes a tar.gz source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) (err error) {
	gz, err := pgzip.NewReader(source)
	if err != nil {
		fyne.LogError("Could not create a gzip reader", err)
		return err
	}

	defer func() {
		if cerr := gz.Close(); cerr != nil {
			fyne.LogError("Could not close the pgzip writer", err)
			err = cerr
		}
	}()

	err = tar.Unarchive(gz, target)
	if err != nil {
		fyne.LogError("Could not unarchive using tar", err)
		return err
	}

	return
}
