package config

import (
	"testing"
)

func TestConfigSetsAllNeededDefaults(t *testing.T) {
	conf := Config{}
	confEmpty := Config{}
	conf.Init()

	// strings
	if conf.folder_to_scan == confEmpty.folder_to_scan {
		t.Errorf("Config's field %q is empty", "folder_to_scan")
	}

	if conf.prefix_to_set == confEmpty.prefix_to_set {
		t.Errorf("Config's field %q is empty", "prefix_to_set")
	}

	if conf.suffix_to_set == confEmpty.suffix_to_set {
		t.Errorf("Config's field %q is empty", "suffix_to_set")
	}

	// slices
	if conf.allowed_file_types == nil {
		t.Errorf("Config's field %q is empty", "allowed_file_types")
	}

	if conf.ignore_files_and_folders == nil {
		t.Errorf("Config's field %q shouldn't be empty", "ignore_files_and_folders")
	}

	if conf.ignore_characters == nil {
		t.Errorf("Config's field %q shouldn't be empty", "ignore_characters")
	}

	if conf.allowed_parse_methods == nil {
		t.Errorf("Config's field %q shouldn't be empty", "allowed_parse_methods")
	}
}
