# GoMD4C

Go言語用の[MD4C](https://github.com/mity/md4c)ラッパーライブラリです。

高速なC言語製マークダウンパーサー [MD4C](https://github.com/mity/md4c) をGoから利用できるようにします。

内部で cgo を使っていますので、Cコンパイラが必須です。

## 前提条件

### 1. MD4Cライブラリ

このライブラリを使用するには、**MD4Cライブラリが事前にインストールされている必要があります**。

#### MD4Cのインストール

##### Ubuntu/Debian
```bash
sudo apt-get install libmd4c-dev
```

##### ソースからビルド
```bash
git clone https://github.com/mity/md4c.git
cd md4c
mkdir build && cd build
cmake ..
make
sudo make install
sudo ldconfig
```

インストール後、以下のコマンドでMD4Cが正しくインストールされていることを確認してください：

```bash
pkg-config --cflags --libs md4c
ldconfig -p | grep md4c
```

### 2. Cgoの有効化

**このライブラリはCgoを使用しています。** ビルド時には `CGO_ENABLED=1` が設定されている必要があります。

## インストール

```bash
go get github.com/devlights/go-md4c
```

## 基本的な使い方

### シンプルなHTML変換

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/devlights/md4c"
)

func main() {
    markdown := `# Hello World
    
This is a **bold** text and this is *italic*.

- Item 1
- Item 2
- Item 3
`

    // GitHub Flavored Markdownとして変換
    html, err := md4c.RenderHTML(markdown, md4c.DialectGitHub)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(html)
}
```

### カスタムフラグを使用したHTML変換

```go
// HTMLレンダラーを作成
renderer := md4c.NewHTMLRenderer(
    md4c.FlagTables | md4c.FlagStrikethrough | md4c.FlagTaskLists,
    md4c.HTMLFlagXHTML,
)

html, err := renderer.Render(markdown)
if err != nil {
    log.Fatal(err)
}

fmt.Println(html)
```

### 低レベルパーサー（コールバックベース）

```go
callbacks := md4c.Callbacks{
    EnterBlock: func(blockType int, detail any) error {
        switch blockType {
        case md4c.BlockH:
            fmt.Println("見出し開始")
        case md4c.BlockP:
            fmt.Println("段落開始")
        }
        return nil
    },
    Text: func(textType int, text string) error {
        fmt.Print(text)
        return nil
    },
}

parser := md4c.NewParser(callbacks, md4c.DialectGitHub)
err := parser.Parse(markdown)
if err != nil {
    log.Fatal(err)
}
```

## 利用可能なフラグ

### パーサーフラグ

```go
md4c.FlagTables                   // テーブル拡張を有効化
md4c.FlagStrikethrough            // 打ち消し線を有効化
md4c.FlagTaskLists                // タスクリストを有効化
md4c.FlagWikiLinks                // Wikiリンクを有効化
md4c.FlagUnderline                // 下線を有効化
md4c.FlagPermissiveURLAutoLinks   // URL自動リンクを有効化
md4c.FlagPermissiveEmailAutoLinks // メール自動リンクを有効化
md4c.FlagPermissiveWWWAutoLinks   // WWW自動リンクを有効化
```

### HTMLレンダラーフラグ

```go
md4c.HTMLFlagDebug            // デバッグ出力を有効化
md4c.HTMLFlagVerbatimEntities // HTMLエンティティをそのまま出力
md4c.HTMLFlagSkipUTF8BOM      // UTF-8 BOMをスキップ
md4c.HTMLFlagXHTML            // XHTML形式で出力
```

### 定義済みダイアレクト

```go
md4c.DialectCommonMark // CommonMark標準
md4c.DialectGitHub     // GitHub Flavored Markdown
```

## サンプル

### HTMLレンダリングのサンプル

```bash
go run examples/html/html_example.go
```

- 基本的なHTML変換
- XHTML形式での出力
- カスタムパーサーフラグの使用
- テーブル、タスクリスト、打ち消し線などの拡張機能

### パーサーのサンプル

```bash
go run examples/parser/parser_example.go
```

- コールバックベースの低レベルパーサー
- ブロック要素とスパン要素の処理
- テキストコンテンツの取得

## コマンドラインツール

マークダウンをHTMLに変換するシンプルなコマンドラインツールも置いています。

### ビルド

```bash
cd cmd/md2html
go build -o md2html main.go
```

### 使用方法

```bash
# 標準入力から変換
echo '# Hello **World**' | ./md2html

# ファイルから変換
cat README.md | ./md2html

# ファイルに出力
cat input.md | ./md2html > output.html
```

出力例：
```html
<h1>Hello <strong>World</strong></h1>
```


## 参考リンク

- [MD4C本家](https://github.com/mity/md4c)
- [CommonMark仕様](https://commonmark.org/)
- [GitHub Flavored Markdown仕様](https://github.github.com/gfm/)
