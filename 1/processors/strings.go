package processors

import (
	"archive/zip"
	"bufio"
	"log"
	"strings"
	"unicode"
)

type stringsProcessor struct{}

func (sp *stringsProcessor) process(file *zip.File) string {
	rc, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer rc.Close()

	s := bufio.NewScanner(rc)
	s.Split(bufio.ScanLines)
	processedContent := ""

	for s.Scan() {
		processedContent += sp.applyTransformation(s.Text()) + lineEnd
	}

	return processedContent

}

func (sp *stringsProcessor) applyTransformation(line string) string {

	transFn := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return unicode.ToLower(r)
		case r >= 'a' && r <= 'z':
			return unicode.ToUpper(r)
		}
		return r
	}

	return strings.Map(transFn, line)
}
