package archives

import (
	"errors"
	"io"
	"strings"

	"github.com/rymdport/archives/tar"
	"github.com/rymdport/archives/tar/bzip2"
	"github.com/rymdport/archives/tar/gzip"
	"github.com/rymdport/archives/tar/xz"
	"github.com/rymdport/archives/tar/zstd"
	"github.com/rymdport/archives/zip"
)

var (
	errorInvalidFormat           = errors.New("unsupported archive format")
	errorZipUnarchiveNotPossible = errors.New("zip unarchive must conform to io.Seeker and io.ReaderAt")
)

// Archive creates a new archive based on the given source and extention.
// The output is written to the writer that is passed.
func Archive(source string, target io.Writer, ext string) error {
	switch strings.TrimPrefix(ext, ".") {
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

// Unarchive will take a source reader, the extension for selecting correct method and the target to unarchive to.
// Note that .zip files requires the reader to conform to both io.Seeker and io.ReaderAt to work.
func Unarchive(source io.Reader, ext, target string) error {
	switch strings.TrimPrefix(ext, ".") {
	case "tar": // No compression
		return tar.Unarchive(source, target)
	case "tar.gz", "tgz":
		return gzip.Decompress(source, target)
	case "tar.xz", "txz":
		return xz.Decompress(source, target)
	case "tar.zst", "tzst":
		return zstd.Decompress(source, target)
	case "tar.bz2", "tbz2":
		return bzip2.Decompress(source, target)
	case "zip":
		return tryExtractZip(source, target)
	}

	return errorInvalidFormat
}

func tryExtractZip(source io.Reader, target string) error {
	reader, ok := source.(io.ReaderAt)
	if !ok {
		return errorZipUnarchiveNotPossible
	}

	seek, ok := source.(io.Seeker)
	if !ok {
		return errorZipUnarchiveNotPossible
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

	return zip.Extract(reader, size, target)
}
