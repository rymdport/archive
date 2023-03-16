package zstd

import (
	"io"

	"github.com/klauspost/compress/zstd"
	"github.com/rymdport/archives/tar/internal/common"
)

// Compress takes a source to compress and a target to compress and archive to.
func Compress(source string, target io.Writer) (err error) {
	zstd, err := zstd.NewWriter(target)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := zstd.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = common.ArchiveAndCompress(source, zstd)
	if err != nil {
		return err
	}

	return
}
