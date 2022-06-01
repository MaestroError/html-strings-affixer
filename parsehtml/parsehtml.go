package parsehtml

type ParseHtml struct {
	file          string
	found_strings map[string][]map[string]string
	content       string
}

/*
* @todo make Get file content funtion and use it in constructor (Init)
* @todo Simple string extraction function
* @todo craete list of Visible HTML attributes
* @todo methods for parsing Visible HTML attributes
* @todo How to catch input type submit's value attribute?
* @todo Parsing class: Remember replaceable strings with pairs, per file ">Some nice string</"
* @todo Parsing class: Replace Prefix and Suffix
* @todo Parsing class: Replace function, which uses Prefix-Suffix and replaceable string pairs per file one by one
 */

func (parse *ParseHtml) Init(file string) {
	parse.found_strings = make(map[string][]map[string]string)
	parse.SetFile(file)
}

// setters
func (parse *ParseHtml) SetFile(file string) {
	parse.file = file
}

func (parse *ParseHtml) AddNewString(found string, need_to_replace string) {
	foundObject := make(map[string]string)
	foundObject["found"] = found
	foundObject["need_to_replace"] = need_to_replace
	foundObject["type"] = need_to_replace
	parse.found_strings[parse.file] = append(parse.found_strings[parse.file], foundObject)
}

func (parse *ParseHtml) GetFoundStrings() map[string][]map[string]string {
	return parse.found_strings
}

// privates
func (parse *ParseHtml) setFoundStrings(found_strings map[string][]map[string]string) {
	parse.found_strings = found_strings
}
