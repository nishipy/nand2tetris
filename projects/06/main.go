package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	filename := os.Args[1]
	fp, _ := os.Open(filename)
	scanner := bufio.NewScanner(fp)
	hackfile := strings.TrimSuffix(filename, ".asm") + ".hack"
	writerfp, _ := os.OpenFile(hackfile, os.O_WRONLY|os.O_CREATE, 0644)
	writer := bufio.NewWriter(writerfp)

	symbolTable := scanSymbol(NewParser(scanner))
	fp.Seek(0, 0)
	scanner = bufio.NewScanner(fp)
	ramAddr := 0x0010
	p := NewParser(scanner)
	for p.HasMoreCommands() {
		p.Advance()
		var output string
		switch p.CommandType() {
		case A_CMD:
			output = "0"
			symbol := p.Symbol()
			addr, err := strconv.Atoi(symbol)
			if err == nil {
				//number-like string
				output += int2bin(addr)
			} else {
				//symbol
				if symbolTable.Contains(symbol) {
					// known symbol
					addr = symbolTable.GetAddress(symbol)
					output += int2bin(addr)
				} else {
					// new symbol
					symbolTable.AddEntry(symbol, ramAddr)
					output += int2bin(ramAddr)
					ramAddr++
				}
			}
			fmt.Fprintln(writer, output)

		case C_CMD:
			//fmt.Println(p.Comp(), p.Dest(), p.Jump())
			output = "111"
			comp := Comp(p.Comp())
			output += comp

			dest := Dest(p.Dest())
			output += dest

			jump := Jump(p.Jump())
			output += jump

			fmt.Fprintln(writer, output)

		case L_CMD:
			// do nothing
		}
		writer.Flush()
	}
}

// 15bit int -> 000 0000 0000 0000
func int2bin(num int) string {
	var bin string
	for i := 1 << 14; i > 0; i = i >> 1 {
		if i&num != 0 {
			bin += "1"
		} else {
			bin += "0"
		}
	}
	return bin
}

func scanSymbol(p *Parser) SymbolTable {
	st := NewSymbolTable()
	romAddr := 0

	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case A_CMD, C_CMD:
			romAddr++
		case L_CMD:
			st.AddEntry(p.Symbol(), romAddr)
		}
	}

	return st
}

//Reference:
// - https://github.com/hirak/Assembler/blob/master/Assembler.go
