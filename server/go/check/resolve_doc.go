package check

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
)

// open doc
func getParagraphs(filename string) (paragraphs []document.Paragraph, err error) {
	doc, error := document.Open(filename)

	if error != nil {
		err = error
		paragraphs = []document.Paragraph{}
		return
	}

	// get all paragraphs from document
	paragraphs = doc.Paragraphs()
	error = nil
	return
}

// check document title
func checkTitle(paragraphs []document.Paragraph) {

	for _, p := range paragraphs {
		paragraphs = append(paragraphs, p)
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {

			// text := r.Text()

			// check font type
			if fontType := *r.Properties().Fonts().X().AsciiAttr; fontType == TitleFontType {
				// correct
			} else {
				// incorrect
				fmt.Println("Document Title Error")
			}
		}
	}

}

// real check function
func Execute() {
	paragraphs, err := getParagraphs("check/demo.docx")

	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	checkTitle(paragraphs)
}
