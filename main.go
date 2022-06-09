package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

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

	// Regex
	var r []byte
	var err1 error
	r, err1 = ioutil.ReadFile("C:\\Users\\XPS\\Desktop\\html-strings-affixer\\testdata\\pages\\test.blade.php")
	if err1 != nil {
		panic(err1)
	}
	html := string(r)

	// @todo add more denied characters
	// @todo make it as package with struct and changeable configs
	// @todo add more check typs (pre-subfix)
	deniedCharacters := []string{"%", "#", "_", ">", "{", "(", "}", ")"}
	deniedCharString := strings.Join(deniedCharacters, "\\")

	prefix := `\>`
	subfix := `\<`

	re := regexp.MustCompile(prefix + `[^` + deniedCharString + `].[^` + deniedCharString + `]+` + subfix)
	fmt.Printf("Pattern: %v\n", re.String()) // print pattern

	fmt.Println("\nText between parentheses:")
	submatchall := re.FindAllString(html, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, ">")
		element = strings.Trim(element, "</")
		fmt.Println(element)
	}

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
