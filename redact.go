package redactomat

import (
	"bytes"
	"bufio"
	"io"
	"code.google.com/p/go.net/html"
)

func Redact(r io.Reader) (string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				switch c.Data {
				case "style", "script", "head", "meta":
					n.RemoveChild(c)
					return
				case "img":
					for i, attr := range c.Attr {
						if attr.Key == "src" {
							c.Attr[i].Key = "data-redacted-src"
						}	
					}
				}
			} else if c.Type == html.CommentNode {
				n.RemoveChild(c)
				return
			}
			f(c)
		}
	}
	f(doc)
	buf := bytes.NewBufferString("")
	err = html.Render(buf, doc)
	return buf.String(), err
}

func RedactString(input string) (string, error) {
	r := bufio.NewReader(bytes.NewBufferString(input))
	return Redact(r)
}
