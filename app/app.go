package app

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MaestroError/html-strings-affixer/config"
)

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

var Configuration config.Config

func Bootstrap() {
	// configs and prepare app for work

	// Set config default and read json file
	Configuration.Init()

	// resolve command options and reset configs (command line argument has higher priority)
	resolveCommands()
}

func Start() {
	// Application runtime code goes here
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}

func resolveCommands() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "replace":
			// Replace command
			resolveReplaceCommand()
		case "check":
			// Check command
			resolveReplaceCommand()
		default:
			fmt.Println("Expected 'replace' or 'check' subcommands")
			os.Exit(1)
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
