package wkhtmltopdf_test

import (
	"strings"
	"testing"

	"github.com/andrewcharlton/wkhtmltopdf-go"
)

func TestPDFCreation(t *testing.T) {

	testcases := []struct {
		Case     string
		Pages    []string
		Filename string
		Err      string
	}{
		{"Simple", []string{"test_data/simple.html"}, "test_data/simple.pdf", ""},
		{"Missing", []string{"test_data/missing.html"}, "test_data/missing.pdf", "Error running wkhtmltopdf"},
	}

	for _, tc := range testcases {

		doc := wkhtmltopdf.NewDocument()
		doc.AddPages(tc.Pages...)
		err := doc.WriteToFile(tc.Filename)
		switch {
		case err == nil && tc.Err != "":
			t.Errorf("Wrong error produced. Expected: %v, Got: %v", tc.Err, err)
		case err == nil:
			continue
		case !strings.HasPrefix(err.Error(), tc.Err):
			t.Errorf("Wrong error produced. Expected: %v, Got: %v", tc.Err, err)
		}
	}
}
