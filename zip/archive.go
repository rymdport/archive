package zip

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/klauspost/compress/zip"
)

// Archive takes a source and a writer for the target to put the resulting archive.
func Archive(source string, target io.Writer) (err error) {
	writer := zip.NewWriter(target)

	defer func() {
		if cerr := writer.Close(); cerr != nil {
			fyne.LogError("Could not close the file", err)
			err = cerr
		}
	}()

	baseDir := ""
	if info, err := os.Stat(source); err != nil {
		fyne.LogError("Could not stat the source", err)
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
		return archiveFile(path, relative, info, writer)
	}); err != nil {
		fyne.LogError("Could not walk", err)
		return err
	}

	return
}

func archiveFile(path, relative string, info os.FileInfo, target *zip.Writer) (err error) {
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		fyne.LogError("Could not create header from fileinfo", err)
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
		fyne.LogError("Could not create header", err)
		return err
	}

	if info.IsDir() {
		return nil
	}

	file, err := os.Open(path) // #nosec - The received path is already cleaned
	if err != nil {
		fyne.LogError("Could not open the file", err)
		return err
	}

	defer func() {
		cerr := file.Close()
		if err != nil {
			fyne.LogError("Could not close the file", err)
			err = cerr
		}
	}()

	_, err = io.Copy(writer, file)
	if err != nil {
		fyne.LogError("Could not copy file contents", err)
		return err
	}

	return
}
