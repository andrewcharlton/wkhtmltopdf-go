# wkhtmltopdf-go

[![GoDoc](https://godoc.org/github.com/andrewcharlton/wkhtmltopdf-go?status.svg)](https://godoc.org/github.com/andrewcharlton/wkhtmltopdf-go)
[![Build Status](https://travis-ci.org/andrewcharlton/wkhtmltopdf-go.svg?branch=master)](https://travis-ci.org/andrewcharlton/wkhtmltopdf-go)
[![Coverage Status](https://coveralls.io/repos/github/andrewcharlton/wkhtmltopdf-go/badge.svg?branch=master)](https://coveralls.io/github/andrewcharlton/wkhtmltopdf-go?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/andrewcharlton/wkhtmltopdf-go)](https://goreportcard.com/report/github.com/andrewcharlton/wkhtmltopdf-go)


wkhtmltopdf-go is a go wrapper for [wkhtmltopdf](www.wkhtmltopdf.org).

It is intended to simplify the production of pdf documents in go, using html/css templates as a source. Multiple input
sources, specified by filename, url or an io.Reader, can be combined to form a single pdf and written to either a file
or an io.Writer.


## Installation

To get the go library:
```
go get -u github.com/andrewcharlton/wkhtmltopdf-go
```

[wkhtmltopdf](www.wkhtmltopdf.org) also needs to be installed. It is assumed that wkhtmltopdf can be found on your PATH.
If this is not the case, you can set the Executable variable to wkhtmltopdf's location.


## Example Usage

To create a simple pdf from a url:

``` go
import "github.com/andrewcharlton/wkhtmltopdf-go"

func GooglePDF() {

	doc := wkhtmltopdf.NewDocument()
	pg := wkhtmltopdf.NewPage("www.google.com")
	doc.AddPages(pg)

	doc.WriteToFile("google.pdf")
}
```

Options can be set when creating both documents and pages. To create a landscape document without images:

``` go
import "github.com/andrewcharlton/wkhtmltopdf-go"

func GooglePDF() {

	doc := wkhtmltopdf.NewDocument(wkhtmltopdf.Landscape())
	pg := wkhtmltopdf.NewPage("www.google.com", wkhtmltopdf.NoImages())
	doc.AddPages(pg)

	doc.WriteToFile("google.pdf")
}
```

See [GoDoc](https://godoc.org/github.com/andrewcharlton/wkhtmltopdf-go) for a full list of options.

Documents can also be created from io.Readers, and written to io.Writers to facilitate use in a server environment.
If multiple readers are used, these will be written to temporary files. Each document creates its own temporary
directory, so *should* be safe for concurrent use. If only a single reader is used, this will be piped to wkhtmltopdf.

``` go
package main

import (
    "bytes"
    "html/template"
    "log"
    "net/http"

    "github.com/andrewcharlton/wkhtmltopdf-go"
)

const page = `
<html>
  <body>
    <h1>Test Page</h1>

	<p>Path: {{.}}</p>
  </body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {

    tmpl := template.Must(template.New("page").Parse(page))
    buf := &bytes.Buffer{}
    tmpl.Execute(buf, r.URL.String())

    doc := wkhtmltopdf.NewDocument()
    pg, err := wkhtmltopdf.NewPageReader(buf)
    if err != nil {
        log.Fatal("Error reading page buffer")
    }
    doc.AddPages(pg)

    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", `attachment; filename="test.pdf"`)
    err = doc.Write(w)
    if err != nil {
        log.Fatal("Error serving pdf")
    }
}

func main() {

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)

}
```


