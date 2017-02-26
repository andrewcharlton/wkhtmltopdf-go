package wkhtmltopdf

import (
	"reflect"
	"testing"
)

func TestGlobalOptions(t *testing.T) {

	testcases := []struct {
		Options []Option
		Args    []string
	}{
		{[]Option{}, []string{}},
		{[]Option{NoCollate()}, []string{"--no-collate"}},
		{[]Option{DPI(300), Grayscale()}, []string{"--dpi", "300", "--grayscale"}},
		{[]Option{ImageDPI(500), ImageQuality(60)}, []string{"--image-dpi", "500", "--image-quality", "60"}},
		{[]Option{LowQuality(), CookieJar("testpath")}, []string{"--low-quality", "--cookie-jar", "testpath"}},
		{[]Option{MarginBottom("1cm"), MarginTop("2cm")}, []string{"--margin-bottom", "1cm", "--margin-top", "2cm"}},
		{[]Option{MarginLeft("15mm"), MarginRight("1.5cm")}, []string{"--margin-left", "15mm", "--margin-right", "1.5cm"}},
		{[]Option{Landscape(), PageHeight("20cm")}, []string{"--orientation", "landscape", "--page-height", "20cm"}},
		{[]Option{PageWidth("35cm")}, []string{"--page-width", "35cm"}},
		{[]Option{NoPDFCompression(), Quiet()}, []string{"--no-pdf-compression", "--quiet"}},
		{[]Option{Title("An excellent pdf")}, []string{"--title", "An excellent pdf"}},
		{[]Option{PageSize("A4")}, []string{"--page-size", "A4"}},
		{[]Option{Outline(), OutlineDepth(2)}, []string{"--outline", "--outline-depth", "2"}},
		{[]Option{NoOutline()}, []string{"--no-outline"}},
		{[]Option{DisableDottedLines(), TocHeaderText("A lovely header")}, []string{"--disable-dotted-lines", "--toc-header-text", "A lovely header"}},
		{[]Option{TocLevelIndentation("2mm")}, []string{"--toc-level-indentation", "2mm"}},
		{[]Option{DisableTocLinks()}, []string{"--disable-toc-links"}},
		{[]Option{TocTextSizeShrink(0.8)}, []string{"--toc-text-size-shrink", "0.800"}},
		{[]Option{XSLStyleSheet("./style.css")}, []string{"--xsl-style-sheet", "./style.css"}},
	}

	for _, tc := range testcases {
		doc := NewDocument(tc.Options...)
		if !reflect.DeepEqual(doc.options, tc.Args) {
			t.Errorf("Wrong arguments created. Expected: %v, Got: %v", tc.Args, doc.options)
		}
	}
}

