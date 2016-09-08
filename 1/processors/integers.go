package processors

import (
	"archive/zip"
	"bufio"
	"log"
	"strconv"
	"strings"
)

type integersProcessor struct{}

func (ip *integersProcessor) process(file *zip.File) string {

	rc, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	s := bufio.NewScanner(rc)
	s.Split(bufio.ScanLines)

	processedContent := ""

	//scan each line
	for s.Scan() {
		processedContent += ip.applyTransformation(s.Text()) + lineEnd
	}

	return processedContent

}

func (ip *integersProcessor) applyTransformation(line string) string {
	tokens := strings.Split(line, spaceSeparator)
	for pos, val := range tokens {
		i, err := strconv.Atoi(val)
		if err == nil {
			tokens[pos] = strconv.Itoa(i + valToAdd)
		}
	}

	return strings.Join(tokens, spaceSeparator)
}
