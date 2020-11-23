package main

import (
	"bufio"
	"strings"
)

type CommandType int

const (
	INVALID_CMD CommandType = 0
	A_CMD       CommandType = 1
	C_CMD       CommandType = 2
	L_CMD       CommandType = 3
)

type Parser struct {
	scanner    *bufio.Scanner
	currentCmd string
	dest       string
	comp       string
	jump       string
}

type ParserInterface interface {
	HasMoreCommands() bool
	Advance()
	CommandType() CommandType
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, "", "", "", ""}
}

//scanner.Scan(): １行読み込む
//scanner.Text(): 読み込んだトークンを文字列として取り出す
func (p *Parser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *Parser) Advance() {
	p.currentCmd = p.scanner.Text()
	if len(p.currentCmd) <= 0 {
		return
	}

	//コメントを排除
	tokens := strings.SplitN(p.currentCmd, "//", 2)
	if len(tokens) > 0 {
		p.currentCmd = tokens[0]
	}
	//スペースを排除
	p.currentCmd = strings.TrimSpace(p.currentCmd)
}

func (p *Parser) CommandType() CommandType {
	if len(p.currentCmd) == 0 {
		return INVALID_CMD
	}
	if p.currentCmd[0] == '@' {
		// @Xxx
		return A_CMD
	}
	if p.currentCmd[0] == '(' {
		// (Xxx)
		return L_CMD
	}

	p.analyze()
	return C_CMD
}

func (p *Parser) Symbol() string {
	if p.CommandType() == A_CMD {
		return strings.TrimPrefix(p.currentCmd, "@")
	}
	if p.CommandType() == L_CMD {
		return strings.Trim(p.currentCmd, "()")
	}
	return ""
}

// parse "dest=comp;jump"
func (p *Parser) analyze() {
	// "dest=comp", "jump"
	tokens := strings.SplitN(p.currentCmd, ";", 2)
	destcomp := tokens[0]
	if len(tokens) == 2 {
		p.jump = tokens[1]
	} else {
		p.jump = ""
	}

	// "dest", "comp"
	tokens = strings.SplitN(destcomp, "=", 2)
	if len(tokens) == 2 {
		p.dest = tokens[0]
		p.comp = tokens[1]
	} else {
		p.dest = ""
		p.comp = destcomp
	}
}

func (p *Parser) Dest() string {
	return p.dest
}

func (p *Parser) Comp() string {
	return p.comp
}

func (p *Parser) Jump() string {
	return p.jump
}

//Reference:
// - https://github.com/hirak/Assembler/blob/master/parser.go
