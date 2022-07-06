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

	fmt.Println(Configuration)
}

func Start() {
	// Application runtime code goes here
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}

func resolveCommands() {

	// Replace command
	// Example: replace -folder=testFolder -only="text,hastag" -allowed="blade.php" -prefix="{{__('" -suffix="')}}"
	replaceCmd := flag.NewFlagSet("replace", flag.ExitOnError)
	replaceFolder := replaceCmd.String("folder", "", "Folder to scan")
	replaceAllowed := replaceCmd.String("allowed", "", "allowed file types, separated by commas")
	replaceMethods := replaceCmd.String("only", "", "Methods to use while parsing, separated by commas. Available: text, placeholder, alt, title, hashtag")
	replacePrefix := replaceCmd.String("prefix", "", "New prefix for strings")
	replaceSuffix := replaceCmd.String("suffix", "", "New suffix for strings")

	// Check command
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)

	switch os.Args[1] {
	case "replace":
		// Replace command resolve
		replaceCmd.Parse(os.Args[2:])
		// set command name
		Configuration.SetCurrentCommand("replace")
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
	case "check":
		// Check command resolve
		checkCmd.Parse(os.Args[2:])
		Configuration.SetCurrentCommand("check")
	default:
		fmt.Println("expected 'replace', 'watch' or 'check' subcommands")
		os.Exit(1)
	}
}

func getCommandExample() {
	// print args
	fmt.Println(len(os.Args), os.Args)

	// creating flags (options)
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	// creating options with different commands
	// command 1
	fooCmd := flag.NewFlagSet("replace", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")
	// command 2
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")
	// exception: one from the commands required
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// check witch command is used with switch
	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
