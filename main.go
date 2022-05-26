package main

import (
	"github.com/MaestroError/html-strings-affixer/scanning"
)

// "testdata", []string{"index.blade.php", "denyThisFolder"}, []string{".blade.php", ".jsx"}
func main() {
	Scan := scanning.Scanning{}
	Scan.SetFolder("testdata")
}
