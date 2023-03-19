package xz

import (
	"io"

	"github.com/rymdport/archive/tar"
	"github.com/ulikunitz/xz"
)

// ArchiveAndCompress takes a source to compress and a target to compress and archive to.
func ArchiveAndCompress(source string, target io.Writer) (err error) {
	xz, err := xz.NewWriter(target)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := xz.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = tar.CreateFromPathToWriter(source, xz)
	return
}
