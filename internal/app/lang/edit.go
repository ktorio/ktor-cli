package lang

import (
	"github.com/antlr4-go/antlr/v4"
)

func InsertLnAfter(rewriter *antlr.TokenStreamRewriter, token antlr.Token, indent, text string) {
	rewriter.InsertAfterDefault(token.GetTokenIndex(), "\n"+indent+text)
}
