package core

import (
	"baliance.com/gooxml/document"
	"fmt"
	"log"
	"strings"
)

// check document title
func checkTitle(paragraphs []document.Paragraph) (nextIndex int) {
	fmt.Println("-------------------- Doc Title Checking --------------------")

	for idx, p := range paragraphs {

		if p.Properties().X().Jc == nil {
			nextIndex = idx
			return
		}

		textAlign := p.Properties().X().Jc.ValAttr.String()

		switch textAlign {
		case "center":
			// confirm 'title'
			for _, r := range p.Runs() {
				// check font type
				fontSize := *r.Properties().X().Sz.ValAttr.ST_UnsignedDecimalNumber

				if fontType := *r.Properties().Fonts().X().EastAsiaAttr; fontType == TitleFontType {

					// check font size
					if fontSize/2 == TitleFontSize {
						// type and size all correct
						fmt.Printf("PASS:\t[%s]\n", r.Text())
					} else {
						fmt.Printf("ERR:\t[%s]Font{expect:%s, actual: %s}\tFontSize{expect:%d, actual:%d}\t\n",
							r.Text(), TitleFontType, fontType, TitleFontSize, fontSize)
					}
				} else {
					fmt.Printf("ERR:\t[%s]Font{expect:%s, actual: %s}\tFontSize{expect:%d, actual:%d}\t\n",
						r.Text(), TitleFontType, fontType, TitleFontSize, fontSize)
				}
			}
		default:
			nextIndex = idx
			return
		}
	}

	return
}

// check document content
func checkContent(paragraphs []document.Paragraph) {
	fmt.Println("-------------------- Doc Content Checking --------------------")

	for _, p := range paragraphs {

		if lineSpacing := *p.Properties().X().Spacing.LineAttr.Int64 / 20; lineSpacing != LineSpacing {
			var tmpTxt string
			for _, r := range p.Runs() {
				tmpTxt += r.Text()
			}
			fmt.Printf("ERR:\t[%s]LineSpacing{expect:%d, actual: %d}\t\n", tmpTxt, LineSpacing, lineSpacing)
		}


		for _, r := range p.Runs() {
			fmt.Printf("[%s]\t%s\n", *r.Properties().Fonts().X().EastAsiaAttr, strings.Trim(r.Text(), " "))

		}
	}
}

// real check function
func Execute() {

	doc, err := document.Open("core/demo.docx")

	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	// get all paragraphs
	paragraphs := doc.Paragraphs()

	// ready to check title
	nextIdx := checkTitle(paragraphs[:4])

	// ready to check content
	checkContent(paragraphs[nextIdx:])

	//p := paragraphs[0].Properties().X().Spacing.LineAttr
	//fmt.Println(p)
}
