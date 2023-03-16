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

	baseDir := ""
	if info, err := os.Stat(source); err != nil {
		return err
	} else if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	if err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relative := ""
		if baseDir != "" {
			relative = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		return archiveFile(path, relative, info, tarball)
	}); err != nil {
		return err
	}

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
		if err != nil {
			err = cerr
		}
	}()

	_, err = io.Copy(target, file)
	if err != nil {
		return err
	}

	return
}
