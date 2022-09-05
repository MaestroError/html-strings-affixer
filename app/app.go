package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/MaestroError/html-strings-affixer/config"
	"github.com/MaestroError/html-strings-affixer/logger"
	"github.com/MaestroError/html-strings-affixer/parsehtml"
	"github.com/MaestroError/html-strings-affixer/replacer"
	"github.com/MaestroError/html-strings-affixer/reporter"
	"github.com/MaestroError/html-strings-affixer/scanning"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

var Configuration config.Config
var Logger logger.Logger

// configs and prepare app for work
func Bootstrap() {
	// Set config default and read json file
	Configuration.Init()

	// Init logger
	if Configuration.Log_folder != "" {
		Logger.Init(Configuration)
	}

	// resolve command options and reset configs (so command line argument has higher priority)
	resolveCommands()
}

// Application runtime code goes here
func Start() {
	command := Configuration.GetCommandName()
	switch command {
	case "replace":
		// Replace command
		runReplaceCommand()
	case "check":
		// Check command\
		runCheckCommand()
	default:
		fmt.Println("No command found")
		Shutdown()
	}
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
	os.Exit(1)
}

func resolveCommands() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "replace":
			// Replace command
			resolveReplaceCommand()
		case "check":
			// Check command
			resolveCheckCommand()
		case "clear-log":
			// run clear-log command
			runClearLogCommand()
			Shutdown()
		case "debug":
			Configuration.SetCurrentCommand("debug")
			debug()
		default:
			fmt.Println("Expected 'replace', 'clear-log' or 'check' subcommands")
			Shutdown()
		}
	}
}

/* Commands */

func resolveReplaceCommand() {
	// Example: replace -folder=testFolder -only="text,hastag" -allowed="blade.php" -prefix="{{__('" -suffix="')}}"
	// Create flag set
	Cmd := flag.NewFlagSet("", flag.ExitOnError)
	// set flags (options)
	Folder := Cmd.String("folder", "", "Folder to scan")
	Allowed := Cmd.String("allowed", "", "allowed file types, separated by commas")
	Methods := Cmd.String("only", "", "Methods to use while parsing, separated by commas. Available: text, placeholder, alt, title, hashtag")
	Prefix := Cmd.String("prefix", "", "New prefix for strings")
	Suffix := Cmd.String("suffix", "", "New suffix for strings")
	Detailed := Cmd.Bool("detailed", false, "If true, detailed report printed")
	Force := Cmd.Bool("force", false, "If true, git status check is ignored")
	oneFile := Cmd.String("file", "", "Use this argument to run command only on one file")
	// Parse arguments
	Cmd.Parse(os.Args[2:])
	// set command name
	Configuration.SetCurrentCommand("replace")
	// Set configs
	if *Folder != "" {
		Configuration.SetFolderToScan(*Folder)
	}
	if *Allowed != "" {
		Configuration.SetAllowedFileTypes(strings.Split(*Allowed, ","))
	}
	if *Methods != "" {
		Configuration.SetAllowedParseMethods(strings.Split(*Methods, ","))
	}
	if *Prefix != "" {
		Configuration.SetPrefixToSet(*Prefix)
	}
	if *Suffix != "" {
		Configuration.SetSuffixToSet(*Suffix)
	}
	if *oneFile != "" {
		Configuration.SetOneFile(*oneFile)
	}
	if *Detailed {
		Configuration.SetDetailedReport(*Detailed)
	}
	if *Force {
		Configuration.SetForce(*Force)
	}
}

func runReplaceCommand() {
	// scan folder and get needed files
	var files []string
	oneFile := Configuration.GetOneFile()
	if oneFile != "" {
		files = []string{oneFile}
	} else {
		files = scanFolder()
	}
	
	// Controls replace execution
	replaceAllowed := true
	
	// Prepare reporter
	reporter := reporter.Reporter{}
	reporter.PrepareReplaceTable()
	var totalReplaced int = 0
	
	// Check git status (Warn if need to commit)
	if !checkGitStatus() {
		if !Configuration.Force {
			if !reporter.AskForConfirmation("You have uncommitted changes. Recommended to commit or stash them first. Continue anyway?", "yes") {
				replaceAllowed = false
			}
			reporter.PrintMsg("You have uncommitted changes, you can use -force parameter in cli or set 'force' to true in config file to avoid confirmation)")
		} else {
			reporter.PrintMsg("You have uncommitted changes")
		}
	}

	if replaceAllowed {
		for _, path := range files {
			parse := parsehtml.Parsehtml{}
			// path = createTestFile(path)
			parse.ParseFile(path, Configuration)
			affix(path, &parse, &reporter, &totalReplaced)
		}
		// add total count
		reporter.AddTotal(totalReplaced)
	}

	// Report
	reporter.Report()
	// Log
	Logger.Log()
}

