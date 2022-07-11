package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	// Scanning
	Folder_to_scan           string   `json:"folder"` // both
	Allowed_file_types       []string // both
	Ignore_files_and_folders []string // only json
	one_file                 string   // only CLI
	// Parse
	Ignore_characters     []string // only json
	Allowed_parse_methods []string // both
	// Replace
	Prefix_to_set string // both
	Suffix_to_set string // both
	// info
	current_command string
}

func (c *Config) Init() {
	c.setDefaults()
	c.parseJsonFile()
}

// Setters

func (c *Config) SetFolderToScan(Folder_to_scan string) {
	c.Folder_to_scan = Folder_to_scan
}

func (c *Config) SetAllowedFileTypes(Allowed_file_types []string) {
	c.Allowed_file_types = Allowed_file_types
}

func (c *Config) SetIgnoreFileTypes(Ignore_files_and_folders []string) {
	c.Ignore_files_and_folders = Ignore_files_and_folders
}

func (c *Config) SetIgnoreCharacters(Ignore_characters []string) {
	c.Ignore_characters = Ignore_characters
}

func (c *Config) SetAllowedParseMethods(Allowed_parse_methods []string) {
	c.Allowed_parse_methods = Allowed_parse_methods
}

func (c *Config) SetPrefixToSet(prefix string) {
	c.Prefix_to_set = prefix
}

func (c *Config) SetSuffixToSet(suffix string) {
	c.Suffix_to_set = suffix
}

func (c *Config) SetCurrentCommand(command_name string) {
	c.current_command = command_name
}

func (c *Config) SetOneFile(one_file string) {
	c.one_file = one_file
}

// Getters

func (c *Config) GetFolder() string {
	return c.Folder_to_scan
}

func (c *Config) GetFileTypes() []string {
	return c.Allowed_file_types
}

func (c *Config) GetIgnoreFiles() []string {
	return c.Ignore_files_and_folders
}

func (c *Config) GetIgnoreCharacters() []string {
	return c.Ignore_characters
}

func (c *Config) GetAllowedMethods() []string {
	return c.Allowed_parse_methods
}

func (c *Config) GetPrefix() string {
	return c.Prefix_to_set
}

func (c *Config) GetSuffix() string {
	return c.Suffix_to_set
}

func (c *Config) GetCommandName() string {
	return c.current_command
}

func (c *Config) GetOneFile() string {
	return c.one_file
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
	if c.Folder_to_scan == "" {
		c.Folder_to_scan = "resources/views"
	}
}

func (c *Config) setDefaultAllowedFileTypes() {
	c.Allowed_file_types = []string{".blade.php", ".jsx", ".vue", ".twig"}
}

func (c *Config) setDefaultIgnoreFilesAndFolders() {
	c.Ignore_files_and_folders = []string{}
}

// Sets predefined ignore characters
func (c *Config) setDefaultIgnoreCharacters() {
	c.Ignore_characters = []string{"%", "#", "_", ">", "{", "(", "}", ")", "^", "$", "*", "=", "'"}
}

// List of available extractions, which starting with "Extract" in parsehtml package
// for example: method for "text" is ExtractText
// used for settings (Allow/deny string types)
func (c *Config) setDefaultAllowedParseMethods() {
	c.Allowed_parse_methods = []string{"text", "placeholder", "alt", "title", "hashtag"}
}

func (c *Config) setDefaultPreffix() {
	c.Prefix_to_set = "{{ __('"
}

func (c *Config) setDefaultSuffix() {
	c.Suffix_to_set = "') }}"
}

func (c *Config) parseJsonFile() {
	jsonFile, err := os.Open("affixer-config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err, "| Config file does not used")
	}

	fmt.Println("Config file used successfully")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &c)
}
