package wkhtmltopdf

import (
	"fmt"
	"strconv"
)

// An Option to be applied to a page or document.
type Option interface {
	opts() []string
}

// A GlobalOption can be applied only to a document.
type GlobalOption struct {
	options []string
}

func (opt GlobalOption) opts() []string { return opt.options }

// A PageOption can be applied to pages and/or documents.
type PageOption struct {
	options []string
}

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

// Outline - put an outline into the pdf
func Outline() GlobalOption {
	return GlobalOption{[]string{"--outline"}}
}

// NoOutline - do not put an outline into the pdf
func NoOutline() GlobalOption {
	return GlobalOption{[]string{"--no-outline"}}
}

// OutlineDepth - set the depth of the outline
func OutlineDepth(level int) GlobalOption {
	return GlobalOption{[]string{"--outline-depth", strconv.Itoa(level)}}
}

// DisableDottedLines - do not use dotted lines in the toc
func DisableDottedLines() GlobalOption {
	return GlobalOption{[]string{"--disable-dotted-lines"}}
}

// TocHeaderText - the header text of the toc
func TocHeaderText(text string) GlobalOption {
	return GlobalOption{[]string{"--toc-header-text", text}}
}

// TocLevelIndentation - for each level of headings in the toc indent by this length
func TocLevelIndentation(width string) GlobalOption {
	return GlobalOption{[]string{"--toc-level-indentation", width}}
}

// DisableTocLinks - do not link from toc to sections
func DisableTocLinks() GlobalOption {
	return GlobalOption{[]string{"--disable-toc-links"}}
}

// TocTextSizeShrink - for each level of headings in the toc the font is scaled
// by this factor
func TocTextSizeShrink(factor float64) GlobalOption {
	return GlobalOption{[]string{"--toc-text-size-shrink", fmt.Sprintf("%.3f", factor)}}
}

