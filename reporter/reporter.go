package reporter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Reporter struct {
	warnings []string
	errors []string
	report_table table.Writer
}


func (reporter *Reporter) AddWarning(message string) {
	reporter.warnings = append(reporter.warnings, message)
}

func (reporter *Reporter) AddError(message string) {
	reporter.errors = append(reporter.errors, message)
}

func (reporter *Reporter) Report() {

	if reporter.report_table != nil {
		reporter.printTable()
	}

	if reporter.warnings != nil {
		reporter.printWarnings()
	}

	if reporter.errors != nil {
		reporter.printErrors()
	}

}


func (reporter *Reporter) AskForConfirmation(s string, def string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		var warnColor text.Color = text.FgYellow
		msg := s + " y/n ["+def+"]: "
		reporter.print(warnColor, "Confirm: " + msg)
		

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "" {
			response = def
		}

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func (reporter *Reporter) PrintMsg(msg string) {
	var msgColor text.Color = text.FgBlue
	reporter.print(msgColor, "- " + msg)
}

/* Strings */

func (reporter *Reporter) printWarnings() {
	var warnColor text.Color = text.FgHiYellow
	for _, warning := range reporter.warnings {
		reporter.print(warnColor, "Warning: " + warning)
	}
}

func (reporter *Reporter) printErrors() {
	var errorColor text.Color = text.FgHiRed
	for _, msg := range reporter.errors {
		reporter.print(errorColor, "Error: " + msg)
	}
}

func (reporter *Reporter) print(color text.Color, message string) {
	fmt.Println(color.Sprint(message))
}

/* Table */

func (reporter *Reporter) PrepareReplaceTable() {
	t := reporter.getTable()
	t.SetTitle("Affixed strings")
	t.AppendHeader(table.Row{"Location", "Replaced/Found"})
	reporter.report_table = t
}

func (reporter *Reporter) PrepareCheckTable() {
	t := reporter.getTable()
	t.SetTitle("Found strings")
	t.AppendHeader(table.Row{"Location", "Warns/Found"})
	reporter.report_table = t
}

func (reporter *Reporter) AddRow(location string, count string) {
	reporter.report_table.AppendRow(table.Row{location, count})
}

func (reporter *Reporter) AddTotal(total int) {
	reporter.report_table.AppendFooter(table.Row{"Total:", total})
}

func (reporter *Reporter) getTable() table.Writer {
	t := table.NewWriter()
	t.SetAutoIndex(true)
	t.SetStyle(table.StyleRounded)
	return t
}

func (reporter *Reporter) printTable() {
	fmt.Println(reporter.report_table.Render())
}