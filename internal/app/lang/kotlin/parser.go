package kotlin

import (
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
)

func NewParser(fp string) (*parser.KotlinParser, error) {
	input, err := antlr.NewFileStream(fp)

	if err != nil {
		return nil, err
	}

	lexer := parser.NewKotlinLexer(&input.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return parser.NewKotlinParser(stream), nil
}
