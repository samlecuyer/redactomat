package redactomat

import (
	"testing"
	"bytes"
	"bufio"
)

type decodeTest struct {
	input    string
	expected string
}

func (test decodeTest) Run(t *testing.T) {
	r := bufio.NewReader(bytes.NewBufferString(test.input))
	ps, err := Redact(r)
	if err != nil {
		t.Fatalf("parsing %s: %s", test.input, err)
	}

	if ps != test.expected {
		t.Fatalf("Expected, Actual(%#v, %#v)", test.expected, ps)
	}
}


func TestDecode_HappyPath(t *testing.T) {
	decodeTest {
		"<html><head></head><body>abc</body></html>",
		"<html><body>abc</body></html>",
	}.Run(t)
}

func TestDecode_Scripts(t *testing.T) {
	decodeTest {
		"<html><head><script></script></head><body>abc</body></html>",
		"<html><body>abc</body></html>",
	}.Run(t)
}

func TestDecode_Images(t *testing.T) {
	decodeTest {
		`<html><head></head><body><img src="blach"/></body></html>`,
		`<html><body><img data-src="blach"/></body></html>`,
	}.Run(t)
}
