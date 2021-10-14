package token

type TokenType string // stringにしているのは簡単にデバッグできるから.
type Token struct {
	Type    TokenType
	Literal string // TODO: stringで良いかどうか?
}

// トークンが未知であることを表す
// ファイルの終わりを表す
// これがあることによって, 構文解析器に終了のタイミングを伝えられる.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

// 識別子, リテラル
const (
	IDENT = "IDENT"
	INT   = "INT"
)

// 記号（演算子）
const (
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
)

// 記号（デリミタ）
const (
	COMMA     = ","
	SEMICOLON = ";"
)

// 記号（かっこ）
const (
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

// キーワード
const (
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
