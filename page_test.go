package wkhtmltopdf

import (
	"bytes"
	"errors"
	"reflect"
	"testing"
)

func TestNewPage(t *testing.T) {

	pg := NewPage("test.html", Allow("images/"), NoBackground())

	if pg.filename != "test.html" {
		t.Errorf("Wrong filename. Expected: test.html, Got: %v", pg.filename)
	}

	args := []string{"--allow", "images/", "--no-background"}
	if !reflect.DeepEqual(args, pg.options) {
		t.Errorf("Wrong options. Expected: %v, Got: %v", args, pg.options)
	}
}

func TestNewPageReader(t *testing.T) {

	msg := "<html><body><h1>Test Page</h1></body></html>"
	buf := bytes.NewBufferString(msg)
	pg, err := NewPageReader(buf, CacheDir("cache/"))

	if err != nil {
		t.Errorf("Error produced: %v", err)
	}

	args := []string{"--cache-dir", "cache/"}
	if !reflect.DeepEqual(pg.options, args) {
		t.Errorf("Wrong options. Expected: %v, Got: %v", args, pg.options)
	}

	out := pg.buf.String()
	if out != msg {
		t.Errorf("Wrong buffer contents. Expected: %v, Got: %v", msg, out)
	}
}

type brokenReader struct{}

func (r brokenReader) Read(p []byte) (int, error) {
	return 0, errors.New("Broken reader")
}

func TestNewBrokenReader(t *testing.T) {

	_, err := NewPageReader(brokenReader{})
	if err == nil {
		t.Errorf("Error expected, but nil returned.")
	}

	if err.Error() != "Error reading from reader: Broken reader" {
		t.Errorf("Wrong error produced: Got: %v", err)
	}
}
