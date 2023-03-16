package ui

import (
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/Jacalz/arkivera/internal/all"
	"github.com/Jacalz/arkivera/internal/assets"
)

// Supported for extraction.
var filterUnarchive = storage.NewExtensionFileFilter([]string{
	".tar", ".zip",
	".tbz2", ".bz2",
	".tgz", ".gz",
	".txz", ".xz",
	".tzst", ".zst",
})

type unarchiveUI struct {
	toUnarchive *widget.Entry
	openFile    *widget.Button
	fileDialog  *dialog.FileDialog
	opened      fyne.URIReadCloser

	cancel  *widget.Button
	extract *widget.Button

	window fyne.Window
}

func (u *unarchiveUI) selectFile(file fyne.URIReadCloser, err error) {
	if err != nil {
		fyne.LogError("Could not open the file", err)
		dialog.ShowError(err, u.window)
		return
	} else if file == nil {
		return
	}

	// Text needs to be set before updating the opened uri.
	u.toUnarchive.SetText(file.URI().Path())
	u.opened = file
}

func (u *unarchiveUI) unarchive(archive fyne.URIReadCloser) error {
	err := all.Unarchive(archive, archive.URI().Extension(), filepath.Dir(archive.URI().Path()))
	if err != nil {
		fyne.LogError("Could not extract the archive", err)
		return err
	}

	return nil
}

func newUnarchiver(path fyne.URI, w fyne.Window) *fyne.Container {
	u := &unarchiveUI{window: w}

	u.fileDialog = dialog.NewFileOpen(u.selectFile, w)
	u.fileDialog.SetFilter(filterUnarchive)
	u.openFile = &widget.Button{Icon: assets.ZipIcon, OnTapped: u.fileDialog.Show}
	u.toUnarchive = &widget.Entry{PlaceHolder: "Archive to extract...", OnChanged: func(_ string) {
		u.opened = nil // Invalidate saved file.
	}}

	if path != nil && filterUnarchive.Matches(path) {
		uri, err := storage.Reader(path)
		if err != nil {
			fyne.LogError("Could not open the file for reading", err)
		} else {
			u.toUnarchive.SetText(uri.URI().Path()) // Text needs to be set before updating the opened uri.
			u.opened = uri
		}
	}

	u.cancel = &widget.Button{Text: "Cancel", OnTapped: func() {
		u.toUnarchive.SetText("")
	}}

	u.extract = &widget.Button{Text: "Extract", Importance: widget.HighImportance, OnTapped: func() {
		if u.opened == nil {
			uri, err := storage.Reader(storage.NewFileURI(u.toUnarchive.Text))
			if err != nil {
				fyne.LogError("Could not create reader from the uri", err)
				dialog.ShowError(err, w)
				return
			}

			u.opened = uri
		}

		defer func() {
			if cerr := u.opened.Close(); cerr != nil {
				fyne.LogError("Could not close the file", cerr)
			}
			u.opened = nil
		}()

		if err := u.unarchive(u.opened); err != nil {
			dialog.ShowError(err, w)
		}
	}}

	return container.NewBorder(nil, container.NewHBox(layout.NewSpacer(), u.cancel, u.extract), nil, nil,
		container.NewVBox(
			container.NewBorder(nil, nil, nil, u.openFile, u.toUnarchive),
		),
	)
}
