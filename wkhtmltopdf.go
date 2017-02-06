/*
Package wkhtmltopdf implements a wrapper for the html to pdf converter wkhtmltopdf.

For more information on wkhtmltopdf see: http://wkhtmltopdf.org/

Creating Documents

Creating a pdf document is simple:

	doc := wkhtmltopdf.NewDocument()
	pg := wkhtmltopdf.NewPage("www.google.com")
	doc.AddPages(pg)
	doc.WriteToFile("google.pdf")

Pages can be sourced from both URLs and local filenames.

Applying Options

You can apply options to both the document and individual pages when creating them.

	doc := wkhtmltopdf.NewDocument(wkhtmltopdf.Grayscale(), wkhtmltopdf.PageSize("A4"))
	pg := wkhtmltopdf.NewPage("www.google.com", wkhtmltopdf.DefaultHeader())
	doc.AddPages(pg)
	doc.WriteToFile("google.pdf")



Using Readers and Writers

As well as URLs/filenames, you can source pages from an io.Reader, and write them to an
io.Writer.

If a single reader is provided, this is piped to wkhtmltopdf through stdin. If multiple
readers are provided, the contents are written to a temporary directory and used from there.
The location of the temporary directories can be changed by setting the TempDir variable.

	doc := wkhtmltopdf.NewDocument()

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






*/
package wkhtmltopdf

var (

	// TempDir is where the directories for creating temporary
	// files are created.
	TempDir = "."
)
