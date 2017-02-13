package wkhtmltopdf

import (
	"bytes"
	"fmt"
	"io"
)

// A Page represents a single html document, which may span multiple pages in
// the finished pdf document. Options can be applied to the page.
type Page struct {
	filename string
	buf      *bytes.Buffer
	reader   bool
	options  []string
}

// NewPage creates a new page from the given filename (which can be a url),
// with the given options.
func NewPage(filename string, opts ...PageOption) *Page {

	pg := &Page{filename: filename, options: []string{}}
	pg.AddOptions(opts...)
	return pg
}

// NewPageReader creates a new page from an io.Reader. The reader will be
// drained on page creation, and stored in a temporary buffer.
func NewPageReader(r io.Reader, opts ...PageOption) (*Page, error) {

	pg := &Page{reader: true, buf: &bytes.Buffer{}, options: []string{}}
	_, err := pg.buf.ReadFrom(r)
	if err != nil {
		return nil, fmt.Errorf("Error reading from reader: %v", err)
	}

	for _, opt := range opts {
		pg.options = append(pg.options, opt.opts()...)
	}

	return pg, nil
}

// AddOptions allows the setting of options after page creation.
func (pg *Page) AddOptions(opts ...PageOption) {

	for _, opt := range opts {
		pg.options = append(pg.options, opt.opts()...)
	}
}
