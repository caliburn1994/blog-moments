package main

import (
	"github.com/russross/blackfriday"
)


func AddMarkdown(context string)string {
	tmp := blackfriday.MarkdownCommon([]byte(context))
	return string(tmp)
}


