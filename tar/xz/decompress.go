package xz

import (
	"io"

	"github.com/rymdport/archives/tar"
	"github.com/ulikunitz/xz"
)

// Decompress takes a tar.xz source to decompress from and a target to decompress to.
func Decompress(source io.Reader, target string) error {
	xz, err := xz.NewReader(source)
	if err != nil {
		return err
	}

	err = tar.Unarchive(xz, target)
	if err != nil {
		return err
	}

	return nil
}
