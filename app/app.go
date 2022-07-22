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
	"github.com/jedib0t/go-pretty/v6/table"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

var Configuration config.Config
var Logger logger.Logger

// configs and prepare app for work
func Bootstrap() {
	// Set config default and read json file
	Configuration.Init()
	Logger.Init(Configuration)


	// resolve command options and reset configs (command line argument has higher priority)
	resolveCommands()
}

// Application runtime code goes here
func Start() {
	command := Configuration.GetCommandName()
	switch command {
	case "replace":
		// Replace command
		// scan folder and get needed files
		files := scanFolder()

		// Controls replace execution
		replaceAllowed := true
		
		// Prepare reporter
		reporter := reporter.Reporter{}
		reporter.PrepareReplaceTable()
		var totalReplaced int = 0
		
		// Check git status (Warn if need to commit)
		if !checkGitStatus() {
			replaceAllowed = false
			// @todo add -force parameter in replace command and replace error with warning if -force is true
			reporter.AddError("You have uncommitted changes, please commit/stash them or run command again with -force parameter (git status)")
		}

		if replaceAllowed {
			for _, path := range files {
				parse := parsehtml.Parsehtml{}
				// path = createTestFile(path)
				parse.ParseFile(path, Configuration)
				Replace(path, &parse, &reporter, &totalReplaced)
			}
			// add total count
			reporter.AddTotal(totalReplaced)
		}

		// Report
		reporter.Report()
		// Log
		Logger.Log()

	case "check":
		// Check command\
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
		case "debug":
			Configuration.SetCurrentCommand("debug")
			debug()
		default:
			fmt.Println("Expected 'replace' or 'check' subcommands")
			Shutdown()
		}
	}
}

func resolveReplaceCommand() {
	// Example: replace -folder=testFolder -only="text,hastag" -allowed="blade.php" -prefix="{{__('" -suffix="')}}"
	replaceCmd := flag.NewFlagSet("replace", flag.ExitOnError)
	replaceFolder := replaceCmd.String("folder", "", "Folder to scan")
	replaceAllowed := replaceCmd.String("allowed", "", "allowed file types, separated by commas")
	replaceMethods := replaceCmd.String("only", "", "Methods to use while parsing, separated by commas. Available: text, placeholder, alt, title, hashtag")
	replacePrefix := replaceCmd.String("prefix", "", "New prefix for strings")
	replaceSuffix := replaceCmd.String("suffix", "", "New suffix for strings")
	replaceDetailed := replaceCmd.Bool("detailed", true, "If true, detailed report printed")
	oneFile := replaceCmd.String("file", "", "Use this argument to run command only on one file")
	// Parse arguments
	replaceCmd.Parse(os.Args[2:])
	// set command name
	Configuration.SetCurrentCommand("replace")
	// Set configs
	if *replaceFolder != "" {
		Configuration.SetFolderToScan(*replaceFolder)
	}
	if *replaceAllowed != "" {
		Configuration.SetAllowedFileTypes(strings.Split(*replaceAllowed, ","))
	}
	if *replaceMethods != "" {
		Configuration.SetAllowedParseMethods(strings.Split(*replaceMethods, ","))
	}
	if *replacePrefix != "" {
		Configuration.SetPrefixToSet(*replacePrefix)
	}
	if *replaceSuffix != "" {
		Configuration.SetSuffixToSet(*replaceSuffix)
	}
	if *oneFile != "" {
		Configuration.SetOneFile(*oneFile)
	}
	if *replaceDetailed {
		Configuration.SetDetailedReport(*replaceDetailed)
	}
}

func resolveCheckCommand() {
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
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

/*
* @todo Create structs with following actions:
* 		- Affixer +
*		- Logger (in json file)
*		- Backup (Zips original content of file )
* 		- Reporter (results in CLI) +
 */
func Replace(path string, parser *parsehtml.Parsehtml, reporter *reporter.Reporter, totalReplaced *int) {
	// get replacement data from parsers
	data := parser.GetFoundStrings()["data"]

	// get file content from parser to not re-read
	var content string = parser.GetOriginalContent()

	// Declare counts
	var countInFile int = 0
	var countReplaced int = 0

	// prepare replacer object for use
	affixer := replacer.Replacer{}
	affixer.SetPath(path).SetContent(content)
	affixer.SetPrefix(Configuration.Prefix_to_set).SetSuffix(Configuration.Suffix_to_set)
	affixer.SetDebug(Configuration.GetCommandName() == "debug")

	// loop on found strings
	for _, element := range data {
		// Extra checks
		approved := true
		// if placeholder attribute contains only "placeholder"
		if element["type"] == "placeholder" {
			if strings.ToLower(element["found"]) == "placeholder" {
				approved = false
				msg := "Couldn't affix, use of 'placeholder' in placeholder attribute not allowed: " + path + ":"+ element["lines"]
				reporter.AddWarning(msg)
			}
		}

		countInFile++
		replaced := false
		// affix found string if all checks passed well
		if approved {
			replaced = affixer.Affix(element, parser)
			countReplaced++
		}

		if !replaced && approved {
			msg := "String '" + element["found"] + "' not found in file " + path + " (Lines: "+ element["lines"] + ") "
			reporter.AddError(msg)
		}

		if Configuration.Detailed_report {
			reporter.AddRow(path + ":" + element["lines"], element["found"])
		}

		if Logger.GetLogFolder() != "" {
			Logger.AddNewItem(path, element)
		}
			
		
	}
	
	if Logger.GetLogFolder() == "" {
		// @todo add msg (before table, gray) property and related methods in reporter, use here message 
		reporter.AddWarning("Logger disabled")
	}
	
	*totalReplaced = *totalReplaced + countReplaced
	replacedString := strconv.Itoa(countReplaced) + "/" + strconv.Itoa(countInFile)

	if !Configuration.Detailed_report {
		reporter.AddRow(path, replacedString)
	}
	
	
	// write file with same name
	err := ioutil.WriteFile(path, []byte(affixer.GetContent()), 0)
	if err != nil {
		panic(err)
	}

}

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

/* Testing */

func debug() {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Location", "Found", "Replaced"})
	t.AppendRow(table.Row{1, "testdata/pages/test.blade.php:20", "Stark", 3000})
	// all rows need not have the same number of columns
	t.AppendRow(table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"})
	// t.AppendSeparator()
	// table.Row is just a shorthand for []interface{}
	t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// time to take a peek
	t.SetCaption("Simple Table with 3 Rows.\n")
	t.SetStyle(table.StyleRounded)
	t.SetTitle("Game Of Thrones")
	fmt.Println(t.Render())

	debugReplace()
}

func debugReplace() {
	parse := parsehtml.Parsehtml{}
	path := createTestFile("test.blade.php")
	parse.ParseFile(path, Configuration)
	reporter := reporter.Reporter{}		
	reporter.PrepareReplaceTable()
	var totalReplaced int = 0
	Replace(path, &parse, &reporter, &totalReplaced)

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
	if exists == false {
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