package main

import (
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./demo.html")
	if err != nil {
		log.Panic(err)
	}
	doc, err := html.Parse(f)
	if err != nil {
		log.Panic(err)
	}

	// log.Printf("doc:%+v \n ", doc)
	var parse func(*html.Node)
	parse = func(n *html.Node) {
		// log.Printf("node type:%d  ", n.Type)
		// switch n.Type {
		// case html.ErrorNode:
		// 	log.Printf("ErrorNode(%p):%+v", n, n)
		// case html.TextNode:
		// 	log.Printf("TextNode(%p):%+v", n, n)
		// case html.DocumentNode:
		// 	log.Printf("DocumentNode(%p):%+v", n, n)
		// case html.ElementNode:
		// 	log.Printf("ElementNode(%p):%+v", n, n)
		// case html.CommentNode:
		// 	log.Printf("CommentNode(%p):%+v", n, n)
		// case html.DoctypeNode:
		// 	log.Printf("DoctypeNode(%p):%+v", n, n)
		// }

		if n.Type == html.ElementNode && n.Data == "a" {
			if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
				log.Printf("href:%s, text:%s \n", n.Attr[0].Val, n.FirstChild.Data)
			} else {
				log.Printf("href:%s \n", n.Attr[0].Val)
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c)
		}
	}
	parse(doc)
}
