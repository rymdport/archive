package ui

import (
	"errors"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/Jacalz/arkivera/internal/assets"
)

// Create sets up the user interface.
func Create(w fyne.Window) *fyne.Container {
	var path fyne.URI
	if len(os.Args) > 1 {
		cleaned := filepath.Clean(os.Args[1])
		if _, err := os.Stat(cleaned); err == nil || errors.Is(err, os.ErrExist) {
			path = storage.NewFileURI(cleaned)
		}
	}

	archiver := newArchiver(path, w)
	unarchiver := newUnarchiver(path, w)

	archive := &widget.Button{Text: "New archive...", Icon: assets.ZipIcon, OnTapped: archiver.Show}
	toolbar := container.NewHBox(archive, layout.NewSpacer())

	return container.NewBorder(toolbar, nil, nil, nil, unarchiver)
}
