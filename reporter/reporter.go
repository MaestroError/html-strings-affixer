package reporter

import (
	"fmt"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Reporter struct {
	data     [][]string
	warnings []string
	errors []string
	detailed bool // if true, printing detailed report (with strings and line)
	report_table table.Writer
}

/*
	@todo make data report (table) and warning report methods
*/


func (reporter *Reporter) AddWarning(message string) {
	reporter.warnings = append(reporter.warnings, message)
}

func (reporter *Reporter) AddError(message string) {
	reporter.errors = append(reporter.errors, message)
}

func (reporter *Reporter) EnableDetailed() {
	reporter.detailed = true
}

func (reporter *Reporter) AddCountData(location string, count int) {
	reporter.data = append(reporter.data, []string{location, strconv.Itoa(count)})
}

func (reporter *Reporter) AddDetailedData(location string, found string) {
	reporter.data = append(reporter.data, []string{location, found})
}

func (reporter *Reporter) Report() {
	// @todo print table here

	if reporter.warnings != nil {
		reporter.printWarnings()
	}

	if reporter.errors != nil {
		reporter.printErrors()
	}
}

/* Strings */

func (reporter *Reporter) printWarnings() {
	var warnColor text.Color = text.FgHiYellow
	for _, warning := range reporter.warnings {
		reporter.print(warnColor, warning)
	}
}

func (reporter *Reporter) printErrors() {
	var errorColor text.Color = text.FgHiRed
	for _, msg := range reporter.errors {
		reporter.print(errorColor, msg)
	}
}

func (reporter *Reporter) print(color text.Color, message string) {
	color.Sprint(message)
}

/* Table */

// @todo make build table function

func (reporter *Reporter) getTable() table.Writer {
	t := table.NewWriter()
	t.SetAutoIndex(true)
	t.SetStyle(table.StyleLight)
	return t
}

func (reporter *Reporter) prepareReplaceTable() {
	t := reporter.getTable()
	t.SetTitle("Affixed strings")
	t.AppendHeader(table.Row{"Location", "Replaced"})
	reporter.report_table = t
}

func (reporter *Reporter) prepareCheckTable() {
	t := reporter.getTable()
	t.SetTitle("Found strings")
	t.AppendHeader(table.Row{"Location", "Found"})
	reporter.report_table = t
}

func (reporter *Reporter) addRow(location string, count string) {
	reporter.report_table.AppendRow(table.Row{location, count})
}

func (reporter *Reporter) addTotal(total int) {
	reporter.report_table.AppendFooter(table.Row{"", "", "Total", total})
}

func (reporter *Reporter) printTable() {
	fmt.Println(reporter.report_table.Render())
}