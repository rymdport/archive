package assets

import "fyne.io/fyne/v2/theme"

//go:generate fyne bundle -package assets -o bundled.go icons

// ZipIcon returns the zip icons for the current theme.
var (
	ZipIcon = theme.NewThemedResource(resourceFolderZipSvg)
	AppIcon = resourceIcon512Png
)
