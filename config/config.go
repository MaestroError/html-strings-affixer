package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	// Scanning
	Folder_to_scan           string   `json:"folder"` // both | just folder to scan
	Allowed_file_types       []string `json:"file_types"` // both | parses file only with given extensions
	Ignore_files_and_folders []string `json:"ignore_names"` // only json | ignores files and folders with given names
	one_file                 string   // only CLI | full path to file, if set, all actions will be performed on one given file 
	// Parse
	Ignore_characters     []string `json:"ignore"` // only json | ignores strings which contains given characters
	Warning_characters     []string `json:"warnings"` // only json | Warns about strings which contains given characters (not replaces)
	Allowed_parse_methods []string `json:"methods"` // both | Uses only given parse methods. Available: text, placeholder, alt, title, hashtag
	// Replace
	Prefix_to_set string `json:"prefix"` // both | Prefix to set
	Suffix_to_set string `json:"suffix"` // both | Suffix to set
	Force bool `json:"force"` // both | If true, git status check is ignored
	// Report
	Detailed_report bool `json:"detailed"` // both | if true, detailed report will be displayed
	Warn_as_table bool `json:"warn_as_Table"` // JSON | if true, warnings will be displayed as table
	// Log
	Log_folder string `json:"log_folder"` // only JSON | folder to store logs
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

func (c *Config) SetWarningCharacters(Warning_characters []string) {
	c.Warning_characters = Warning_characters
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

func (c *Config) SetDetailedReport(detailed bool) {
	c.Detailed_report = detailed
}

func (c *Config) SetForce(force bool) {
	c.Force = force
}

func (c *Config) SetLogFolder(folder string) {
	c.Log_folder = folder
}

func (c *Config) SetWarnAsTable(warn bool) {
	c.Warn_as_table = warn
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

func (c *Config) GetWarningCharacters() []string {
	return c.Warning_characters
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

func (c *Config) GetWarnAsTable() bool {
	return c.Warn_as_table
}

// Defaults

func (c *Config) setDefaults() {
	// check commented function's configs before running commands
	// c.setDefaultFolder()
	// c.setDefaultAllowedFileTypes()
	c.setDefaultIgnoreFilesAndFolders()
	c.setDefaultIgnoreCharacters()
	c.setDefaultWarningCharacters()
	c.setDefaultAllowedParseMethods()
	// c.setDefaultPreffix()
	// c.setDefaultSuffix()
	c.setDefaultDetailed()
	c.setDefaultWarnAsTable()
	c.setDefaultForce()
	c.setDefaultLogFolder()
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
	c.Ignore_characters = []string{"#", "_", ">", "^", "*", "=", "@"}
}

// Sets predefined warning characters
func (c *Config) setDefaultWarningCharacters() {
	c.Warning_characters = []string{"%", "{", "(", "}", ")", "$", `'`}
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

func (c *Config) setDefaultDetailed() {
	c.Detailed_report = false
}

func (c *Config) setDefaultWarnAsTable() {
	c.Warn_as_table = false
}

func (c *Config) setDefaultForce() {
	c.Force = false
}

func (c *Config) setDefaultLogFolder() {
	c.Log_folder = ""
}

func (c *Config) parseJsonFile() {
	jsonFile, err := os.Open("affixer.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err, "| Config file does not used")
	} else {
		fmt.Println("Config file used successfully")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray
	json.Unmarshal(byteValue, &c)
}
