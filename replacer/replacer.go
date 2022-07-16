package replacer

import (
	"fmt"
	"strings"

	"github.com/MaestroError/html-strings-affixer/backup"
	"github.com/MaestroError/html-strings-affixer/logger"
)

type Replacer struct {
	path    string
	content string
	prefix  string
	suffix  string
	logger  logger.Logger
	backup  backup.Backup
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

func (r *Replacer) SetLogger(logger logger.Logger) *Replacer {
	r.logger = logger
	return r
}

func (r *Replacer) SetBackup(backup backup.Backup) *Replacer {
	r.backup = backup
	return r
}

func (r *Replacer) Affix(original_string string, found string) {
	str := original_string
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
		fmt.Println("Replaced: ", found, startIndex, endIndex)
	} else {
		fmt.Println("Not replaced: ", found, startIndex)
	}
}

func (r *Replacer) GetContent() string {
	return r.content
}
