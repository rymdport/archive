package archive

import "path/filepath"

type Format = string

const (
	Zip Format = ".zip" // Zip file extension for a regular zip archive.

	Tar Format = ".tar" // Tar file extension for a regular tar archive.

	// Long versions (two extensions) of compressed tar archive formats.
	TarBzip2 Format = ".tar.bz2" // TarBzip2 file extension for bzip2 compressed tar archive.
	TarGzip  Format = ".tar.gz"  // TarGzip file extension for gzip compressed tar archive.
	TarXz    Format = ".tar.xz"  // TarXz file extension for xz compressed tar archive.
	TarZstd  Format = ".tar.zst" // TarZstd file extension for zstd compressed tar archive.

	// Short versions (one extension) of compressed tar archive formats.
	TarBzip2Short Format = ".tbz2" // TarBzip2 file extension for bzip2 compressed tar archive.
	TarGzipShort  Format = ".tgz"  // TarGzip file extension for gzip compressed tar archive.
	TarXzShort    Format = ".txz"  // TarXz file extension for xz compressed tar archive.
	TarZstdShort  Format = ".tzst" // TarZstd file extension for zstd compressed tar archive.
)

// extensionsFromFile grabs the two last file extensions of a file.
// This make sure that extensions like .tar.gz gets treated as one format.
func extensionsFromFile(path string) Format {
	second := filepath.Ext(path)

	argumentStart := len(path) - len(second)
	first := filepath.Ext(path[:argumentStart])

	argumentStart -= len(first)
	return path[argumentStart:]
}
