package wkhtmltopdf

type Page struct {
	filename string

	options []string
}

func NewPage(filename string, opts ...PageOption) *Page {

	pg := &Page{filename: filename, options: []string{}}
	for _, opt := range opts {
		pg.options = append(pg.options, opt.opts()...)
	}
	return pg
}
