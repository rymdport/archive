package common

import (
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"github.com/Jacalz/arkivera/internal/tar"
)

// ArchiveAndCompress provides a common way to archive and compress.
// It opens the file at source and writes the archive to the compress writer.
func ArchiveAndCompress(source string, compress io.Writer) (err error) {
	file, err := os.Open(filepath.Clean(source))
	if err != nil {
		fyne.LogError("Could not opent eh file", err)
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fyne.LogError("Could not close the file", err)
			err = cerr
		}
	}()

	return tar.Archive(source, compress)
}
