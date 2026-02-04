package md4c

import "fmt"

// ParseError はパースエラーを表す
type ParseError struct {
	Code int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("md4c parse error: code %d", e.Code)
}

// Callbacks はマークダウンパース時のコールバック関数群
type Callbacks struct {
	EnterBlock func(blockType int, detail interface{}) error
	LeaveBlock func(blockType int, detail interface{}) error
	EnterSpan  func(spanType int, detail interface{}) error
	LeaveSpan  func(spanType int, detail interface{}) error
	Text       func(textType int, text string) error
}
