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
	}

	for _, tc := range testcases {
		pg := NewPage("", tc.Options...)
		if !reflect.DeepEqual(pg.options, tc.Args) {
			t.Errorf("Wrong arguments created. Expected: %v, Got: %v", tc.Args, pg.options)
		}
	}

}
