package wkhtmltopdf

// A Page represents a single html document, which may span multiple pages in
// the finished pdf document. Options can be applied to the page.
type Page struct {
	filename string

	options []string
}

// NewPage creates a new page from the given filename (which can be a url),
// with the given options.
func NewPage(filename string, opts ...PageOption) *Page {

	pg := &Page{filename: filename, options: []string{}}
	for _, opt := range opts {
		pg.options = append(pg.options, opt.opts()...)
	}
	return pg
}
