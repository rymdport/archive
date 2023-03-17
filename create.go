package archive

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rymdport/archive/tar"
	"github.com/rymdport/archive/tar/gzip"
	"github.com/rymdport/archive/tar/xz"
	"github.com/rymdport/archive/tar/zstd"
	"github.com/rymdport/archive/zip"
)

// CreateFromPath uses the contents at the source path and creates a new archive at the target path.
func CreateFromPath(source, target string) error {
	saveTo, err := os.Create(filepath.Clean(target))
	if err != nil {
		return err
	}

	return CreateToWriter(source, saveTo, extensionsFromFile(target))
}

// CreateToWriter creates a new archive based on the given source and extention.
// The output is written to the writer that is passed.
func CreateToWriter(source string, target io.Writer, ext Format) error {
	switch ext {
	case Tar:
		return tar.Archive(source, target)
	case TarGzip, TarGzipShort:
		return gzip.Compress(source, target)
	case TarXz, TarXzShort:
		return xz.Compress(source, target)
	case TarZstd, TarZstdShort:
		return zstd.Compress(source, target)
	case Zip:
		return zip.Archive(source, target)
	}

	return ErrInvalidFormat
}
