package config

import (
	"testing"
)

func TestConfigSetsAllNeededDefaults(t *testing.T) {
	conf := Config{}
	confEmpty := Config{}
	conf.Init()

	// strings
	if conf.Folder_to_scan == confEmpty.Folder_to_scan {
		t.Errorf("Config's field %q is empty", "Folder_to_scan")
	}

	if conf.Prefix_to_set == confEmpty.Prefix_to_set {
		t.Errorf("Config's field %q is empty", "Prefix_to_set")
	}

	if conf.Suffix_to_set == confEmpty.Suffix_to_set {
		t.Errorf("Config's field %q is empty", "Suffix_to_set")
	}

	// slices
	if conf.Allowed_file_types == nil {
		t.Errorf("Config's field %q is empty", "Allowed_file_types")
	}

	if conf.Ignore_files_and_folders == nil {
		t.Errorf("Config's field %q shouldn't be empty", "Ignore_files_and_folders")
	}

	if conf.Ignore_characters == nil {
		t.Errorf("Config's field %q shouldn't be empty", "Ignore_characters")
	}

	if conf.Allowed_parse_methods == nil {
		t.Errorf("Config's field %q shouldn't be empty", "Allowed_parse_methods")
	}
}
