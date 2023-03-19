package xz

import (
	"io"

	"github.com/rymdport/archive/tar"
	"github.com/ulikunitz/xz"
)

// DecompressArchive takes a tar.xz source to decompress from and a target to decompress to.
func DecompressArchive(source io.Reader, target string) error {
	xz, err := xz.NewReader(source)
	if err != nil {
		return err
	}

	return tar.ExtractFromReader(xz, target)
}
