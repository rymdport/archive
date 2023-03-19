package zip

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/klauspost/compress/zip"
)

// CreateToWriter creates a new zip archive in the target writer.
// The output is written to the writer that is passed.
func CreateToWriter(source string, target io.Writer) (err error) {
	writer := zip.NewWriter(target)

	defer func() {
		if cerr := writer.Close(); cerr != nil {
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
		return archiveFile(path, relative, info, writer)
	})

	return
}

func archiveFile(path, relative string, info os.FileInfo, target *zip.Writer) (err error) {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	if relative != "" {
		header.Name = relative
	}

	if info.IsDir() {
		header.Name += "/"
	} else {
		header.Method = zip.Deflate
	}

	writer, err := target.CreateHeader(header)
	if err != nil {
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

	_, err = io.Copy(writer, file)
	return
}
