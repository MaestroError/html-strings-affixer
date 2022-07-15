package affixer

type affixer struct {
	path       string
	content    string
	prefix     string
	suffix     string
	use_logger bool
	use_backup bool
}
