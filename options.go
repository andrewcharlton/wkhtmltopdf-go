package wkhtmltopdf

import "strconv"

// An Option to be applied to a page or document.
type Option interface {
	opts() []string
}

// A GlobalOption can be applied only to a document.
type GlobalOption struct {
	options []string
}

// opts returns the options, to satisfy the options interface
func (opt GlobalOption) opts() []string { return opt.options }

// A PageOption can be applied only to a page.
type PageOption struct {
	options []string
}

// opts returns the options, to satisfy the options interface
func (opt PageOption) opts() []string { return opt.options }

// Global Options ----------------------------------------------------------

// NoCollate - do not collate when printing multiple copies.
func NoCollate() GlobalOption {
	return GlobalOption{[]string{"--no-collate"}}
}

// CookieJar - read and write cookies from and to the supplied
// cookie jar file.
func CookieJar(path string) GlobalOption {
	return GlobalOption{[]string{"--cookie-jar", path}}
}

// DPI - change the dpi explicitly.
func DPI(dpi int) GlobalOption {
	return GlobalOption{[]string{"--dpi", strconv.Itoa(dpi)}}
}

// Grayscale - PDF will be generated in grayscale.
func Grayscale() GlobalOption {
	return GlobalOption{[]string{"--grayscale"}}
}

// ImageDPI - When embedding images, scale them down to this dpi.
func ImageDPI(dpi int) GlobalOption {
	return GlobalOption{[]string{"--image-dpi", strconv.Itoa(dpi)}}
}

// ImageQuality - When jpeg compressing images, use this quality (default 94).
func ImageQuality(quality int) GlobalOption {
	return GlobalOption{[]string{"--image-quality", strconv.Itoa(quality)}}
}

// LowQuality - Generates lower quality pdf/ps. Useful to shrink the result document space.
func LowQuality() GlobalOption {
	return GlobalOption{[]string{"--low-quality"}}
}

// MarginBottom - Set the page bottom margin.
func MarginBottom(units string) GlobalOption {
	return GlobalOption{[]string{"--margin-bottom", units}}
}

// MarginLeft - Set the page left margin.
func MarginLeft(units string) GlobalOption {
	return GlobalOption{[]string{"--margin-left", units}}
}

// MarginRight - Set the page right margin.
func MarginRight(units string) GlobalOption {
	return GlobalOption{[]string{"--margin-right", units}}
}

// MarginTop - Set the page top margin.
func MarginTop(units string) GlobalOption {
	return GlobalOption{[]string{"--margin-top", units}}
}

// Landscape - Set the page orientation to landscape.
func Landscape() GlobalOption {
	return GlobalOption{[]string{"--orientation", "landscape"}}
}

// PageHeight - Set the page height.
func PageHeight(units string) GlobalOption {
	return GlobalOption{[]string{"--page-height", units}}
}

// PageSize - Set paper size to A4, letter etc.
func PageSize(size string) GlobalOption {
	return GlobalOption{[]string{"--page-size", size}}
}

// PageWidth - Set the page width.
func PageWidth(units string) GlobalOption {
	return GlobalOption{[]string{"--page-width", units}}
}

// NoPDFCompression - Do not use lossless compression on pdf objects.
func NoPDFCompression() GlobalOption {
	return GlobalOption{[]string{"--no-pdf-compression"}}
}

// Quiet - Be less verbose.
func Quiet() GlobalOption {
	return GlobalOption{[]string{"--quiet"}}
}

// Title - the title of the generated pdf file (the title of the first document is used
// if not specified).
func Title(title string) GlobalOption {
	return GlobalOption{[]string{"--title", title}}
}

// Page Options -------------------------------------------------------------------------

// Allow the file or files from the specified folder to be loaded (repeatable)
func Allow(path string) PageOption {
	return PageOption{[]string{"--allow", path}}
}

// Background - print background (default)
func Background() PageOption {
	return PageOption{[]string{"--background"}}
}

// NoBackground - do not print background
func NoBackground() PageOption {
	return PageOption{[]string{"--no-background"}}
}

// BypassProxy - bypass proxy for host (repeatable)
func BypassProxy(host string) PageOption {
	return PageOption{[]string{"--bypass-proxy-for", host}}
}

// CacheDir - web cache directory
func CacheDir(path string) PageOption {
	return PageOption{[]string{"--cache-dir", path}}
}

// CheckboxCheckedSVG - Use this svg file when rendering checked checkboxes
func CheckboxCheckedSVG(path string) PageOption {
	return PageOption{[]string{"--checkbox-checked-svg", path}}
}

// CheckboxSVG - Use this svg file when rendering unchecked checkboxes
func CheckboxSVG(path string) PageOption {
	return PageOption{[]string{"--checkbox-svg", path}}
}

// Cookie - Set an additional cookie (repeatable), value should be url encoded.
func Cookie(name, value string) PageOption {
	return PageOption{[]string{"--cookie", name, value}}
}

// CustomHeader - Set an additional HTTP header (repeatable)
func CustomHeader(name, value string) PageOption {
	return PageOption{[]string{"--custom-header", name, value}}
}

// CustomHeaderPropagation - Add HTTP headers specified by --custom-header for
// each resource request.
func CustomHeaderPropagation() PageOption {
	return PageOption{[]string{"--custom-header-propagation"}}
}

// NoCustomHeaderPropagation - Do not add HTTP headers specified by --custom-header for
// each resource request.
func NoCustomHeaderPropagation() PageOption {
	return PageOption{[]string{"--no-custom-header-propagation"}}
}
