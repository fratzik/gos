package processors

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	archiveHandlers "github.com/fratzik/gos/1/archive_handlers"
)

const stringsFileIdent = "_strings_"
const integersFileIdent = "_integers_"
const spaceSeparator, lineEnd, valToAdd = " ", "\n", 123

var wg sync.WaitGroup

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

func process(file *zip.File /*, tunnel chan int*/) {

	// fmt.Printf("Process the file %s\n", file.Name)
	processedContent := getProcessor(file.Name).process(file)
	writeContentToTempFile(processedContent, file.Name)
	// fmt.Println("Processing done.")
	// fi := fileInfo{file.Name, processedContent}
	wg.Done()

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
		wg.Add(1)
		process(file)
	}

	wg.Wait()

	archiveHandlers.TryToWrite("result.tar", "./tmp/")
	// fmt.Println("All good - you verify your archive. ;) ")

}
