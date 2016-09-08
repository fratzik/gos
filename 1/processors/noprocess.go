package processors

import (
	"archive/zip"
	"bufio"
	"log"
)

type noprocessProcessor struct{}

func (np *noprocessProcessor) process(file *zip.File) string {

	rc, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	s := bufio.NewScanner(rc)
	s.Split(bufio.ScanLines)

	processedContent := ""

	for s.Scan() {
		processedContent += s.Text() + lineEnd
	}

	writeContentToTempFile(processedContent, file.Name)

	return ""
}
