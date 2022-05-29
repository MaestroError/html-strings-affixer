package scanning

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestScanningStructWorks(t *testing.T) {
	folderName := "ScanningStructWorksFolder"
	os.Mkdir(folderName, 0755)
	path := filepath.Join(folderName, "testing.txt")
	path_not_found := filepath.Join(folderName, "testing.html")
	createFile(path)
	createFile(path_not_found)
	allowed_file_types := []string{".txt"}

	Scan := Scanning{}
	result := Scan.Get(folderName, allowed_file_types, []string{".denyFIleType"})

	if result[0] != path {
		t.Errorf("got %q, wanted %q", result[0], path)
	}

	defer os.RemoveAll(folderName)

}

func BenchmarkScanningSpeed(b *testing.B) {

}

func ExampleScanningStructOutput() {
	fmt.Println("hello")
	// Output: hello
}

func createFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
