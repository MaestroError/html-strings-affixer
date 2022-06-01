package scanning

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Scanning struct {
	folder                   string
	allowed_file_types       []string
	denied_files_and_folders []string
	found_files              []string
	found_folders            []string
}

func (scan *Scanning) Run() []string {
	scan.scan_recursive(false)
	return scan.found_files
}

func (scan *Scanning) RunWithFolders() ([]string, []string) {
	scan.scan_recursive(true)
	return scan.found_files, scan.found_folders
}

func (scan *Scanning) Get(folder string, allowed_file_types []string, deny_files_and_folders []string) []string {
	scan.SetFolder(folder)
	scan.SetAllowedFiles(allowed_file_types)
	scan.SetDeniedFilesAndFolders(deny_files_and_folders)
	return scan.Run()
}

/* Setters */
// publics
func (scan *Scanning) SetDeniedFilesAndFolders(files []string) {
	scan.denied_files_and_folders = files
}

func (scan *Scanning) SetAllowedFiles(files []string) {
	scan.allowed_file_types = files
}

func (scan *Scanning) SetFolder(folder string) {
	scan.folder = folder
}

// privates
func (scan *Scanning) setFoundFiles(files []string) {
	scan.found_files = files
}

func (scan *Scanning) setFoundFolders(folders []string) {
	scan.found_folders = folders
}

/* Setters End */

func (scan *Scanning) scan_recursive(with_dirs bool) ([]string, []string) {

	folders := []string{}
	files := []string{}

	// Scan
	filepath.Walk(scan.folder, func(path string, f os.FileInfo, err error) error {

		_continue := false

		// ignore unwanted folders and files
		scan.deny_files_and_folders(path, &_continue)

		if _continue == false {

			f, err = os.Stat(path)

			// If no error
			if err != nil {
				log.Fatal(err)
			}

			needed := false

			// check if name contains at least one from allowed files
			scan.check_allowed_types(f.Name(), &needed)

			if needed {
				// File & Folder Mode
				f_mode := f.Mode()

				// Is folder
				if f_mode.IsDir() {

					if with_dirs {
						// Append to Folders Array
						scan.found_folders = append(scan.found_folders, path)
					}

					// Is file
				} else if f_mode.IsRegular() {

					// Append to Files Array
					scan.found_files = append(scan.found_files, path)
				}
			}

		}

		return nil
	})

	return folders, files
}

func (scan *Scanning) deny_files_and_folders(path string, _continue *bool) {
	if len(scan.denied_files_and_folders) > 0 {
		// Loop : Ignore Files & Folders
		for _, i := range scan.denied_files_and_folders {

			// If ignored path
			if strings.Index(path, i) != -1 {
				// Continue
				*_continue = true
			}
		}
	}
}

func (scan *Scanning) check_allowed_types(path string, needed *bool) {
	if len(scan.allowed_file_types) > 0 {
		for _, i := range scan.allowed_file_types {
			// If file allowed
			if strings.Contains(path, i) {
				// Continue
				*needed = true
			}
		}
	}
}
