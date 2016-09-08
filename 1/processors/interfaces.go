package processors

import "archive/zip"

// Processor interface
type Processor interface {
	process(file *zip.File) string
}
