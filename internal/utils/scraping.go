package utils

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func printDOMGoQuery(sel *goquery.Selection, indent int) {
	prefix := strings.Repeat("  ", indent)

	for _, node := range sel.Nodes {
		fmt.Printf("%s<%s>\n", prefix, node.Data)
		for _, attr := range node.Attr {
			fmt.Printf("%s  @%s=%s\n", prefix, attr.Key, attr.Val)
		}
		// Recurse over children
		printDOMGoQuery(sel.Children(), indent+1)
	}
}
