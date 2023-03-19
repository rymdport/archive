package zstd

import (
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/rymdport/archive/tar"
)

// ArchiveAndCompress takes a source to compress and a target to compress and archive to.
func ArchiveAndCompress(source string, target io.Writer) (err error) {
	zstd, err := zstd.NewWriter(target)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := zstd.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tar.CreateFromPathToWriter(source, zstd)
	return
}
