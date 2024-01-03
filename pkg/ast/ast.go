package ast

import (
	"github.com/israel-duff/go-monkey-interpreter/pkg/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // The token.LEt token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // THE Token.IDENt token
	Value string
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (id *Identifier) statementNode()       {}
func (id *Identifier) TokenLiteral() string { return id.Token.Literal }

// RETURN STATEMENT
type ReturnStatement struct {
	Token       token.Token // represeting the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
