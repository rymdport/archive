package tar

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
)

// Unarchive takes a reader as a tarball source and extracts it to the target directory.
func Unarchive(source io.Reader, target string) (err error) {
	if err = os.MkdirAll(target, 0750); err != nil {
		fyne.LogError("Could not create the directory", err)
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

		// TODO: Fix potential directory traversal security issue
		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, 0750); err != nil {
				return err
			}
			continue
		}

		if err = extractFile(path, info, tarball); err != nil {
			fyne.LogError("Could not extract the file", err)
			return err
		}
	}

	return
}

func extractFile(path string, info os.FileInfo, target *tar.Reader) (err error) {
	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		fyne.LogError("Could not create the directory", err)
		return err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, info.Mode()) // #nosec - path is already cleaned by filepath.Join()
	if err != nil {
		fyne.LogError("Could not create the target file", err)
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fyne.LogError("Could not close the target file", err)
			err = cerr
		}
	}()

	_, err = io.Copy(file, target)
	if err != nil {
		fyne.LogError("Could not copy file contents", err)
		return err
	}

	return
}
