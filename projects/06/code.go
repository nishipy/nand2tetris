package main

func Dest(s string) string {
	var out string
	switch s {
	case "":
		out = "000"
	case "M":
		out = "001"
	case "D":
		out = "010"
	case "MD":
		out = "011"
	case "A":
		out = "100"
	case "AM":
		out = "101"
	case "AD":
		out = "110"
	case "AMD":
		out = "111"
	}

	return out
}

func Comp(s string) string {
	var out string
	switch s {
	//a=0
	case "0":
		out = "0101010"
	case "1":
		out = "0111111"
	case "-1":
		out = "0111010"
	case "D":
		out = "0001100"
	case "A":
		out = "0110000"
	case "!D":
		out = "0001101"
	case "!A":
		out = "0110001"
	case "-D":
		out = "0001111"
	case "-A":
		out = "0110011"
	case "D+1":
		out = "0011111"
	case "A+1":
		out = "0110111"
	case "D-1":
		out = "0001110"
	case "A-1":
		out = "0110010"
	case "D+A":
		out = "0000010"
	case "D-A":
		out = "0010011"
	case "A-D":
		out = "0000111"
	case "D&A":
		out = "0000000"
	case "D|A":
		out = "0010101"

	// a=1
	case "M":
		out = "1110000"
	case "!M":
		out = "1110001"
	case "-M":
		out = "1110011"
	case "M+1":
		out = "1110111"
	case "M-1":
		out = "1110010"
	case "D+M":
		out = "1000010"
	case "D-M":
		out = "1010011"
	case "M-D":
		out = "1000111"
	case "D&M":
		out = "1000000"
	case "D|M":
		out = "1010101"
	}

	return out
}

func Jump(s string) string {
	var out string
	switch s {
	case "":
		out = "000"
	case "JGT":
		out = "001"
	case "JEQ":
		out = "010"
	case "JGE":
		out = "011"
	case "JLT":
		out = "100"
	case "JNE":
		out = "101"
	case "JLE":
		out = "110"
	case "JMP":
		out = "111"
	}
	return out
}
