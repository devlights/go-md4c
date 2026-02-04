package md4c

/*
#include <md4c.h>
*/
import "C"

// Parser flags
const (
	FlagCollapseWhitespace       = C.MD_FLAG_COLLAPSEWHITESPACE
	FlagPermissiveATXHeaders     = C.MD_FLAG_PERMISSIVEATXHEADERS
	FlagPermissiveURLAutoLinks   = C.MD_FLAG_PERMISSIVEURLAUTOLINKS
	FlagPermissiveEmailAutoLinks = C.MD_FLAG_PERMISSIVEEMAILAUTOLINKS
	FlagNoIndentedCodeBlocks     = C.MD_FLAG_NOINDENTEDCODEBLOCKS
	FlagNoHTMLBlocks             = C.MD_FLAG_NOHTMLBLOCKS
	FlagNoHTMLSpans              = C.MD_FLAG_NOHTMLSPANS
	FlagTables                   = C.MD_FLAG_TABLES
	FlagStrikethrough            = C.MD_FLAG_STRIKETHROUGH
	FlagPermissiveWWWAutoLinks   = C.MD_FLAG_PERMISSIVEWWWAUTOLINKS
	FlagTaskLists                = C.MD_FLAG_TASKLISTS
	FlagLatexMathSpans           = C.MD_FLAG_LATEXMATHSPANS
	FlagWikiLinks                = C.MD_FLAG_WIKILINKS
	FlagUnderline                = C.MD_FLAG_UNDERLINE
	FlagHardSoftBreaks           = C.MD_FLAG_HARD_SOFT_BREAKS

	// 複合フラグ
	FlagPermissiveAutoLinks = C.MD_FLAG_PERMISSIVEAUTOLINKS
	FlagNoHTML              = C.MD_FLAG_NOHTML
)

// Dialect shortcuts
const (
	DialectCommonMark = 0
	DialectGitHub     = FlagPermissiveATXHeaders | FlagPermissiveURLAutoLinks |
		FlagPermissiveWWWAutoLinks | FlagPermissiveEmailAutoLinks |
		FlagTables | FlagStrikethrough | FlagTaskLists
)

// Block types
const (
	BlockDoc   = C.MD_BLOCK_DOC
	BlockQuote = C.MD_BLOCK_QUOTE
	BlockUL    = C.MD_BLOCK_UL
	BlockOL    = C.MD_BLOCK_OL
	BlockLI    = C.MD_BLOCK_LI
	BlockHR    = C.MD_BLOCK_HR
	BlockH     = C.MD_BLOCK_H
	BlockCode  = C.MD_BLOCK_CODE
	BlockHTML  = C.MD_BLOCK_HTML
	BlockP     = C.MD_BLOCK_P
	BlockTable = C.MD_BLOCK_TABLE
	BlockThead = C.MD_BLOCK_THEAD
	BlockTbody = C.MD_BLOCK_TBODY
	BlockTR    = C.MD_BLOCK_TR
	BlockTH    = C.MD_BLOCK_TH
	BlockTD    = C.MD_BLOCK_TD
)

// Span types
const (
	SpanEM               = C.MD_SPAN_EM
	SpanStrong           = C.MD_SPAN_STRONG
	SpanA                = C.MD_SPAN_A
	SpanImg              = C.MD_SPAN_IMG
	SpanCode             = C.MD_SPAN_CODE
	SpanDel              = C.MD_SPAN_DEL
	SpanLatexMath        = C.MD_SPAN_LATEXMATH
	SpanLatexMathDisplay = C.MD_SPAN_LATEXMATH_DISPLAY
	SpanWikiLink         = C.MD_SPAN_WIKILINK
	SpanU                = C.MD_SPAN_U
)

// Text types
const (
	TextNormal    = C.MD_TEXT_NORMAL
	TextNullChar  = C.MD_TEXT_NULLCHAR
	TextBR        = C.MD_TEXT_BR
	TextSoftBR    = C.MD_TEXT_SOFTBR
	TextEntity    = C.MD_TEXT_ENTITY
	TextCode      = C.MD_TEXT_CODE
	TextHTML      = C.MD_TEXT_HTML
	TextLatexMath = C.MD_TEXT_LATEXMATH
)
