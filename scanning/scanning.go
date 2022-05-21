package scanning

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Scan struct {
	folder                   string
	allowed_files            []string
	denied_files_and_folders []string
	found                    []string
}

/*
	@todo make following function as methods of Scan
	@todo give Scan's properties from main
	@todo make tests for Scan methods
*/

func Run() {
	folders, files := scan_recursive("testdata", []string{"index.blade.php", "denyThisFolder"}, []string{".blade.php", ".jsx"}, true)

	// Files
	for _, i := range files {
		fmt.Println(i)
	}

	// Folders
	for _, i := range folders {
		fmt.Println(i)
	}
}

func scan_recursive(dir_path string, ignore []string, allowed_files []string, with_dirs bool) ([]string, []string) {

	folders := []string{}
	files := []string{}

	// Scan
	filepath.Walk(dir_path, func(path string, f os.FileInfo, err error) error {

		_continue := false

		// ignore unwanted folders and files
		deny_files_and_folders(ignore, path, &_continue)

		if _continue == false {

			f, err = os.Stat(path)

			// If no error
			if err != nil {
				log.Fatal(err)
			}

			needed := false

			// check if name contains at least one from allowed files
			check_allowed_types(allowed_files, f.Name(), &needed)

			if needed {
				// File & Folder Mode
				f_mode := f.Mode()

				// Is folder
				if f_mode.IsDir() {

					if with_dirs {
						// Append to Folders Array
						folders = append(folders, path)
					}

					// Is file
				} else if f_mode.IsRegular() {

					// Append to Files Array
					files = append(files, path)
				}
			}

		}

		return nil
	})

	return folders, files
}

func deny_files_and_folders(ignore []string, path string, _continue *bool) {
	// Loop : Ignore Files & Folders
	for _, i := range ignore {

		// If ignored path
		if strings.Index(path, i) != -1 {
			// Continue
			*_continue = true
		}
	}
}

func check_allowed_types(allowed_files []string, path string, needed *bool) {
	for _, i := range allowed_files {
		// If file allowed
		if strings.Contains(path, i) {
			// Continue
			*needed = true
		}
	}
}
