package wkhtmltopdf_test

import (
	"log"

	"github.com/andrewcharlton/wkhtmltopdf-go"
)

func main() {

	doc := wkhtmltopdf.NewDocument()
	pg := wkhtmltopdf.NewPage("www.google.com")
	doc.AddPages(pg)

	err := doc.WriteToFile("google.pdf")
	if err != nil {
		log.Fatal("Error writing to google.pdf")
	}

}
