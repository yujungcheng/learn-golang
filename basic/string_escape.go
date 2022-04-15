package main

import (
	"fmt"
	"html"
	"net/url"
)

func main() {

	// escape html. html.EscapeString
	// encode a string which can be safely placed in HTML text.
	// it encapes <, >, &, ' and "
	const s = `"Foo's Bar" <myemail@example.com>`
	fmt.Println(html.EscapeString(s))

	// escape url. url.PathEscape
	// ref: https://yourbasic.org/golang/multiline-string/
	const u = `Foo's Bar?=`
	fmt.Println(url.PathEscape(u))
}
