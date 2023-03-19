package gzip

import (
	"io"

	"github.com/klauspost/compress/gzip"
	"github.com/rymdport/archive/tar"
)

// DecompressArchive takes a tar.gz source to decompress from and a target to decompress to.
func DecompressArchive(source io.Reader, target string) (err error) {
	gz, err := gzip.NewReader(source)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := gz.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tar.ExtractFromReader(gz, target)
	return
}
