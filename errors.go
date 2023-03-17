package archive

import "errors"

var (
	ErrInvalidFormat           = errors.New("unsupported archive format")
	ErrZipUnarchiveNotPossible = errors.New("zip unarchive must conform to io.Seeker and io.ReaderAt")
)
