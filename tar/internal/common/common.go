package common

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rymdport/archives/tar"
)

// ArchiveAndCompress provides a common way to archive and compress.
// It opens the file at source and writes the archive to the compress writer.
func ArchiveAndCompress(source string, compress io.Writer) (err error) {
	file, err := os.Open(filepath.Clean(source))
	if err != nil {
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tar.Archive(source, compress)
	return
}
