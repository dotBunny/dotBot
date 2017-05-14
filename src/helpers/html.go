package helpers

import (
	"strings"

	"golang.org/x/net/html"
)

// GetAttribute returns the attribute value in a token
func GetAttribute(attr string, token html.Token) (ok bool, value string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range token.Attr {
		if a.Key == attr {
			return true, a.Val
		}
	}
	return false, ""
}

// GetElementContent will look for the following text nodes content
func GetElementContent(tag string, content string, offset int) (found bool, result string) {

	foundTag := false
	z := html.NewTokenizer(strings.NewReader(content))

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == tag {
				if offset == 0 {
					foundTag = true
				}
				offset--
			}
		case tt == html.TextToken:
			t := z.Token()

			if t.Data == tag {
				if offset == 0 {
					foundTag = true
				}
				offset--
			}
			if foundTag {
				return true, t.Data
			}
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == tag {
				if offset == 0 {
					foundTag = true
				}
				offset--
			} else {
				if foundTag {
					foundTag = false
				}
			}
		}
	}

}

// GetElement returns the first tag found with the selector
func GetElement(tag string, content string, offset int) (found bool, element html.Token) {

	z := html.NewTokenizer(strings.NewReader(content))

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == tag {
				if offset == 0 {
					return true, t
				}
				offset--
			}
		case tt == html.TextToken:
			t := z.Token()
			if t.Data == tag {
				if offset == 0 {
					return true, t
				}
				offset--
			}
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == tag {
				if offset == 0 {
					return true, t
				}
				offset--
			}
		}
	}
}

// Clean content, removing special codes
func Clean(content string) (cleaned string) {
	cleaned = content
	cleaned = strings.Replace(content, "&amp;", "&", -1)
	return cleaned
}
