package processors

import (
	"archive/zip"
	"log"
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

// Process a file with the specifications for the theme one
func process(file *zip.File, tunnel chan fileInfo) {

	// fmt.Printf("Process the file %s\n", file.Name)
	processedContent := getProcessor(file.Name).process(file)
	fi := fileInfo{file.Name, processedContent}
	tunnel <- fi

}

// ProcessArchive that needs to be transformed
func ProcessArchive(archiveName string) {
	rc, err := zip.OpenReader(archiveName)

	if err != nil {
		log.Fatalf("Failed opening the archive file %s - %s", archiveName, err)
	}

	defer rc.Close()

	// channels := make([]chan fileInfo, len(rc.File))

	for _, file := range rc.File {
		go process(file, make(chan fileInfo))
	}
}
