package main

import (
	"fmt"
	"log"

	"github.com/devlights/md4c"
)

func main() {
	markdown := `# Hello World

This is a **bold** text and this is *italic*.

## Code Example

` + "```go" + `
package main

func main() {
    fmt.Println("Hello")
}
` + "```" + `

- Item 1
- Item 2
- Item 3
`

	callbacks := md4c.Callbacks{
		EnterBlock: func(blockType int, detail interface{}) error {
			switch blockType {
			case md4c.BlockH:
				fmt.Println("Enter: Header")
			case md4c.BlockP:
				fmt.Println("Enter: Paragraph")
			case md4c.BlockCode:
				fmt.Println("Enter: Code Block")
			case md4c.BlockUL:
				fmt.Println("Enter: Unordered List")
			case md4c.BlockLI:
				fmt.Println("Enter: List Item")
			}
			return nil
		},
		LeaveBlock: func(blockType int, detail interface{}) error {
			switch blockType {
			case md4c.BlockH:
				fmt.Println("Leave: Header")
			case md4c.BlockP:
				fmt.Println("Leave: Paragraph")
			case md4c.BlockCode:
				fmt.Println("Leave: Code Block")
			case md4c.BlockUL:
				fmt.Println("Leave: Unordered List")
			case md4c.BlockLI:
				fmt.Println("Leave: List Item")
			}
			return nil
		},
		EnterSpan: func(spanType int, detail interface{}) error {
			switch spanType {
			case md4c.SpanStrong:
				fmt.Print("<strong>")
			case md4c.SpanEM:
				fmt.Print("<em>")
			case md4c.SpanCode:
				fmt.Print("<code>")
			}
			return nil
		},
		LeaveSpan: func(spanType int, detail interface{}) error {
			switch spanType {
			case md4c.SpanStrong:
				fmt.Print("</strong>")
			case md4c.SpanEM:
				fmt.Print("</em>")
			case md4c.SpanCode:
				fmt.Print("</code>")
			}
			return nil
		},
		Text: func(textType int, text string) error {
			fmt.Print(text)
			return nil
		},
	}

	parser := md4c.NewParser(callbacks, md4c.DialectGitHub)

	if err := parser.Parse(markdown); err != nil {
		log.Fatal(err)
	}
}
