我们在编写爬虫软件或者抓取某一个单一网页的时候，一般都是使用正则表达式技术获取需要的内容。这种方式比较灵活，但是使用门槛较高，而且效率不高，很多不了解正则表达式语法的小伙伴面对一连串符号看起来就是像在看天书一样。
  
今天给大家分享的是golang官方开源的net/html包， 这个包给我们封装了解析的html元素用到的基础方法。下载地址 [https://github.com/golang/net](https://github.com/golang/net)
>不能直接使用git clone 的方式下载源代码，必须使用go get 命令获取

```
<!DOCTYPE html>
<html>
<head>
	<title>This is a demo</title>
</head>
<body>
这是body下面的示例一

<a href="http://shang.qq.com/wpa/qunwpa?idkey=1720f7b75f19d952b80e10549ce35c6cc922c25b6505cd6f6680ac4fc7259484">欢迎加入 dogo 技术交流群：437274005 点击右侧按钮快捷加入</a>
<img src="https://github.com/wuciyou/dogo/blob/master/example/web/static/img/dogo.png">

这是body下面的示例二
</body>
</html>
```

 
```
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

```
  
  
欢迎加入 dogo 技术交流群：437274005 点击右侧按钮快捷加入
[![dogo交流群](http://pub.idqqimg.com/wpa/images/group.png)](http://shang.qq.com/wpa/qunwpa?idkey=1720f7b75f19d952b80e10549ce35c6cc922c25b6505cd6f6680ac4fc7259484)