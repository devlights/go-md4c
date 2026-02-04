package md4c

/*
#cgo LDFLAGS: -lmd4c
#include <stdlib.h>
#include <md4c.h>

// コールバック関数のGo側ラッパー
extern int goEnterBlockCallback(MD_BLOCKTYPE, void*, void*);
extern int goLeaveBlockCallback(MD_BLOCKTYPE, void*, void*);
extern int goEnterSpanCallback(MD_SPANTYPE, void*, void*);
extern int goLeaveSpanCallback(MD_SPANTYPE, void*, void*);
extern int goTextCallback(MD_TEXTTYPE, MD_CHAR*, MD_SIZE, void*);  // constを削除

// C側のコールバック関数（Goの関数を呼び出す）
static int c_enter_block(MD_BLOCKTYPE type, void* detail, void* userdata) {
    return goEnterBlockCallback(type, detail, userdata);
}

static int c_leave_block(MD_BLOCKTYPE type, void* detail, void* userdata) {
    return goLeaveBlockCallback(type, detail, userdata);
}

static int c_enter_span(MD_SPANTYPE type, void* detail, void* userdata) {
    return goEnterSpanCallback(type, detail, userdata);
}

static int c_leave_span(MD_SPANTYPE type, void* detail, void* userdata) {
    return goLeaveSpanCallback(type, detail, userdata);
}

static int c_text(MD_TEXTTYPE type, const MD_CHAR* text, MD_SIZE size, void* userdata) {
    // constキャストを行う
    return goTextCallback(type, (MD_CHAR*)text, size, userdata);
}

// パーサー初期化用のヘルパー
static MD_PARSER create_parser(unsigned int flags) {
    MD_PARSER parser = {
        .abi_version = 0,
        .flags = flags,
        .enter_block = c_enter_block,
        .leave_block = c_leave_block,
        .enter_span = c_enter_span,
        .leave_span = c_leave_span,
        .text = c_text,
        .debug_log = NULL,
        .syntax = NULL
    };
    return parser;
}
*/
import "C"
import (
	"sync"
	"unsafe"
)

// Parser はMD4Cのパーサーラッパー
type Parser struct {
	callbacks Callbacks
	flags     uint32
}

var (
	parserRegistry   = make(map[uintptr]*Parser)
	parserRegistryMu sync.RWMutex
)

// NewParser は新しいパーサーを作成
func NewParser(callbacks Callbacks, flags uint32) *Parser {
	return &Parser{
		callbacks: callbacks,
		flags:     flags,
	}
}

// Parse はマークダウンテキストをパース
func (p *Parser) Parse(markdown string) error {
	// レジストリにパーサーを登録
	parserID := uintptr(unsafe.Pointer(p))
	parserRegistryMu.Lock()
	parserRegistry[parserID] = p
	parserRegistryMu.Unlock()

	defer func() {
		parserRegistryMu.Lock()
		delete(parserRegistry, parserID)
		parserRegistryMu.Unlock()
	}()

	cMarkdown := C.CString(markdown)
	defer C.free(unsafe.Pointer(cMarkdown))

	cParser := C.create_parser(C.uint(p.flags))

	//nolint:unsafeptr
	result := C.md_parse(
		(*C.MD_CHAR)(unsafe.Pointer(cMarkdown)),
		C.MD_SIZE(len(markdown)),
		&cParser,
		unsafe.Pointer(parserID),
	)

	if result != 0 {
		return &ParseError{Code: int(result)}
	}

	return nil
}

// コールバック関数のGo実装（C側から呼ばれる）

//export goEnterBlockCallback
func goEnterBlockCallback(blockType C.MD_BLOCKTYPE, detail unsafe.Pointer, userdata unsafe.Pointer) C.int {
	parserID := uintptr(userdata)
	parserRegistryMu.RLock()
	parser, ok := parserRegistry[parserID]
	parserRegistryMu.RUnlock()

	if !ok || parser.callbacks.EnterBlock == nil {
		return 0
	}

	err := parser.callbacks.EnterBlock(int(blockType), detail)
	if err != nil {
		return -1
	}
	return 0
}

//export goLeaveBlockCallback
func goLeaveBlockCallback(blockType C.MD_BLOCKTYPE, detail unsafe.Pointer, userdata unsafe.Pointer) C.int {
	parserID := uintptr(userdata)
	parserRegistryMu.RLock()
	parser, ok := parserRegistry[parserID]
	parserRegistryMu.RUnlock()

	if !ok || parser.callbacks.LeaveBlock == nil {
		return 0
	}

	err := parser.callbacks.LeaveBlock(int(blockType), detail)
	if err != nil {
		return -1
	}
	return 0
}

//export goEnterSpanCallback
func goEnterSpanCallback(spanType C.MD_SPANTYPE, detail unsafe.Pointer, userdata unsafe.Pointer) C.int {
	parserID := uintptr(userdata)
	parserRegistryMu.RLock()
	parser, ok := parserRegistry[parserID]
	parserRegistryMu.RUnlock()

	if !ok || parser.callbacks.EnterSpan == nil {
		return 0
	}

	err := parser.callbacks.EnterSpan(int(spanType), detail)
	if err != nil {
		return -1
	}
	return 0
}

//export goLeaveSpanCallback
func goLeaveSpanCallback(spanType C.MD_SPANTYPE, detail unsafe.Pointer, userdata unsafe.Pointer) C.int {
	parserID := uintptr(userdata)
	parserRegistryMu.RLock()
	parser, ok := parserRegistry[parserID]
	parserRegistryMu.RUnlock()

	if !ok || parser.callbacks.LeaveSpan == nil {
		return 0
	}

	err := parser.callbacks.LeaveSpan(int(spanType), detail)
	if err != nil {
		return -1
	}
	return 0
}

//export goTextCallback
func goTextCallback(textType C.MD_TEXTTYPE, text *C.MD_CHAR, size C.MD_SIZE, userdata unsafe.Pointer) C.int {
	parserID := uintptr(userdata)
	parserRegistryMu.RLock()
	parser, ok := parserRegistry[parserID]
	parserRegistryMu.RUnlock()

	if !ok || parser.callbacks.Text == nil {
		return 0
	}

	// C文字列をGoの文字列に変換
	goText := C.GoStringN((*C.char)(unsafe.Pointer(text)), C.int(size))

	err := parser.callbacks.Text(int(textType), goText)
	if err != nil {
		return -1
	}
	return 0
}
