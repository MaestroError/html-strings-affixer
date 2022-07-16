package reporter

type Reporter struct {
	data     map[string]string
	warnings []string
	detailed bool // if true, printing detailed report (with strings and line)
}

/*
	@todo make data report (table) and warning report methods
*/