func TestPageOptions(t *testing.T) {

	testcases := []struct {
		Options []PageOption
		Args    []string
	}{
		{[]PageOption{Allow("test/")}, []string{"--allow", "test/"}},
		{[]PageOption{Background()}, []string{"--background"}},
		{[]PageOption{NoBackground()}, []string{"--no-background"}},
		{[]PageOption{BypassProxy("test.com")}, []string{"--bypass-proxy-for", "test.com"}},
		{[]PageOption{CacheDir("cache/")}, []string{"--cache-dir", "cache/"}},
		{[]PageOption{CheckboxCheckedSVG("path")}, []string{"--checkbox-checked-svg", "path"}},
		{[]PageOption{CheckboxSVG("path")}, []string{"--checkbox-svg", "path"}},
		{[]PageOption{Cookie("name", "value")}, []string{"--cookie", "name", "value"}},
		{[]PageOption{CustomHeader("name", "value")}, []string{"--custom-header", "name", "value"}},
		{[]PageOption{CustomHeaderPropagation()}, []string{"--custom-header-propagation"}},
		{[]PageOption{NoCustomHeaderPropagation()}, []string{"--no-custom-header-propagation"}},
		{[]PageOption{DefaultHeader()}, []string{"--default-header"}},
		{[]PageOption{Encoding("UTF8")}, []string{"--encoding", "UTF8"}},
		{[]PageOption{DisableExternalLinks()}, []string{"--disable-external-links"}},
		{[]PageOption{EnableExternalLinks()}, []string{"--enable-external-links"}},
		{[]PageOption{DisableForms()}, []string{"--disable-forms"}},
		{[]PageOption{EnableForms()}, []string{"--enable-forms"}},
		{[]PageOption{Images()}, []string{"--images"}},
		{[]PageOption{NoImages()}, []string{"--no-images"}},
		{[]PageOption{DisableInternalLinks()}, []string{"--disable-internal-links"}},
		{[]PageOption{EnableInternalLinks()}, []string{"--enable-internal-links"}},
		{[]PageOption{EnableJavascript()}, []string{"--enable-javascript"}},
		{[]PageOption{DisableJavascript()}, []string{"--disable-javascript"}},
		{[]PageOption{JavascriptDelay(100)}, []string{"--javascript-delay", "100"}},
		{[]PageOption{KeepRelativeLinks()}, []string{"--keep-relative-links"}},
		{[]PageOption{LoadErrorHandling("abort")}, []string{"--load-error-handling", "abort"}},
		{[]PageOption{LoadMediaErrorHandling("skip")}, []string{"--load-media-error-handling", "skip"}},
		{[]PageOption{DisableLocalFileAccess()}, []string{"--disable-local-file-access"}},
		{[]PageOption{EnableLocalFileAccess()}, []string{"--enable-local-file-access"}},
		{[]PageOption{MinFontSize(12)}, []string{"--minimum-font-size", "12"}},
		{[]PageOption{ExcludeFromOutline()}, []string{"--exclude-from-outline"}},
		{[]PageOption{IncludeInOutline()}, []string{"--include-in-outline"}},
		{[]PageOption{PageOffset(5)}, []string{"--page-offset", "5"}},
		{[]PageOption{Password("p4ssw0rd")}, []string{"--password", "p4ssw0rd"}},
		{[]PageOption{DisablePlugins()}, []string{"--disable-plugins"}},
		{[]PageOption{EnablePlugins()}, []string{"--enable-plugins"}},
		{[]PageOption{Post("name", "value")}, []string{"--post", "name", "value"}},
		{[]PageOption{PostFile("file", "postfile.txt")}, []string{"--post-file", "file", "postfile.txt"}},
		{[]PageOption{PrintMediaType()}, []string{"--print-media-type"}},
		{[]PageOption{NoPrintMediaType()}, []string{"--no-print-media-type"}},
		{[]PageOption{Proxy("proxy name")}, []string{"--proxy", "proxy name"}},
		{[]PageOption{RadioButton("unchecked.svg")}, []string{"--radiobutton-svg", "unchecked.svg"}},
		{[]PageOption{RadioButtonChecked("checked.svg")}, []string{"--radiobutton-checked-svg", "checked.svg"}},
		{[]PageOption{ResolveRelativeLinks()}, []string{"--resolve-relative-links"}},
		{[]PageOption{RunScript("myscript.js")}, []string{"--run-script", "myscript.js"}},
		{[]PageOption{DisableSmartShrinking()}, []string{"--disable-smart-shrinking"}},
		{[]PageOption{EnableSmartShrinking()}, []string{"--enable-smart-shrinking"}},
		{[]PageOption{StopSlowScripts()}, []string{"--stop-slow-scripts"}},
		{[]PageOption{NoStopSlowScripts()}, []string{"--no-stop-slow-scripts"}},
		{[]PageOption{DisableTocBackLinks()}, []string{"--disable-toc-back-links"}},
		{[]PageOption{EnableTocBackLinks()}, []string{"--enable-toc-back-links"}},
		{[]PageOption{UserStyleSheet("style.css")}, []string{"--user-style-sheet", "style.css"}},
		{[]PageOption{Username("user")}, []string{"--username", "user"}},
		{[]PageOption{ViewportSize("size")}, []string{"--viewport-size", "size"}},
		{[]PageOption{WindowStatus("done")}, []string{"--window-status", "done"}},
		{[]PageOption{Zoom(0.9)}, []string{"--zoom", "0.90"}},
		{[]PageOption{FooterCenter("Footer Text")}, []string{"--footer-center", "Footer Text"}},
		{[]PageOption{FooterFontName("DejaVu")}, []string{"--footer-font-name", "DejaVu"}},
		{[]PageOption{FooterFontSize(12)}, []string{"--footer-font-size", "12"}},
		{[]PageOption{FooterHTML("footer.html")}, []string{"--footer-html", "footer.html"}},
		{[]PageOption{FooterLeft("Footer Text")}, []string{"--footer-left", "Footer Text"}},
		{[]PageOption{FooterLine()}, []string{"--footer-line"}},
		{[]PageOption{NoFooterLine()}, []string{"--no-footer-line"}},
		{[]PageOption{FooterRight("Footer Text")}, []string{"--footer-right", "Footer Text"}},
		{[]PageOption{FooterSpacing(2.5)}, []string{"--footer-spacing", "2.50"}},
		{[]PageOption{HeaderCenter("Header Text")}, []string{"--header-center", "Header Text"}},
		{[]PageOption{HeaderFontName("DejaVu")}, []string{"--header-font-name", "DejaVu"}},
		{[]PageOption{HeaderFontSize(12)}, []string{"--header-font-size", "12"}},
		{[]PageOption{HeaderHTML("header.html")}, []string{"--header-html", "header.html"}},
		{[]PageOption{HeaderLeft("Header Text")}, []string{"--header-left", "Header Text"}},
		{[]PageOption{HeaderLine()}, []string{"--header-line"}},
		{[]PageOption{NoHeaderLine()}, []string{"--no-header-line"}},
		{[]PageOption{HeaderRight("Header Text")}, []string{"--header-right", "Header Text"}},
		{[]PageOption{HeaderSpacing(2.5)}, []string{"--header-spacing", "2.50"}},
		{[]PageOption{Replace("name", "Donald")}, []string{"--replace", "name", "Donald"}},
	}

	for _, tc := range testcases {
		pg := NewPage("", tc.Options...)
		if !reflect.DeepEqual(pg.options, tc.Args) {
			t.Errorf("Wrong page arguments created. Expected: %v, Got: %v", tc.Args, pg.options)
		}
	}

}
