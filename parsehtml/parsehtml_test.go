package parsehtml

import (
	"testing"

	"github.com/MaestroError/html-strings-affixer/config"
)

func TestStringsExtractionWorks(t *testing.T) {
	parse := Parsehtml{}
	file := "..\\testdata\\parsehtml\\test_string.blade.php"
	config := config.Config{}
	config.Init()
	parse.Init(file, config)
	parse.ExtractText()
	found := parse.GetFoundStrings()

	needToFind := "Testing string"

	if found["data"][0]["found"] != needToFind {
		t.Errorf("got %q, wanted %q", found[file][0]["found"], needToFind)
	}
}

func TestHashtagExtractionWorks(t *testing.T) {
	parse := Parsehtml{}
	file := "..\\testdata\\parsehtml\\test_hashtag.blade.php"
	config := config.Config{}
	config.Init()
	parse.Init(file, config)
	parse.ExtractHashtag()
	found := parse.GetFoundStrings()

	needToFind := "Tagged string"

	if found["data"][0]["found"] != needToFind {
		t.Errorf("got %q, wanted %q", found[file][0]["found"], needToFind)
	}
}
