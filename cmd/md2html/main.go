package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/devlights/md4c"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	html, err := md4c.RenderHTML(string(b), md4c.DialectGitHub)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html)
}
