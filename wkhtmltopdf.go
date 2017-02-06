/*
Package wkhtmltopdf implements a wrapper for the html to pdf converter wkhtmltopdf.

For more information on wkhtmltopdf see: http://wkhtmltopdf.org/

Creating Documents

Creating a pdf document is simple:

	import "github.com/andrewcharlton/wkhtmltopdf-go"

	func main() {

		doc := wkhtmltopdf.NewDocument()
		pg := wkhtmltopdf.NewPage("www.google.com")
		doc.AddPages(pg)
		doc.WriteToFile("google.pdf")

	}

Pages can be sourced from both URLs and local filenames.

Applying Options

You can apply options to both the document and individual pages when creating them.
Both GlobalOptions and PageOptions can be applied to a document, but only PageOptions
can be applied to pages. For example:

	import "github.com/andrewcharlton/wkhtmltopdf-go"

	func main() {

		doc := wkhtmltopdf.NewDocument(wkhtmltopdf.Grayscale())
		pg := wkhtmltopdf.NewPage("www.google.com", wkhtmltopdf.DefaultHeader())
		doc.AddPages(pg)
		doc.WriteToFile("google.pdf")

	}


Will convert the google homepage to a grayscale pdf, with wkhtmltopdf's default heading.

Using Readers and Writers

As well as URLs/filenames, you can source pages from an io.Reader, and write them to an
io.Writer. If using a reader as source, it will be drained on page creation, and used to
create a temporary file. Modify the TempDir variable to specify where you would like these
files to be created. Each document creates its temp files in its own temporary directory so
*should* be safe for concurrent use.

	import (
		"bytes"

		"github.com/andrewcharlton/wkhtmltopdf-go"
	)

	func main() {

		doc := wkhtmltopdf.NewDocument(wkhtmltopdf.Grayscale())

		buf := bytes.NewBufferString("<html><body><h1>Test Page</h1></body></html>")
		pg, err := wkhtmltopdf.NewPageReader(buf)
		if err != nil {
			log.Fatal("Error reading from reader.")
		}
		doc.AddPages(pg)

		output := &bytes.Buffer{}
		err := doc.Write(output)
		if err != nil {
			log.Fatal("Error writing to writer.")
		}

	}






*/
package wkhtmltopdf

var (

	// TempDir is where the directories for creating temporary
	// files are created.
	TempDir = "."
)
