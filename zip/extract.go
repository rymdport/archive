package zip

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zip"
)

var errDangerousFilename = errors.New("dangerous filename detected")

// ExtractFromReader takes a source reader and a target to extract to.
func ExtractFromReader(source io.ReaderAt, length int64, target string) error {
	reader, err := zip.NewReader(source, length)
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		if err := extractFile(file, target); err != nil {
			return err
		}
	}

	return nil
}

func extractFile(file *zip.File, target string) (err error) {
	path, err := filepath.Abs(filepath.Join(target, file.Name))
	if err != nil {
		return err
	}

	if !strings.HasPrefix(path, target) {
		return errDangerousFilename
	}

	fileReader, err := file.Open()
	if err != nil {
		return err
	}

	defer func() {
		if cerr := fileReader.Close(); cerr != nil {
			err = cerr
		}
	}()

	if file.FileInfo().IsDir() {
		err = os.MkdirAll(path, 0o750)
		return
	}

	err = os.MkdirAll(filepath.Dir(path), 0o750)
	if err != nil {
		return
	}

	targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode()) // #nosec - path is already cleaned by filepath.Abs()
	if err != nil {
		return
	}

	defer func() {
		if cerr := targetFile.Close(); cerr != nil {
			err = cerr
		}
	}()

	_, err = io.Copy(targetFile, fileReader)
	return
}