func resolveCheckCommand() {
	// Create flag set
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	// set flags (options)
	folder := checkCmd.String("folder", "", "Folder to scan")
	allowed := checkCmd.String("allowed", "", "allowed file types, separated by commas")
	methods := checkCmd.String("only", "", "Methods to use while parsing, separated by commas. Available: text, placeholder, alt, title, hashtag")
	oneFile := checkCmd.String("file", "", "Use this argument to run command only on one file")
	// Parse arguments
	checkCmd.Parse(os.Args[2:])
	// Set configs
	Configuration.SetCurrentCommand("check")
	if *folder != "" {
		Configuration.SetFolderToScan(*folder)
	}
	if *allowed != "" {
		Configuration.SetAllowedFileTypes(strings.Split(*allowed, ","))
	}
	if *methods != "" {
		Configuration.SetAllowedParseMethods(strings.Split(*methods, ","))
	}
	if *oneFile != "" {
		Configuration.SetOneFile(*oneFile)
	}
}

func runCheckCommand() {
	// scan folder and get needed files
	var files []string
	oneFile := Configuration.GetOneFile()
	if oneFile != "" {
		files = []string{oneFile}
	} else {
		files = scanFolder()
	}
	
	// Prepare reporter
	reporter := reporter.Reporter{}
	reporter.PrepareCheckTable()

	for _, path := range files {
		parse := parsehtml.Parsehtml{}
		// path = createTestFile(path)
		parse.ParseFile(path, Configuration)
		data := parse.GetFoundStrings()["data"]

		var countWarnings int
		for _, element := range data {
			// Get warning characters if exists
			foundChars := checkForWarningChars(element["found"])

			if element["type"] == "placeholder" {
				if !checkForPlaceholder(element, path, &reporter) {
					countWarnings++
				}
			}

			// Add warning messages if warning characters found
			if len(foundChars) > 0 {
				countWarnings++
				msg := "Couldn't affix, found warning characters: '" + element["found"] + "' -> " + path + ":"+ element["lines"]
				reporter.AddWarning(msg)
			}
		}

		reporter.AddRow(path, strconv.Itoa(countWarnings) + "/" + strconv.Itoa(len(data)))
	}

	// Report
	reporter.Report()
}

func runClearLogCommand() {
	reporter := reporter.Reporter{}
	if reporter.AskForConfirmation("Are you sure to clear all logs?", "no") {
		if Configuration.Log_folder != "" {
			if Logger.ClearLogs() {
				// @todo add success message function for reporter
				reporter.PrintMsg("All logs removed successfully!")
			} else {
				reporter.AddError("Logs didn't remove")
			}
		}
	}
	reporter.Report()
}

/* Actions */

func scanFolder() []string {
	Scan := scanning.Scanning{}

	// set scaning folder
	Scan.SetFolder(Configuration.Folder_to_scan)

	// set allowed extensions
	Scan.SetAllowedFiles(Configuration.Allowed_file_types)

	// set not allowed files and folder names
	Scan.SetDeniedFilesAndFolders(Configuration.Ignore_files_and_folders)

	return Scan.Run()
}

