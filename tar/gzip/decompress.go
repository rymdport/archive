package gzip

import (
	"io"

	"github.com/klauspost/compress/gzip"
	"github.com/rymdport/archives/tar"
)

// Decompress takes a tar.gz source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) (err error) {
	gz, err := gzip.NewReader(source)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := gz.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tar.Unarchive(gz, target)
	if err != nil {
		return err
	}

	return
}
