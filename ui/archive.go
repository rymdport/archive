package ui

import (
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Jacalz/arkivera/internal/all"
)

var archives = []string{".tar", ".tar.gz", ".tar.xz", ".tar.zst", ".zip"}

type archiveUI struct {
	toArchive    *widget.Entry
	opened       fyne.URI
	file         *widget.Button
	fileDialog   *dialog.FileDialog
	folder       *widget.Button
	folderDialog *dialog.FileDialog

	archiveName *widget.Entry
	archiveType *widget.Select

	window fyne.Window
}

func (a *archiveUI) selectFile(file fyne.URIReadCloser, err error) {
	if err != nil {
		fyne.LogError("Could not open the file", err)
		dialog.ShowError(err, a.window)
		return
	} else if file == nil {
		return
	}

	// Text needs to be set before updating the opened uri.
	a.toArchive.SetText(file.URI().Path())
	a.archiveName.SetText(file.URI().Name())
	a.opened = file.URI()

	if err := file.Close(); err != nil {
		fyne.LogError("Could not close the file", err)
	}
}

func (a *archiveUI) selectFolder(folder fyne.ListableURI, err error) {
	if err != nil {
		fyne.LogError("Could not open the folder", err)
		dialog.ShowError(err, a.window)
		return
	} else if folder == nil {
		return
	}

	// Text needs to be set before updating the opened uri.
	a.toArchive.SetText(folder.Path())
	a.archiveName.SetText(folder.Name())
	a.opened = folder
}

func (a *archiveUI) archive(source fyne.URI, target io.Writer, ext string) error {
	err := all.Archive(source.Path(), target, ext)
	if err != nil {
		fyne.LogError("Could not create the archive", err)
		return err
	}

	return nil
}

// newArchiver creates the ui used for archiving files and folders.
func newArchiver(path fyne.URI, w fyne.Window) dialog.Dialog {
	a := &archiveUI{window: w}

	a.fileDialog = dialog.NewFileOpen(a.selectFile, w)
	a.folderDialog = dialog.NewFolderOpen(a.selectFolder, w)

	a.toArchive = &widget.Entry{PlaceHolder: "Path to file/folder to archive...", OnChanged: func(_ string) {
		a.opened = nil // Invalidate saved file or folder.
	}}
	a.file = &widget.Button{Icon: theme.FileIcon(), OnTapped: a.fileDialog.Show}
	a.folder = &widget.Button{Icon: theme.FolderOpenIcon(), OnTapped: a.folderDialog.Show}

	a.archiveName = &widget.Entry{PlaceHolder: "Name of archive..."}
	a.archiveType = &widget.Select{Options: archives, Selected: ".tar.zst"}

	content := container.NewBorder(nil, nil /*container.NewHBox(layout.NewSpacer(), a.cancel, a.create)*/, nil, nil,
		container.NewVBox(
			container.NewBorder(nil, nil, nil, container.NewHBox(a.file, a.folder), a.toArchive),
			container.NewBorder(nil, nil, nil, a.archiveType, a.archiveName),
		),
	)

	custom := dialog.NewCustomConfirm("Create a new archive", "Create", "Cancel", content, func(create bool) {
		if create {
			if a.opened == nil { // If the filename was set in the entry.
				a.opened = storage.NewFileURI(a.toArchive.Text)
			}

			file, err := os.Create(filepath.Join(filepath.Dir(a.opened.Path()), a.archiveName.Text+a.archiveType.Selected))
			if err != nil {
				fyne.LogError("Could not open the target file", err)
				dialog.ShowError(err, a.window)
				return
			}

			defer func() {
				if cerr := file.Close(); cerr != nil {
					fyne.LogError("Could not close the file", cerr)
				}
			}()

			if err = a.archive(a.opened, file, a.archiveType.Selected); err != nil {
				dialog.ShowError(err, a.window)
			}
		}

		a.opened = nil
		a.toArchive.SetText("")
		a.archiveName.SetText("")
	}, w)

	if path != nil && !filterUnarchive.Matches(path) {
		// Text needs to be set before updating the opened uri.
		a.toArchive.SetText(path.Path())
		a.archiveName.SetText(path.Name())
		a.opened = path

		custom.Show()
	}

	return custom
}
