package wkhtmltopdf_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/andrewcharlton/wkhtmltopdf-go"
)

func TestWriteToFile(t *testing.T) {

	testcases := []struct {
		Case     string
		Pages    []string
		Filename string
		Err      string
	}{
		{"Simple", []string{"test_data/simple.html"}, "test_data/simple.pdf", ""},
		{"Missing", []string{"test_data/missing.html"}, "test_data/missing.pdf", "Error running wkhtmltopdf"},
		{"BadFile", []string{"test_data/simple.html"}, "<>/!//bad.pdf", "Error creating file"},
	}

	for _, tc := range testcases {

		doc := wkhtmltopdf.NewDocument()
		for _, pg := range tc.Pages {
			doc.AddPages(wkhtmltopdf.NewPage(pg))
		}
		err := doc.WriteToFile(tc.Filename)
		switch {
		case err == nil && tc.Err != "":
			t.Errorf("%v. Wrong error produced. Expected: %v, Got: %v", tc.Case, tc.Err, err)
		case err == nil:
			continue
		case !strings.HasPrefix(err.Error(), tc.Err):
			t.Errorf("%v. Wrong error produced. Expected: %v, Got: %v", tc.Case, tc.Err, err)
		}
	}
}

type BadWriter struct{}

func (w BadWriter) Write(p []byte) (int, error) {
	return 0, errors.New("Bad writer doesn't write")
}

func TestWriteToReader(t *testing.T) {

	testcases := []struct {
		Case   string
		Pages  []string
		Writer io.Writer
		Err    string
	}{
		{"Simple", []string{"test_data/simple.html"}, &bytes.Buffer{}, ""},
		{"Missing", []string{"test_data/missing.html"}, &bytes.Buffer{}, "Error running wkhtmltopdf"},
		{"Bad Writer", []string{"test_data/simple.html"}, BadWriter{}, "Error writing to writer"},
	}

	for _, tc := range testcases {

		doc := wkhtmltopdf.NewDocument()
		for _, pg := range tc.Pages {
			doc.AddPages(wkhtmltopdf.NewPage(pg))
		}
		err := doc.Write(tc.Writer)
		switch {
		case err == nil && tc.Err != "":
			t.Errorf("%v. Wrong error produced. Expected: %v, Got: %v", tc.Case, tc.Err, err)
		case err == nil:
			continue
		case !strings.HasPrefix(err.Error(), tc.Err):
			t.Errorf("%v. Wrong error produced. Expected: %v, Got: %v", tc.Case, tc.Err, err)
		}
	}
}
