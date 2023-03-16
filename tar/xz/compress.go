package xz

import (
	"io"

	"github.com/rymdport/archives/tar/internal/common"
	"github.com/ulikunitz/xz"
)

// Compress takes a source to compress and a target to compress and archive to.
func Compress(source string, target io.Writer) (err error) {
	xz, err := xz.NewWriter(target)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := xz.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = common.ArchiveAndCompress(source, xz)
	if err != nil {
		return err
	}

	return
}
