package wkhtmltopdf

import (
	"reflect"
	"testing"
)

func TestNewDocument(t *testing.T) {

	doc := NewDocument()
	exp := &Document{pages: []*Page{}, options: []string{}}
	if !reflect.DeepEqual(doc, exp) {
		t.Errorf("NewDocument not produced correctly. Expected: %v, Got: %v", exp, doc)
	}
}

func TestAddPages(t *testing.T) {

	pg1 := NewPage("test1.html")
	pg2 := NewPage("test2.html")
	pg3 := NewPage("test3.html")
	pg4 := NewPage("test4.html")

	testcases := []struct {
		Pages []*Page
		All   []*Page
	}{
		{[]*Page{}, []*Page{}},
		{[]*Page{pg1}, []*Page{pg1}},
		{[]*Page{}, []*Page{pg1}},
		{[]*Page{pg2}, []*Page{pg1, pg2}},
		{[]*Page{pg3, pg4}, []*Page{pg1, pg2, pg3, pg4}},
	}

	doc := NewDocument()
	for _, tc := range testcases {
		for _, pg := range tc.Pages {
			doc.AddPages(pg)
		}
		if !reflect.DeepEqual(doc.pages, tc.All) {
			t.Errorf("Wrong page list. Expected: %v, Got: %v", tc.All, doc.pages)
		}
	}
}

func TestArgs(t *testing.T) {

	doc := NewDocument(Grayscale(), PageSize("A5"))
	pg1 := NewPage("page1.html", Allow("images/"))
	pg2 := NewPage("page2.html", NoBackground())
	doc.AddPages(pg1, pg2)

	args := doc.args()
	exp := []string{"--grayscale", "--page-size", "A5", "page1.html", "--allow",
		"images/", "page2.html", "--no-background"}

	if !reflect.DeepEqual(args, exp) {
		t.Errorf("Wrong error produced. Expected: %v, Got: %v", exp, args)
	}
}