// XSLStyleSheet - use the supplied xsl style sheet for printing the
// table of content
func XSLStyleSheet(file string) GlobalOption {
	return GlobalOption{[]string{"--xsl-style-sheet", file}}
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

// DefaultHeader - Add a default header, with the name of the page to the left
// and the page numner to the right.
func DefaultHeader() PageOption {
	return PageOption{[]string{"--default-header"}}
}

// Encoding - Set the default text encoding for text input
func Encoding(encoding string) PageOption {
	return PageOption{[]string{"--encoding", encoding}}
}

// DisableExternalLinks - Do not make links to remote web pages
func DisableExternalLinks() PageOption {
	return PageOption{[]string{"--disable-external-links"}}
}

// EnableExternalLinks - Make links to remote web pages
func EnableExternalLinks() PageOption {
	return PageOption{[]string{"--enable-external-links"}}
}

// DisableForms - Do not turn HTML form fields into pdf form fields
func DisableForms() PageOption {
	return PageOption{[]string{"--disable-forms"}}
}

// EnableForms - Turn HTML form fields into pdf form fields
func EnableForms() PageOption {
	return PageOption{[]string{"--enable-forms"}}
}

// Images - do load or print images
func Images() PageOption {
	return PageOption{[]string{"--images"}}
}

// NoImages - do not load or print images
func NoImages() PageOption {
	return PageOption{[]string{"--no-images"}}
}

// DisableInternalLinks - do not make local links
func DisableInternalLinks() PageOption {
	return PageOption{[]string{"--disable-internal-links"}}
}

// EnableInternalLinks - make local links
func EnableInternalLinks() PageOption {
	return PageOption{[]string{"--enable-internal-links"}}
}

// EnableJavascript - do allow web pages to run javascript
func EnableJavascript() PageOption {
	return PageOption{[]string{"--enable-javascript"}}
}

// DisableJavascript - do not allow web pages to run javascript
func DisableJavascript() PageOption {
	return PageOption{[]string{"--disable-javascript"}}
}

// JavascriptDelay - Wait some milliseconds for javascript to finish
func JavascriptDelay(msec int) PageOption {
	return PageOption{[]string{"--javascript-delay", strconv.Itoa(msec)}}
}

// KeepRelativeLinks - keep relative external links as relative external links
func KeepRelativeLinks() PageOption {
	return PageOption{[]string{"--keep-relative-links"}}
}

// LoadErrorHandling - Specify how to handle pages that fail to load: abort, ignore or skip.
func LoadErrorHandling(handler string) PageOption {
	return PageOption{[]string{"--load-error-handling", handler}}
}

// LoadMediaErrorHandling - specify how to handle media pages that fail to load: abort, ignore or skip.
func LoadMediaErrorHandling(handler string) PageOption {
	return PageOption{[]string{"--load-media-error-handling", handler}}
}

// DisableLocalFileAccess - do not allow conversion of a local file to read in other local
// files unless explicitly allowed with Allow()
func DisableLocalFileAccess() PageOption {
	return PageOption{[]string{"--disable-local-file-access"}}
}

// EnableLocalFileAccess - do not allow conversion of a local file to read in other local
// files unless explicitly allowed with Allow()
func EnableLocalFileAccess() PageOption {
	return PageOption{[]string{"--enable-local-file-access"}}
}

// MinFontSize - minimum font size
func MinFontSize(size int) PageOption {
	return PageOption{[]string{"--minimum-font-size", strconv.Itoa(size)}}
}

// ExcludeFromOutline - do not include in the table of contents and outlines
func ExcludeFromOutline() PageOption {
	return PageOption{[]string{"--exclude-from-outline"}}
}

// IncludeInOutline - include in the table of contents and outlines
func IncludeInOutline() PageOption {
	return PageOption{[]string{"--include-in-outline"}}
}

// PageOffset - set the starting page number
func PageOffset(offset int) PageOption {
	return PageOption{[]string{"--page-offset", strconv.Itoa(offset)}}
}

// Password - HTTP authentication password
func Password(password string) PageOption {
	return PageOption{[]string{"--password", password}}
}

// DisablePlugins - disable installed plugins
func DisablePlugins() PageOption {
	return PageOption{[]string{"--disable-plugins"}}
}

// EnablePlugins - enable installed plugins (plugins will likely not work)
func EnablePlugins() PageOption {
	return PageOption{[]string{"--enable-plugins"}}
}

// Post - add an additional post field
func Post(name, value string) PageOption {
	return PageOption{[]string{"--post", name, value}}
}

// PostFile - post an additional file (repeatable)
func PostFile(name, path string) PageOption {
	return PageOption{[]string{"--post-file", name, path}}
}

// PrintMediaType - use print media type instead of screen
func PrintMediaType() PageOption {
	return PageOption{[]string{"--print-media-type"}}
}

// NoPrintMediaType - do not use print media type instead of screen
func NoPrintMediaType() PageOption {
	return PageOption{[]string{"--no-print-media-type"}}
}

// Proxy - use a proxy
func Proxy(proxy string) PageOption {
	return PageOption{[]string{"--proxy", proxy}}
}

// RadioButton - use this svg file when rendering unchecked radio buttons
func RadioButton(path string) PageOption {
	return PageOption{[]string{"--radiobutton-svg", path}}
}

// RadioButtonChecked - use this svg file when rendering checked radio buttons
func RadioButtonChecked(path string) PageOption {
	return PageOption{[]string{"--radiobutton-checked-svg", path}}
}

// ResolveRelativeLinks
func ResolveRelativeLinks() PageOption {
	return PageOption{[]string{"--resolve-relative-links"}}
}

// RunScript
func RunScript(js string) PageOption {
	return PageOption{[]string{"--run-script", js}}
}

// DisableSmartShrinking - disable the intelligent shrinking strategy
// used by webkit that makes the pixel/dpi ratio none constant.
func DisableSmartShrinking() PageOption {
	return PageOption{[]string{"--disable-smart-shrinking"}}
}

// EnableSmartShrinking - enable the intelligent shrinking strategy
// used by webkit that makes the pixel/dpi ratio none constant.
func EnableSmartShrinking() PageOption {
	return PageOption{[]string{"--enable-smart-shrinking"}}
}

// StopSlowScripts - stop slow running javascripts
func StopSlowScripts() PageOption {
	return PageOption{[]string{"--stop-slow-scripts"}}
}

// NoStopSlowScripts
func NoStopSlowScripts() PageOption {
	return PageOption{[]string{"--no-stop-slow-scripts"}}
}

// DisableTocBackLinks - do not link from section header to toc
func DisableTocBackLinks() PageOption {
	return PageOption{[]string{"--disable-toc-back-links"}}
}

// EnableTocBackLinks - link from section header to toc
func EnableTocBackLinks() PageOption {
	return PageOption{[]string{"--enable-toc-back-links"}}
}

// UserStyleSheet - specify a user style sheet, to load with every page
func UserStyleSheet(url string) PageOption {
	return PageOption{[]string{"--user-style-sheet", url}}
}

// Username - HTTP authentication username
func Username(username string) PageOption {
	return PageOption{[]string{"--username", username}}
}

// ViewportSize - set viewport size if you have custom scrollbars or css
// attribute over-flow to emulate window size
func ViewportSize(size string) PageOption {
	return PageOption{[]string{"--viewport-size", size}}
}

// WindowStatus - wait until window.status is equal to this string before
// rendering page
func WindowStatus(status string) PageOption {
	return PageOption{[]string{"--window-status", status}}
}

// Zoom - use this zoom factor
func Zoom(factor float64) PageOption {
	return PageOption{[]string{"--zoom", fmt.Sprintf("%.2f", factor)}}
}

// Footer Options -------------------------------------------------------------------------

// FooterCenter - centered footer text
func FooterCenter(text string) PageOption {
	return PageOption{[]string{"--footer-center", text}}
}

// FooterFontName - set footer font name
func FooterFontName(font string) PageOption {
	return PageOption{[]string{"--footer-font-name", font}}
}

// FooterFontSize - set footer font size
func FooterFontSize(size int) PageOption {
	return PageOption{[]string{"--footer-font-size", strconv.Itoa(size)}}
}

// FooterHTML - Adds an html footer
func FooterHTML(url string) PageOption {
	return PageOption{[]string{"--footer-html", url}}
}

// FooterLeft - left aligned footer text
func FooterLeft(text string) PageOption {
	return PageOption{[]string{"--footer-left", text}}
}

// FooterLine - display line above the footer
func FooterLine() PageOption {
	return PageOption{[]string{"--footer-line"}}
}

// NoFooterLine - do not display line above the footer
func NoFooterLine() PageOption {
	return PageOption{[]string{"--no-footer-line"}}
}

// FooterRight - right aligned footer text
func FooterRight(text string) PageOption {
	return PageOption{[]string{"--footer-right", text}}
}

// FooterSpacing - spacing between the footer and content in mm.
func FooterSpacing(spacing float64) PageOption {
	return PageOption{[]string{"--footer-spacing", fmt.Sprintf("%.2f", spacing)}}
}

// Header Options -------------------------------------------------------------------------

// HeaderCenter - centered header text
func HeaderCenter(text string) PageOption {
	return PageOption{[]string{"--header-center", text}}
}

// HeaderFontName - set header font name
func HeaderFontName(font string) PageOption {
	return PageOption{[]string{"--header-font-name", font}}
}

// HeaderFontSize - set header font size
func HeaderFontSize(size int) PageOption {
	return PageOption{[]string{"--header-font-size", strconv.Itoa(size)}}
}

// HeaderHTML - Adds an html header
func HeaderHTML(url string) PageOption {
	return PageOption{[]string{"--header-html", url}}
}

// HeaderLeft - left aligned header text
func HeaderLeft(text string) PageOption {
	return PageOption{[]string{"--header-left", text}}
}

// HeaderLine - display line above the header
func HeaderLine() PageOption {
	return PageOption{[]string{"--header-line"}}
}

// NoHeaderLine - do not display line above the header
func NoHeaderLine() PageOption {
	return PageOption{[]string{"--no-header-line"}}
}

// HeaderRight - right aligned header text
func HeaderRight(text string) PageOption {
	return PageOption{[]string{"--header-right", text}}
}

// HeaderSpacing - spacing between the header and content in mm.
func HeaderSpacing(spacing float64) PageOption {
	return PageOption{[]string{"--header-spacing", fmt.Sprintf("%.2f", spacing)}}
}

// Replace - replace 'name' with value in header and footer (repeatable).
func Replace(name, value string) PageOption {
	return PageOption{[]string{"--replace", name, value}}
}
