package xz

import (
	"io"

	"github.com/rymdport/archive/tar"
	"github.com/ulikunitz/xz"
)

// Decompress takes a tar.xz source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) error {
	xz, err := xz.NewReader(source)
	if err != nil {
		return err
	}

	return tar.Unarchive(xz, target)
}
