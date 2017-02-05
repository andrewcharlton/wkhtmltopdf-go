package wkhtmltopdf

import (
	"bytes"
	"fmt"
	"os/exec"
)

// A Document represents a single pdf document.
type Document struct {
	pages   []*Page
	options []string
}

// NewDocument creates a new document.
func NewDocument(opts ...Option) *Document {

	doc := &Document{pages: []*Page{}, options: []string{}}
	for _, opt := range opts {
		doc.options = append(doc.options, opt.opts()...)
	}
	return doc
}

// AddPages to the document. Pages will be included in
// the final pdf in the order they are added.
func (doc *Document) AddPages(pages ...*Page) {
	doc.pages = append(doc.pages, pages...)
}

// args calculates the args needed to run wkhtmltopdf
func (doc *Document) args() []string {

	args := []string{}
	args = append(args, doc.options...)
	for _, pg := range doc.pages {
		args = append(args, pg.filename)
		args = append(args, pg.options...)
	}
	return args
}

// WriteToFile creates the pdf document and writes it
// to the specified filename.
func (doc *Document) WriteToFile(filename string) error {

	args := append(doc.args(), filename)

	cmd := exec.Command("wkhtmltopdf", args...)
	errbuf := &bytes.Buffer{}
	cmd.Stderr = errbuf
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("Error running wkhtmltopdf: %v", errbuf.String())
	}

	return nil
}
