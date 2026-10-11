package feed

import (
	"strings"

	"golang.org/x/net/html"
)

func cleanDescription(description string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(description))
	var text strings.Builder

	for {
		switch tokenizer.Next() {
		case html.ErrorToken:
			return strings.Join(strings.Fields(text.String()), " ")

		case html.TextToken:
			text.Write(tokenizer.Text())

		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken:
			name, _ := tokenizer.TagName()

			switch string(name) {
			case "p", "br", "div", "li":
				text.WriteByte(' ')
			}
		}
	}
}
