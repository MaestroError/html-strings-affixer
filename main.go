package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/MaestroError/html-strings-affixer/app"
	"github.com/MaestroError/html-strings-affixer/parsehtml"
)

// "testdata", []string{"index.blade.php", "denyThisFolder"}, []string{".blade.php", ".jsx"}
func main() {

	// @todo - inject config in scanning struct and finish replace command

	app.Bootstrap()

	debug()

}

func debug() {
	parse := parsehtml.Parsehtml{}
	path := "testdata\\test.blade.php"
	parse.ParseFile(path, app.Configuration)
	// PrettyPrint(data)
	app.Replace(path, &parse)

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
