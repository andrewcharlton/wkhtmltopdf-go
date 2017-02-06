// Package wkhtmltopdf implements a wrapper for the html to pdf
// converter wkhtmltopdf.
package wkhtmltopdf

var (

	// Executable is the command to run wkhtmltopdf. If wkhtmltopdf
	// cannot be found on your path, amend this to its location.
	Executable = "wkhtmltopdf"

	// TempDir is where the directories for creating temporary
	// files are created.
	TempDir = "."
)
