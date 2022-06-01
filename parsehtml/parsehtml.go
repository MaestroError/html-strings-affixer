package parsehtml

type ParseHtml struct {
	file          string
	found_strings map[string][]map[string]string
}

func (parse *ParseHtml) setFile(file string) {
	parse.file = file
}

func (parse *ParseHtml) setFoundStrings(found_strings map[string][]map[string]string) {
	parse.found_strings = found_strings
}

func (parse *ParseHtml) addNewString(found string, need_to_replace string) {
	foundObject := make(map[string]string)
	foundObject["found"] = found
	foundObject["need_to_replace"] = need_to_replace
	parse.found_strings[parse.file] = append(parse.found_strings[parse.file], foundObject)
}
