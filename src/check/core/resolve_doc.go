package core

import (
	"baliance.com/gooxml/document"
	"bytes"
	"fmt"
	"log"
	"strings"
)

// check document title
func checkTitle(paragraphs []document.Paragraph) (nextIndex int, txtBuffer bytes.Buffer) {
	fmt.Fprintln(&txtBuffer, "-------------------- Doc Title Checking --------------------")

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
						fmt.Fprintf(&txtBuffer, "PASS:\t[%s]\n", r.Text())
					} else {
						fmt.Fprintf(&txtBuffer, "ERR:\t[%s]Font{expect:%s, actual: %s}\tFontSize{expect:%d, actual:%d}\t\n",
							r.Text(), TitleFontType, fontType, TitleFontSize, fontSize)
					}
				} else {
					fmt.Fprintf(&txtBuffer, "ERR:\t[%s]Font{expect:%s, actual: %s}\tFontSize{expect:%d, actual:%d}\t\n",
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
func checkContent(paragraphs []document.Paragraph) (txtBuffer bytes.Buffer) {
	fmt.Fprintln(&txtBuffer, "-------------------- Doc Content Checking --------------------")

	for _, p := range paragraphs {

		if lineSpacing := *p.Properties().X().Spacing.LineAttr.Int64 / 20; lineSpacing != LineSpacing {
			var tmpTxt string
			for _, r := range p.Runs() {
				tmpTxt += r.Text()
			}
			fmt.Fprintf(&txtBuffer, "ERR:\t[%s]LineSpacing{expect:%d, actual: %d}\t\n", tmpTxt, LineSpacing, lineSpacing)
		}

		for _, r := range p.Runs() {
			fmt.Fprintf(&txtBuffer, "[%s]\t%s\n", *r.Properties().Fonts().X().EastAsiaAttr, strings.Trim(r.Text(), " "))

		}
	}

	return
}

// real check function
func Execute(filename string) string {

	doc, err := document.Open(filename)

	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	// get all paragraphs
	paragraphs := doc.Paragraphs()


	// ready to check title
	nextIdx, txtBuffer1 := checkTitle(paragraphs[:4])

	// ready to check content
	txtBuffer2 := checkContent(paragraphs[nextIdx:])

	return txtBuffer1.String() + txtBuffer2.String()
}
