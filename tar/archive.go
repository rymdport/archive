package tar

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Archive creates a tar archive using the target writer provided.
func Archive(source string, target io.Writer) (err error) {
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
