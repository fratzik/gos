package achive_handlers

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func addFile(fileName string, tw *tar.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed opening %s: %s", fileName, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Failed to read the stat of the file %s", fileName)
	}

	header := &tar.Header{
		ModTime: stat.ModTime(),
		Name:    fileName,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode().Perm()),
	}

	if err := tw.WriteHeader(header); err != nil {
		return fmt.Errorf("Failed to write archive for %s: %s", fileName, err)
	}

	copied, err := io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("Failed to write %s to tar: %s", fileName, err)
	}

	if copied < stat.Size() {
		return fmt.Errorf("Wrote %d bytes of %s, expected to write %d", copied, fileName, stat.Size())
	}

	return nil
}

func TryToWrite(targetFile string, dirToRead string) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(targetFile, flags, 0644)
	if err != nil {
		log.Fatalf("Fail opening tar for writing %s.", targetFile)
	}
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	files, err := ioutil.ReadDir(dirToRead)

	for _, f := range files {
		if err := addFile(f.Name(), tw); err != nil {
			log.Fatalf("Failed adding file %s to tar: %s", f.Name(), err)
		}
	}
}
