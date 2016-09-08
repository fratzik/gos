package processors

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const stringsFileIdent = "_strings_"
const integersFileIdent = "_integers_"
const spaceSeparator, lineEnd, valToAdd = " ", "\n", 123

type fileInfo struct {
	ProcessedContent string
	FileName         string
}

func shouldProcessStrings(fileName string) bool {
	return strings.Contains(fileName, stringsFileIdent)
}

func shouldProcessIntegers(fileName string) bool {
	return strings.Contains(fileName, integersFileIdent)
}

func getProcessor(fileName string) Processor {

	if shouldProcessIntegers(fileName) {
		return new(integersProcessor)
	} else if shouldProcessStrings(fileName) {
		return new(stringsProcessor)
	}

	return new(noprocessProcessor)
}

func process(file *zip.File, tunnel chan fileInfo) {

	fmt.Printf("Process the file %s\n", file.Name)
	processedContent := getProcessor(file.Name).process(file)
	writeContentToTempFile(processedContent, file.Name)
	fi := fileInfo{file.Name, processedContent}
	tunnel <- fi

}

func writeContentToTempFile(content string, fileName string) {
	contentToWrite := []byte(content)
	tmpfile, err := ioutil.TempFile("tmp", fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(contentToWrite); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

// ProcessArchive that needs to be transformed
func ProcessArchive(archiveName string) {
	rc, err := zip.OpenReader(archiveName)

	if err != nil {
		log.Fatalf("Failed opening the archive file %s - %s", archiveName, err)
	}

	defer rc.Close()

	for _, file := range rc.File {
		ch := make(chan fileInfo)
		process(file, ch)

		fi := <-ch

		fmt.Printf("Res: %v", fi)
	}
}
