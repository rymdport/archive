package all

import (
	"errors"
	"io"

	"github.com/Jacalz/arkivera/internal/tar"
	"github.com/Jacalz/arkivera/internal/tar/gzip"
	"github.com/Jacalz/arkivera/internal/tar/xz"
	"github.com/Jacalz/arkivera/internal/tar/zstd"
	"github.com/Jacalz/arkivera/internal/zip"
)

var errorInvalidFormat = errors.New("unsupported archive format")

// Archive will
func Archive(source string, target io.Writer, ext string) error {
	switch getExtension(ext) {
	case "tar": // No compression
		return tar.Archive(source, target)
	case "tar.gz", "tgz":
		return gzip.Compress(source, target)
	case "tar.xz", "txz":
		return xz.Compress(source, target)
	case "tar.zst", "tzst":
		return zstd.Compress(source, target)
	case "zip":
		return zip.Archive(source, target)
	}

	return errorInvalidFormat
}

func getExtension(ext string) string {
	if ext[0] == '.' {
		return ext[1:]
	}

	return ext
}
