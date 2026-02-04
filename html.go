package md4c

/*
#cgo LDFLAGS: -lmd4c-html
#include <md4c-html.h>
#include <stdlib.h>
#include <string.h>

// HTML出力を文字列として受け取るための構造体
typedef struct {
    char* data;
    size_t size;
    size_t capacity;
} output_buffer_t;

// 出力バッファを初期化
static output_buffer_t* create_output_buffer() {
    output_buffer_t* buf = (output_buffer_t*)malloc(sizeof(output_buffer_t));
    buf->capacity = 4096;
    buf->size = 0;
    buf->data = (char*)malloc(buf->capacity);
    return buf;
}

// 出力バッファに追記
static void append_output(const MD_CHAR* text, MD_SIZE size, void* userdata) {
    output_buffer_t* buf = (output_buffer_t*)userdata;

    // 容量が足りない場合は拡張
    while (buf->size + size > buf->capacity) {
        buf->capacity *= 2;
        buf->data = (char*)realloc(buf->data, buf->capacity);
    }

    memcpy(buf->data + buf->size, text, size);
    buf->size += size;
}

// 出力バッファを解放
static void free_output_buffer(output_buffer_t* buf) {
    if (buf) {
        if (buf->data) {
            free(buf->data);
        }
        free(buf);
    }
}

// ラッパー関数
static int render_html(const MD_CHAR* input, MD_SIZE input_size,
                       output_buffer_t* buf, unsigned parser_flags, unsigned renderer_flags) {
    return md_html(input, input_size, append_output, (void*)buf, parser_flags, renderer_flags);
}
*/
import "C"
import (
	"unsafe"
)

// HTML renderer flags
const (
	HTMLFlagDebug            = C.MD_HTML_FLAG_DEBUG
	HTMLFlagVerbatimEntities = C.MD_HTML_FLAG_VERBATIM_ENTITIES
	HTMLFlagSkipUTF8BOM      = C.MD_HTML_FLAG_SKIP_UTF8_BOM
	HTMLFlagXHTML            = C.MD_HTML_FLAG_XHTML
)

// HTMLRenderer はマークダウンをHTMLに変換するレンダラー
type HTMLRenderer struct {
	parserFlags   uint32
	rendererFlags uint32
}

// NewHTMLRenderer は新しいHTMLレンダラーを作成
func NewHTMLRenderer(parserFlags uint32, rendererFlags uint32) *HTMLRenderer {
	return &HTMLRenderer{
		parserFlags:   parserFlags,
		rendererFlags: rendererFlags,
	}
}

// Render はマークダウンをHTMLに変換
func (r *HTMLRenderer) Render(markdown string) (string, error) {
	cMarkdown := C.CString(markdown)
	defer C.free(unsafe.Pointer(cMarkdown))

	// 出力バッファを作成
	buf := C.create_output_buffer()
	defer C.free_output_buffer(buf)

	// HTMLレンダリング（ラッパー関数を使用）
	result := C.render_html(
		(*C.MD_CHAR)(unsafe.Pointer(cMarkdown)),
		C.MD_SIZE(len(markdown)),
		buf,
		C.uint(r.parserFlags),
		C.uint(r.rendererFlags),
	)

	if result != 0 {
		return "", &ParseError{Code: int(result)}
	}

	// C文字列をGoの文字列に変換
	html := C.GoStringN(buf.data, C.int(buf.size))

	return html, nil
}

// RenderHTML はマークダウンをHTMLに変換する便利関数
func RenderHTML(markdown string, parserFlags uint32) (string, error) {
	renderer := NewHTMLRenderer(parserFlags, 0)
	return renderer.Render(markdown)
}

// RenderHTMLWithFlags はパーサーとレンダラーの両方のフラグを指定してHTMLに変換
func RenderHTMLWithFlags(markdown string, parserFlags uint32, rendererFlags uint32) (string, error) {
	renderer := NewHTMLRenderer(parserFlags, rendererFlags)
	return renderer.Render(markdown)
}
