package archive

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rymdport/archive/tar"
	"github.com/rymdport/archive/tar/bzip2"
	"github.com/rymdport/archive/tar/gzip"
	"github.com/rymdport/archive/tar/xz"
	"github.com/rymdport/archive/tar/zstd"
	"github.com/rymdport/archive/zip"
)

// ExtractFromPath extracts the archive at the source path to the target path.
func ExtractFromPath(source, target string) error {
	readFrom, err := os.Open(filepath.Clean(source))
	if err != nil {
		return err
	}

	ext := extensionsFromFile(source)
	if ext == Zip { // We know that *os.File can use ReaderAt() and Stat().
		info, err := readFrom.Stat()
		if err != nil {
			return err
		}

		return zip.ExtractFromReader(readFrom, info.Size(), target)
	}

	return ExtractFromReader(readFrom, target, ext)
}

// ExtractFromReader takes a source reader, a path to save the archive at and a format to select archive and compression type.
// Note that .zip files requires the reader to conform to both io.Seeker and io.ReaderAt to work.
func ExtractFromReader(source io.Reader, target string, ext Format) error {
	switch ext {
	case Tar:
		return tar.ExtractFromReader(source, target)
	case TarGzip, TarGzipShort:
		return gzip.DecompressArchive(source, target)
	case TarXz, TarXzShort:
		return xz.DecompressArchive(source, target)
	case TarZstd, TarZstdShort:
		return zstd.DecompressArchive(source, target)
	case TarBzip2, TarBzip2Short:
		return bzip2.DecompressArchive(source, target)
	case Zip:
		return tryExtractZip(source, target)
	}

	return ErrInvalidFormat
}

func tryExtractZip(source io.Reader, target string) error {
	reader, ok := source.(io.ReaderAt)
	if !ok {
		return ErrZipUnarchiveNotPossible
	}

	seek, ok := source.(io.Seeker)
	if !ok {
		return ErrZipUnarchiveNotPossible
	}

	size, err := seek.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	// Seek back to start to reset the reading offset
	_, err = seek.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	return zip.ExtractFromReader(reader, size, target)
}
