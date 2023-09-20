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
	warnings       []string
	errors         []string
	success        []string
	report_table   table.Writer
	warnAsTable    bool
	warnings_table table.Writer
}

func (reporter *Reporter) AddWarning(message string) {
	reporter.warnings = append(reporter.warnings, message)
}

func (reporter *Reporter) AddError(message string) {
	reporter.errors = append(reporter.errors, message)
}

func (reporter *Reporter) AddSuccess(message string) {
	reporter.success = append(reporter.success, message)
}

func (reporter *Reporter) WarnAsTable() {
	reporter.warnAsTable = true
}

func (reporter *Reporter) Report() {

	if reporter.report_table != nil {
		reporter.printTable()
	}

	if reporter.warnings != nil {
		if reporter.warnAsTable {
			reporter.printWarningsTable()
		} else {
			reporter.printWarnings()
		}
	}

	if reporter.errors != nil {
		reporter.printErrors()
	}

	if reporter.success != nil {
		reporter.printSuccesses()
	}

}

func (reporter *Reporter) AskForConfirmation(s string, def string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		var warnColor text.Color = text.FgYellow
		msg := s + " y/n [" + def + "]: "
		reporter.print(warnColor, "Confirm: "+msg)

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

func (reporter *Reporter) PrintMsg(msg ...string) {
	var msgColor text.Color = text.FgBlue
	for _, m := range msg {
		reporter.print(msgColor, "- "+m)
	}
}

func (reporter *Reporter) PrintSuccess(msg ...string) {
	var msgColor text.Color = text.FgGreen
	for _, m := range msg {
		reporter.print(msgColor, "- "+m)
	}
}

/* Strings */

func (reporter *Reporter) printWarnings() {
	var warnColor text.Color = text.FgHiYellow
	for _, warning := range reporter.warnings {
		reporter.print(warnColor, "- Warning: "+warning+"\n -----")
	}
}

func (reporter *Reporter) printWarningsTable() {
	reporter.prepareWarningsTable()
	for _, warning := range reporter.warnings {
		warn := strings.Split(warning, "->")
		if len(warn) > 1 {
			str := strings.Split(warn[0], ":")
			reporter.warnings_table.AppendRow(table.Row{str[1], warn[1]})
		} else {
			reporter.warnings_table.AppendRow(table.Row{warn[0], ""})
		}
	}
	fmt.Println(reporter.warnings_table.Render())
}

func (reporter *Reporter) printErrors() {
	var errorColor text.Color = text.FgHiRed
	for _, msg := range reporter.errors {
		reporter.print(errorColor, "Error: "+msg)
	}
}

func (reporter *Reporter) printSuccesses() {
	var sucColor text.Color = text.FgHiGreen
	for _, msg := range reporter.success {
		reporter.print(sucColor, "Success: "+msg)
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

func (reporter *Reporter) prepareWarningsTable() {
	t := reporter.getTable()
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Warning on:", "location"})
	reporter.warnings_table = t
}

// Getters
func (reporter *Reporter) GetWarnings() []string {
	return reporter.warnings
}
