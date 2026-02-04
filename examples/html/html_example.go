package main

import (
	"fmt"
	"log"

	"github.com/devlights/md4c"
)

func main() {
	markdown := `# Hello World

This is a **bold** text and this is *italic*.

## Features

- GitHub Flavored Markdown
- Tables support
- ~~Strikethrough~~
- Task lists

## Code Example

` + "```go" + `
package main

func main() {
    fmt.Println("Hello, MD4C!")
}
` + "```" + `

## Table

| Name | Age | City |
|------|-----|------|
| Alice | 30 | Tokyo |
| Bob | 25 | Osaka |

## Task List

- [x] Completed task
- [ ] Pending task
`

	// GitHub Flavored Markdownとしてレンダリング
	html, err := md4c.RenderHTML(markdown, md4c.DialectGitHub)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== HTML Output ===")
	fmt.Println(html)

	// XHTML形式で出力
	html2, err := md4c.RenderHTMLWithFlags(
		markdown,
		md4c.DialectGitHub,
		md4c.HTMLFlagXHTML,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== XHTML Output ===")
	fmt.Println(html2)

	// カスタムパーサーフラグ
	renderer := md4c.NewHTMLRenderer(
		md4c.FlagTables|md4c.FlagStrikethrough|md4c.FlagTaskLists,
		0,
	)

	html3, err := renderer.Render(markdown)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== Custom Parser Flags Output ===")
	fmt.Println(html3)
}
