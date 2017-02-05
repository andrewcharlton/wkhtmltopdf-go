package wkhtmltopdf

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

// createPDF creates the pdf and writes it to the buffer,
// which can then be written to file or writer.
func (doc *Document) createPDF() (*bytes.Buffer, error) {

	args := append(doc.args(), "-")

	buf := &bytes.Buffer{}
	errbuf := &bytes.Buffer{}

	cmd := exec.Command("wkhtmltopdf", args...)
	cmd.Stdout = buf
	cmd.Stderr = errbuf
	err := cmd.Run()

	if err != nil {
		return nil, fmt.Errorf("Error running wkhtmltopdf: %v", errbuf.String())
	}
	return buf, nil
}

// WriteToFile creates the pdf document and writes it
// to the specified filename.
func (doc *Document) WriteToFile(filename string) error {

	buf, err := doc.createPDF()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0666)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}

	return nil
}

// Write creates the pdf document and writes it
// to the provided reader.
func (doc *Document) Write(w io.Writer) error {

	buf, err := doc.createPDF()
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return fmt.Errorf("Error writing to writer: %v", err)
	}

	return nil
}
