package tar

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateToPath uses the contents at the source path and creates a new archive at the target path.
func CreateToPath(source, target string) error {
	saveTo, err := os.Create(filepath.Clean(target))
	if err != nil {
		return err
	}

	return CreateToWriter(source, saveTo)
}

// CreateFromPathToWriter provides a common way to archive and compress.
// It opens the file at source and writes the archive to the compress writer.
func CreateFromPathToWriter(source string, compress io.Writer) (err error) {
	file, err := os.Open(filepath.Clean(source))
	if err != nil {
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = CreateToWriter(source, compress)
	return
}

// CreateToWriter creates a new tar archive in the target writer.
// The output is written to the writer that is passed.
func CreateToWriter(source string, target io.Writer) (err error) {
	tarball := tar.NewWriter(target)

	defer func() {
		if cerr := tarball.Close(); cerr != nil {
			err = cerr
		}
	}()

	info, err := os.Stat(source)
	if err != nil {
		return
	}

	baseDir := ""
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relative := ""
		if baseDir != "" {
			relative = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		return archiveFile(path, relative, info, tarball)
	})

	return
}

func archiveFile(path, relative string, info os.FileInfo, target *tar.Writer) (err error) {
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	if relative != "" {
		header.Name = relative
	}

	if err := target.WriteHeader(header); err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path) // #nosec - The received path is already cleaned
	if err != nil {
		return err
	}

	defer func() {
		cerr := file.Close()
		if cerr != nil {
			err = cerr
		}
	}()

	_, err = io.Copy(target, file)
	return
}
