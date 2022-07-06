package parsehtml

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/MaestroError/html-strings-affixer/config"
)

type Parsehtml struct {
	file             string
	found_strings    map[string][]map[string]string
	content          string
	original_content string
	// options
	ignore_characters []string
	extractions       []string
	// Affixes to search string
	prefix string
	suffix string
	// regex
	regexp       *regexp.Regexp
	search_regex string
}

/*
* @todo global extract method, with configs and extraction methods check
 */

func (parse *Parsehtml) Init(file string, c config.Config) {
	parse.found_strings = make(map[string][]map[string]string)
	parse.SetFile(file)
	parse.getFileContent()

	// set options from config
	parse.setIgnoreCharacters(c.GetIgnoreCharacters())
	// @todo check if available, while extracting
	parse.setExtractions(c.GetAllowedMethods())
}

// setters
func (parse *Parsehtml) SetFile(file string) {
	parse.file = file
}

// Adds new string in found_strings
// sets trimmed string as "found" and original string as "original_string"
// type -> string describing type of visible html, you can specify it while calling parse.parseContent method
// lines -> lines where found string exists, you can get it with parse.findLineOfString method
func (parse *Parsehtml) AddNewString(found string, original_string string, found_type string, lines string) {
	foundObject := make(map[string]string)
	foundObject["found"] = found
	foundObject["original_string"] = original_string
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

// Simple strings extraction method - just plain strings in HTML
func (parse *Parsehtml) ExtractText() {
	// set affixes for simple strings extraction
	parse.SetPrefix("\\>")
	parse.SetSuffix("\\<")
	// Generates regex based on prefix, suffix and denied characters
	parse.generateRegex()
	// Parses content and adds strings in found_strings with specific type
	parse.parseContent("text")
}

// HTML input's Placeholders attributes extraction method
// XX - Can't use word "placeholder" inside placeholder - XX ?? why? it does well
func (parse *Parsehtml) ExtractPlaceholder() {
	// set affixes for simple strings extraction
	// (?i) = case insensitive
	parse.SetPrefix("(?i)placeholder=(\"|')")
	parse.SetSuffix("(\"|')")
	// Generates regex based on prefix, suffix and denied characters
	parse.generateRegex()
	// Parses content and adds strings in found_strings with specific type
	parse.parseContent("placeholder")
}

// HTML img's alt attributes extraction method
func (parse *Parsehtml) ExtractAlt() {
	// set affixes for simple strings extraction
	parse.SetPrefix("(?i)alt=(\"|')")
	parse.SetSuffix("(\"|')")
	// Generates regex based on prefix, suffix and denied characters
	parse.generateRegex()
	// Parses content and adds strings in found_strings with specific type
	parse.parseContent("alt")
}

// HTML title attributes extraction method
func (parse *Parsehtml) ExtractTitle() {
	// set affixes for simple strings extraction
	parse.SetPrefix("(?i)title=(\"|')")
	parse.SetSuffix("(\"|')")
	// Generates regex based on prefix, suffix and denied characters
	parse.generateRegex()
	// Parses content and adds strings in found_strings with specific type
	parse.parseContent("title")
}

// Extracts "#text" type (selected) strings
func (parse *Parsehtml) ExtractHashtag() {
	// set affixes for simple strings extraction
	parse.SetPrefix("(\"|'|>)\\s*#")
	parse.SetSuffix("(\"|'|<)")
	// Generates regex based on prefix, suffix and denied characters
	parse.generateRegex()
	// Parses content and adds strings in found_strings with specific type
	// @todo add "#" as strip to remove it while replacing
	parse.parseContent("hashtag")

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
	// check each line for founded string existence
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), str) {
			// append line as string in foundOnLines array
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

// Reads file and sets content (as content and original_content properties)
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

func (parse *Parsehtml) setIgnoreCharacters(ignore_characters []string) {
	parse.ignore_characters = ignore_characters
}

func (parse *Parsehtml) setExtractions(allowed_parse_methods []string) {
	parse.extractions = allowed_parse_methods
}

// Generates regex based on prefix, suffix and denied characters
// sets search_regex as regular expression string
// and regexp as regexp object
func (parse *Parsehtml) generateRegex() {
	if parse.prefix != "" && parse.suffix != "" {
		deniedCharString := strings.Join(parse.ignore_characters, "\\")
		// [^\s+] -> used to not match whitespace
		reg := regexp.MustCompile(parse.prefix + `[^` + deniedCharString + `].[^\s+][^` + deniedCharString + `]+` + parse.suffix)
		parse.search_regex = reg.String()
		parse.regexp = reg
	}
}

// parses content, trims found strings and adds in found_strings if not already exists
func (parse *Parsehtml) parseContent(htmlType string) {
	// find all strings based on regex
	submatchall := parse.regexp.FindAllString(parse.content, -1)
	for _, element := range submatchall {
		// removes (trims) finding prefix and suffix
		re := regexp.MustCompile(parse.prefix)
		found := re.ReplaceAllString(element, "")
		re = regexp.MustCompile(parse.suffix)
		found = re.ReplaceAllString(found, "")
		// add as new string if no duplicates found
		if !parse.checkDuplicate(found) {
			lines := parse.findLineOfString(found)
			parse.AddNewString(found, element, htmlType, strings.Join(lines, ", "))
		}
	}
}

// check if string already exists in found strings
func (parse *Parsehtml) checkDuplicate(found string) bool {
	result := false
	// @todo check also type of string or "original_string" (maybe some string will need different methods to replace)
	for _, fs := range parse.found_strings[parse.file] {
		if fs["found"] == found {
			result = true
			break
		}
	}
	return result
}
