package greb

import (
	"github.com/alecthomas/participle/v2"
)

var _InternalParser = &Parser{}

type Parser struct {
	parser *participle.Parser
}

func (p *Parser) Parse(rawText string) (*Greb, error) {
	ast := &Greb{}
	if err := p.parser.ParseString("", rawText, ast); err != nil {
		return nil, err
	}
	return ast, nil
}

func init() {
	grammar := &Greb{}
	parser := participle.MustBuild(
		grammar,
		participle.Unquote())

	_InternalParser = &Parser{
		parser: parser,
	}
}
