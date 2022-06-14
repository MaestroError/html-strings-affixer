package parsehtml

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Parsehtml struct {
	file              string
	found_strings     map[string][]map[string]string
	content           string
	original_content  string
	ignore_characters []string
	prefix            string
	suffix            string
	regexp            *regexp.Regexp
	search_regex      string
}

/*
* @todo make Get file content funtion and use it in constructor (Init) +
* @todo Simple string extraction function and Test it +
* @todo line extraction function from found string (consider dublicate strings) +
* @todo Comment everything what you done, consider all logical parts !!!
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

// type -> string describing type of visible html
func (parse *Parsehtml) AddNewString(found string, need_to_replace string, found_type string, lines string) {
	foundObject := make(map[string]string)
	foundObject["found"] = found
	foundObject["need_to_replace"] = need_to_replace
	foundObject["type"] = found_type
	foundObject["lines"] = lines
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

func (parse *Parsehtml) ExtractText() {
	// set affixes for simple strings extraction
	parse.SetPrefix("\\>")
	parse.SetSuffix("\\<")
	parse.generateRegex()
	parse.parseContent("text")
}

// privates
func (parse *Parsehtml) setFoundStrings(found_strings map[string][]map[string]string) {
	parse.found_strings = found_strings
}

func (parse *Parsehtml) renewContent() {
	parse.content = parse.original_content
}

func (parse *Parsehtml) findLineOfString(str string) []string {
	f, err := os.Open(parse.file)
	if err != nil {
		// return 0, err
		panic(err)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	foundOnLines := []string{}
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), str) {
			foundOnLines = append(foundOnLines, strconv.Itoa(line))
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
		panic(err)
	}
	return foundOnLines
}

func (parse *Parsehtml) getFileContent() {
	var r []byte
	var err error
	r, err = ioutil.ReadFile(parse.file)
	if err != nil {
		panic(err)
	}
	content := string(r)
	parse.content = content
	parse.original_content = content
}

func (parse *Parsehtml) setIgnoreCharacters() {
	parse.ignore_characters = []string{"%", "#", "_", ">", "{", "(", "}", ")"}
}

func (parse *Parsehtml) generateRegex() {
	if parse.prefix != "" && parse.suffix != "" {
		deniedCharString := strings.Join(parse.ignore_characters, "\\")
		// [^\s+] -> used to not match whitespace
		reg := regexp.MustCompile(parse.prefix + `[^` + deniedCharString + `].[^\s+][^` + deniedCharString + `]+` + parse.suffix)
		parse.search_regex = reg.String()
		parse.regexp = reg
	}
}

func (parse *Parsehtml) parseContent(htmlType string) {
	submatchall := parse.regexp.FindAllString(parse.content, -1)
	for _, element := range submatchall {
		found := strings.Trim(element, parse.prefix)
		found = strings.Trim(found, parse.suffix)
		if !parse.checkDuplicate(found) {
			lines := parse.findLineOfString(found)
			parse.AddNewString(found, element, htmlType, strings.Join(lines, ", "))
		}
	}
}

func (parse *Parsehtml) checkDuplicate(found string) bool {
	result := false
	for _, fs := range parse.found_strings[parse.file] {
		if fs["found"] == found {
			result = true
			break
		}
	}
	return result
}
