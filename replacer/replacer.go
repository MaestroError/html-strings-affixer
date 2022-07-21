package replacer

import (
	"fmt"
	"strings"

	"github.com/MaestroError/html-strings-affixer/backup"
	"github.com/MaestroError/html-strings-affixer/parsehtml"
)

type Replacer struct {
	path    string
	content string
	prefix  string
	suffix  string
	backup  backup.Backup
	debug bool
}

func (r *Replacer) SetPath(path string) *Replacer {
	r.path = path
	return r
}

func (r *Replacer) SetContent(content string) *Replacer {
	r.content = content
	return r
}

func (r *Replacer) SetPrefix(prefix string) *Replacer {
	r.prefix = prefix
	return r
}

func (r *Replacer) SetSuffix(suffix string) *Replacer {
	r.suffix = suffix
	return r
}

func (r *Replacer) SetBackup(backup backup.Backup) *Replacer {
	r.backup = backup
	return r
}

func (r *Replacer) SetDebug(debug bool) *Replacer {
	r.debug = debug
	return r
}

func (r *Replacer) Affix(element map[string]string, parser *parsehtml.Parsehtml) bool {
	// Set main data
	str := element["original_string"]
	original_string := element["original_string"]
	// @todo Create it controllable with config trimSpaces=true (both json & cli)
	// Trim whitespace from sides
	found := strings.TrimSpace(element["found"])

	// Extra manipulation
	if element["type"] == "hashtag" {
		str = parser.RemoveFirstOccurrence(str, "#")
	}
	// Find prefix index
	startIndex := strings.Index(str, found)
	if startIndex > -1 {
		// set prefix
		str = str[:startIndex] + r.prefix + str[startIndex:]
		// find suffix index from edited string (str)
		endIndex := strings.Index(str, found) + len(found)
		// set suffix
		str = str[:endIndex] + r.suffix + str[endIndex:]
		// replace original with edited string (str)
		r.content = strings.Replace(r.content, original_string, str, -1)

		// print if debug
		if r.debug {
			fmt.Println("Replaced: ", element)
		}
		// Replacement done successfully
		return true
	}
	
	// print if debug
	if r.debug {
		fmt.Println("Not replaced: ", element)
	}

	// it means that "found" not found in "str"
	return false
}

func (r *Replacer) GetContent() string {
	return r.content
}
