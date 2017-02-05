package wkhtmltopdf

import (
	"bytes"
	"fmt"
	"os/exec"
)

// A Document represents a single pdf document.
type Document struct {
	pages []string
}

// NewDocument creates a new document.
func NewDocument() *Document {
	return &Document{pages: []string{}}
}

// AddPages to the document. Pages will be included in
// the final pdf in the order they are added.
func (doc *Document) AddPages(pages ...string) {
	doc.pages = append(doc.pages, pages...)
}

// WriteToFile creates the pdf document and writes it
// to the specified filename.
func (doc *Document) WriteToFile(filename string) error {

	args := append(doc.pages, filename)
	cmd := exec.Command("wkhtmltopdf", args...)
	errbuf := &bytes.Buffer{}
	cmd.Stderr = errbuf
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("Error running wkhtmltopdf: %v", errbuf.String())
	}

	return nil
}
