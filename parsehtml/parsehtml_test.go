package parsehtml

import (
	"testing"
)

func TestStringsExtractionWorks(t *testing.T) {
	parse := Parsehtml{}
	file := "..\\testdata\\parsehtml\\test_string.blade.php"
	parse.Init(file)
	parse.ExtractText()
	found := parse.GetFoundStrings()

	needToFind := "Testing string"

	if found[file][0]["found"] != needToFind {
		t.Errorf("got %q, wanted %q", found[file][0]["found"], needToFind)
	}
}

func TestHashtagExtractionWorks(t *testing.T) {
	parse := Parsehtml{}
	file := "..\\testdata\\parsehtml\\test_hashtag.blade.php"
	parse.Init(file)
	parse.ExtractHashtag()
	found := parse.GetFoundStrings()

	needToFind := "Tagged string"

	if found[file][0]["found"] != needToFind {
		t.Errorf("got %q, wanted %q", found[file][0]["found"], needToFind)
	}
}
