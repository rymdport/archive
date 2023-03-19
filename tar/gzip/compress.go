package gzip

import (
	"io"

	"github.com/klauspost/compress/gzip"
	"github.com/rymdport/archive/tar"
)

// ArchiveAndCompress takes a source to compress and a target to compress and archive to.
func ArchiveAndCompress(source string, target io.Writer) (err error) {
	gz := gzip.NewWriter(target)

	defer func() {
		if cerr := gz.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tar.CreateFromPathToWriter(source, gz)
	return
}
