package gzip

import (
	"io"

	"github.com/klauspost/compress/gzip"
	"github.com/rymdport/archives/tar/internal/common"
)

// Compress takes a source to compress and a target to compress and archive to.
func Compress(source string, target io.Writer) (err error) {
	gz := gzip.NewWriter(target)

	defer func() {
		if cerr := gz.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = common.ArchiveAndCompress(source, gz)
	return
}
