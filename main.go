package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/MaestroError/html-strings-affixer/app"
	"github.com/MaestroError/html-strings-affixer/parsehtml"
	"github.com/MaestroError/html-strings-affixer/scanning"
)

// "testdata", []string{"index.blade.php", "denyThisFolder"}, []string{".blade.php", ".jsx"}
func main() {

	// @todo Create config package and Config struct:
	//			- parse configs from json file +
	//			- parse from CLI +
	//			- make setter and getter methods +
	//			- inject config in other structs (parsehtml, scanning)

	app.Bootstrap()

	fmt.Println(app.Configuration)

	// files := scanFolder()
	// fmt.Println(files)

	parse := parsehtml.Parsehtml{}
	data := parse.ParseFile("testdata\\pages\\test.blade.php", app.Configuration)
	PrettyPrint(data.GetFoundStrings())

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

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// deprecated function
func testExtraction() {

	// Regex
	var r []byte
	var err1 error
	r, err1 = ioutil.ReadFile("C:\\Users\\XPS\\Desktop\\html-strings-affixer\\testdata\\pages\\test.blade.php")
	if err1 != nil {
		panic(err1)
	}
	html := string(r)

	deniedCharacters := []string{"%", "#", "_", ">", "{", "(", "}", ")"}
	deniedCharString := strings.Join(deniedCharacters, "\\")

	// ("|') OR condition with |
	// prefix := `placeholder=("|')`
	// subfix := `("|')`
	prefix := `\>`
	subfix := `\<\/`

	re := regexp.MustCompile(prefix + `[^` + deniedCharString + `].[^` + deniedCharString + `]+` + subfix)
	fmt.Printf("Pattern: %v\n", re.String()) // print pattern

	fmt.Println("\nText between parentheses:")
	submatchall := re.FindAllString(html, -1)
	fmt.Println(submatchall)
	for _, element := range submatchall {
		element = strings.Trim(element, ">")
		element = strings.Trim(element, "</")
		fmt.Println(element)
	}
}
