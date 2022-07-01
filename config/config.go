package config

type Config struct {
	// Scanning
	folder_to_scan           string   // both
	allowed_file_types       []string // both
	ignore_files_and_folders []string // only json
	// Parse
	ignore_characters     []string // only json
	allowed_parse_methods []string // both
	// Replace
	prefix_to_set string // both
	suffix_to_set string // both
}

func (c *Config) Init() {
	c.setDefaults()
}

// Setters

func (c *Config) SetFolderToScan(folder_to_scan string) {
	c.folder_to_scan = folder_to_scan
}

func (c *Config) SetAllowedFileTypes(allowed_file_types []string) {
	c.allowed_file_types = allowed_file_types
}

func (c *Config) SetIgnoreFileTypes(ignore_files_and_folders []string) {
	c.ignore_files_and_folders = ignore_files_and_folders
}

func (c *Config) SetIgnoreCharacters(ignore_characters []string) {
	c.ignore_characters = ignore_characters
}

func (c *Config) SetAllowedParseMethods(allowed_parse_methods []string) {
	c.allowed_parse_methods = allowed_parse_methods
}

func (c *Config) SetPrefixToSet(prefix string) {
	c.prefix_to_set = prefix
}

func (c *Config) SetSuffixToSet(suffix string) {
	c.suffix_to_set = suffix
}

// Getters

func (c *Config) GetFolder() string {
	return c.folder_to_scan
}

func (c *Config) GetFileTypes() []string {
	return c.allowed_file_types
}

func (c *Config) GetIgnoreFiles() []string {
	return c.ignore_files_and_folders
}

func (c *Config) GetIgnoreCharacters() []string {
	return c.ignore_characters
}

func (c *Config) GetAllowedMethods() []string {
	return c.allowed_parse_methods
}

func (c *Config) GetPrefix() string {
	return c.prefix_to_set
}

func (c *Config) GetSuffix() string {
	return c.suffix_to_set
}

// Defaults

func (c *Config) setDefaults() {
	c.setDefaultFolder()
	c.setDefaultAllowedFileTypes()
	c.setDefaultIgnoreFilesAndFolders()
	c.setDefaultIgnoreCharacters()
	c.setDefaultAllowedParseMethods()
	c.setDefaultPreffix()
	c.setDefaultSuffix()
}

func (c *Config) setDefaultFolder() {
	c.folder_to_scan = "resources/views"
}

func (c *Config) setDefaultAllowedFileTypes() {
	c.allowed_file_types = []string{".blade.php", ".jsx", ".vue", ".twig"}
}

func (c *Config) setDefaultIgnoreFilesAndFolders() {
	c.ignore_files_and_folders = []string{}
}

// Sets predefined ignore characters
func (c *Config) setDefaultIgnoreCharacters() {
	c.ignore_characters = []string{"%", "#", "_", ">", "{", "(", "}", ")", "^", "$", "*", "="}
}

// List of available extractions, which starting with "Extract" in parsehtml package
// for example: method for "text" is ExtractText
// used for settings (Allow/deny string types)
func (c *Config) setDefaultAllowedParseMethods() {
	c.allowed_parse_methods = []string{"text", "placeholder", "alt", "title", "hashtag"}
}

func (c *Config) setDefaultPreffix() {
	c.prefix_to_set = "{{ __(\""
}

func (c *Config) setDefaultSuffix() {
	c.suffix_to_set = "\") }}"
}