func affix(path string, parser *parsehtml.Parsehtml, reporter *reporter.Reporter, totalReplaced *int) {
	// get replacement data from parsers
	data := parser.GetFoundStrings()["data"]

	// get file content from parser to not re-read
	var content string = parser.GetOriginalContent()

	// Declare counts (found in file and replaced with success)
	var countInFile int = 0
	var countReplaced int = 0

	// prepare replacer object for use
	affixer := replacer.Replacer{}
	affixer.SetPath(path).SetContent(content)
	affixer.SetPrefix(Configuration.Prefix_to_set).SetSuffix(Configuration.Suffix_to_set)
	affixer.SetDebug(Configuration.GetCommandName() == "debug")

	// loop on found strings
	for _, element := range data {
		// EXTRA CHECKS
		approved := true
		// if placeholder attribute contains only "placeholder"
		if element["type"] == "placeholder" {
			approved = checkForPlaceholder(element, path, reporter)
		}

		// Get warning characters if exists
		foundChars := checkForWarningChars(element["found"])

		// @todo make this warnings printed as table
		// Add warning messages if warning characters found
		if len(foundChars) > 0 {
			approved = false
			msg := "Couldn't affix, found warning characters: '" + element["found"] + "' -> " + path + ":"+ element["lines"]
			reporter.AddWarning(msg)
		}

		countInFile++
		replaced := false

		// affix found string if all checks passed well
		if approved {
			replaced = affixer.Affix(element, parser)
			countReplaced++
		}

		// Add error message if string is approved but not replaced
		if !replaced && approved {
			msg := "String '" + element["found"] + "' not found in file " + path + " (Lines: "+ element["lines"] + ") "
			reporter.AddError(msg)
		}

		// Add string data in detailed report if replaced with success and detailed report requested
		if Configuration.Detailed_report && replaced {
			filePath := path
			if element["lines"] != "" {
				filePath = path + ":" + element["lines"]
			}
			reporter.AddRow(filePath, strings.TrimSpace(element["found"]))
		}

		// add file data in logger if log folder is set
		if Logger.GetLogFolder() != "" {
			Logger.AddNewItem(path, element)
		}
	}
	
	// Msg for developer about disabled logger in this run
	if Logger.GetLogFolder() == "" {
		reporter.PrintMsg("Logger disabled (You will not able to check/undo this changes")
	}
	
	// Set total for report
	*totalReplaced = *totalReplaced + countReplaced
	// Make report string r/f format
	replacedString := strconv.Itoa(countReplaced) + "/" + strconv.Itoa(countInFile)

	// add count if detailed report not requested
	if !Configuration.Detailed_report {
		reporter.AddRow(path, replacedString)
	}
	
	// write file with same name
	err := ioutil.WriteFile(path, []byte(affixer.GetContent()), 0)
	if err != nil {
		panic(err)
	}

}

/* Helpers */

// nothing to commit, working tree clean
func checkGitStatus() bool {
	out, err := exec.Command("git", "status").Output()
	if err != nil {
		panic(err)
	}
	asString := string(out)
	if strings.Contains(asString, "nothing to commit, working tree clean") {
		return true
	} else {
		return false
	}
}

func checkForWarningChars(found string) []string {
	// Check for warning characters existence
	warnings := Configuration.GetWarningCharacters()
	var foundChars []string
	for _, char := range warnings {
		if strings.Contains(found, char) {
			foundChars = append(foundChars, char)
		}
	}
	return foundChars
}

func checkForPlaceholder(element map[string]string, path string, reporter *reporter.Reporter) bool {
	// if placeholder attribute contains only "placeholder"
	if strings.ToLower(element["found"]) == "placeholder" {
		msg := "Couldn't affix, use of 'placeholder' in placeholder attribute not allowed: " + path + ":"+ element["lines"]
		reporter.AddWarning(msg)
		return false
	}
	return true
}

/* Debug */

func debug() {
	debugReplace()
}

func debugReplace() {
	parse := parsehtml.Parsehtml{}
	path := createTestFile("test.blade.php")
	parse.ParseFile(path, Configuration)
	reporter := reporter.Reporter{}		
	reporter.PrepareReplaceTable()
	var totalReplaced int = 0
	affix(path, &parse, &reporter, &totalReplaced)

	reporter.AddTotal(totalReplaced)
	// Report
	reporter.Report()
	// Log
	PrettyPrint(Logger.GetLogData())
}

func createTestFile(path string) string {
	var r []byte
	var err error
	r, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	createDirIfNotExists("testdata\\replaceTests")
	content := string(r)
	rand.Seed(time.Now().UnixNano())
	randomNumber := strconv.Itoa(rand.Intn(100))
	path = strings.ReplaceAll(path, "\\", "-")
	path = strings.ReplaceAll(path, "/", "-")
	newFileName := "testdata\\replaceTests\\" + path + "_" + randomNumber + ".blade.php"
	err = ioutil.WriteFile(newFileName, []byte(content), 0775)
	if err != nil {
		panic(err)
	}
	return newFileName
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func createDirIfNotExists(path string) {
	exists, err := exists(path)
	if err != nil {
		panic(err)
	}
	if !exists {
		os.Mkdir(path, 0777)
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}