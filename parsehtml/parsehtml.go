package parsehtml

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

type Parsehtml struct {
	file              string
	found_strings     map[string][]map[string]string
	content           string
	ignore_characters []string
	prefix            string
	suffix            string
	search_regex      string
}

/*
* @todo make Get file content funtion and use it in constructor (Init) +
* @todo Simple string extraction function
* @todo line extraction function from found string (consider dublicate strings)
* @todo craete list of Visible HTML attributes
* @todo methods for parsing Visible HTML attributes
* @todo How to catch input type submit's value attribute?
* @todo Parsing class: Remember replaceable strings with pairs, per file ">Some nice string</"
* @todo Parsing class: Replace Prefix and Suffix
* @todo Parsing class: Replace function, which uses Prefix-Suffix and replaceable string pairs per file one by one
* @todo Tests for Parsehtml
 */

func (parse *Parsehtml) Init(file string) {
	parse.found_strings = make(map[string][]map[string]string)
	parse.SetFile(file)
	parse.getFileContent()
	parse.setIgnoreCharacters()
}

// setters
func (parse *Parsehtml) SetFile(file string) {
	parse.file = file
}

func (parse *Parsehtml) AddNewString(found string, need_to_replace string) {
	foundObject := make(map[string]string)
	foundObject["found"] = found
	foundObject["need_to_replace"] = need_to_replace
	foundObject["type"] = need_to_replace
	parse.found_strings[parse.file] = append(parse.found_strings[parse.file], foundObject)
}

func (parse *Parsehtml) GetFoundStrings() map[string][]map[string]string {
	return parse.found_strings
}

func (parse *Parsehtml) AddIgnoreCharacter(char string) {
	parse.ignore_characters = append(parse.ignore_characters, char)
}

func (parse *Parsehtml) SetPrefix(prefix string) {
	parse.prefix = prefix
}

func (parse *Parsehtml) SetSuffix(suffix string) {
	parse.suffix = suffix
}

// privates
func (parse *Parsehtml) setFoundStrings(found_strings map[string][]map[string]string) {
	parse.found_strings = found_strings
}

func (parse *Parsehtml) findLineOfString(str string) (int, error) {
	f, err := os.Open(parse.file)
	if err != nil {
		// return 0, err
		panic(err)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), str) {
			return line, nil
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
		panic(err)
	}
	return line, err
}

func (parse *Parsehtml) getFileContent() {
	var r []byte
	var err error
	r, err = ioutil.ReadFile(parse.file)
	if err != nil {
		panic(err)
	}
	parse.content = string(r)
}

func (parse *Parsehtml) setIgnoreCharacters() {
	parse.ignore_characters = []string{"%", "#", "_", ">", "{", "(", "}", ")"}
}

func (parse *Parsehtml) generateRegex() {

}
