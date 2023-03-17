package tar

import (
	"archive/tar"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var errDangerousFilename = errors.New("dangerous filename detected")

// Unarchive takes a reader as a tarball source and extracts it to the target directory.
func Unarchive(source io.Reader, target string) (err error) {
	if err = os.MkdirAll(target, 0o750); err != nil {
		return err
	}

	tarball := tar.NewReader(source)

	for {
		header, err := tarball.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		if !strings.HasPrefix(path, target) {
			return errDangerousFilename
		}

		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, 0o750); err != nil {
				return err
			}
			continue
		}

		if err = extractFile(path, info, tarball); err != nil {
			return err
		}
	}

	return
}

func extractFile(path string, info os.FileInfo, target *tar.Reader) (err error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, info.Mode()) // #nosec - path is already cleaned by filepath.Join()
	if err != nil {
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = cerr
		}
	}()

	_, err = io.Copy(file, target)
	return
}
