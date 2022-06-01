package main

import (
	"fmt"

	"github.com/MaestroError/html-strings-affixer/parsehtml"
	"github.com/MaestroError/html-strings-affixer/scanning"
)

// "testdata", []string{"index.blade.php", "denyThisFolder"}, []string{".blade.php", ".jsx"}
func main() {
	files := scanFolder()
	fmt.Println(files)

	parse := parsehtml.Parsehtml{}
	parse.Init(files[0])

	parse.AddNewString("some founded string", "some founded string to replace")
	parse.AddNewString("some founded 1", "string to replace 1")

	fmt.Println(parse.GetFoundStrings())
}

func scanFolder() []string {
	Scan := scanning.Scanning{}

	// set scaning folder
	Scan.SetFolder("testdata")

	// set allowed extensions
	Scan.SetAllowedFiles([]string{".blade.php", ".jsx"})

	// set not allowed files and folder names
	Scan.SetDeniedFilesAndFolders([]string{})

	return Scan.Run()
}
