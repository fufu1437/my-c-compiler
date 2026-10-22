package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO (compiler-pipeline): implement per the lesson description.

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		// fmt.Println("[DEBUG]", line)

		tks := Tokenizer(line)
		PrintTokens(tks)
		fmt.Println()
		// files := strings.Split(line, ".")
		// switch files[len(files)-1] {
		// case "c":
		// 	fmt.Println("SOURCE")
		// case "i":
		// 	fmt.Println("PREPROCESSED")
		// case "s":
		// 	fmt.Println("ASSEMBLY")
		// case "o":
		// 	fmt.Println("OBJECT")
		// case "a":
		// 	fmt.Println("STATIC_LIB")
		// case "so", "dylib", "dll":
		// 	fmt.Println("SHARED_LIB")
		// default:
		// 	fmt.Println("EXECUTABLE")
		// }
	}
}

type TokenType string

const (
	KW  = "KW"
	ID  = "ID"
	INT = "INT"
	OP  = "OP"
)

var KeyWords = map[string]bool{
	"int":    true,
	"if":     true,
	"else":   true,
	"while":  true,
	"return": true,
}

var Operators = map[string]bool{
	"+":  true,
	"-":  true,
	"*":  true,
	"/":  true,
	"==": true,
	"!=": true,
	"<":  true,
	">":  true,
	"=":  true,
	"(":  true,
	")":  true,
	"{":  true,
	"}":  true,
	";":  true,
	"!":  true,
}

type Token struct {
	Type      TokenType
	Value     string
	Line, Cow uint64
}

func Tokenizer(code string) []Token {
	tks := make([]Token, 0, 10)
	codes := []rune(code)
	len_codes := len(codes)
	var line uint64 = 0
	var cow uint64 = 0

	for i := 0; i < len_codes; i++ {
		if isSpace(codes[i]) {
			for isSpace(codes[i]) {
				if codes[i] == '\n' {
					line++
					cow = 0
				} else {
					cow++
				}
				i++
			}
			// i--

		}

		if isAsciiLetter(codes[i]) || codes[i] == '_' {
			var buf strings.Builder
			buf.Grow(5)
			for i < len_codes && (isAsciiLetter(codes[i]) || isNumber(codes[i]) || codes[i] == '_') {
				buf.WriteRune(codes[i])
				i++
				cow++
			}
			i--
			cow--
			if isKeyWords(buf.String()) {
				tks = append(tks, Token{
					Value: buf.String(),
					Type:  KW,
					Line:  line,
					Cow:   cow,
				})
			} else {
				tks = append(tks, Token{
					Value: buf.String(),
					Type:  ID,
					Line:  line,
					Cow:   cow,
				})
			}

		} else if isNumber(codes[i]) {
			var buf strings.Builder
			buf.Grow(5)
			for isNumber(codes[i]) {
				buf.WriteRune(codes[i])
				i++
				cow++
			}
			i--
			cow--
			tks = append(tks, Token{
				Value: buf.String(),
				Type:  INT,
				Line:  line,
				Cow:   cow,
			})
		} else if isOperators(string(codes[i])) {

			var buf strings.Builder
			buf.Grow(2)

			// if (codes[i] == '!' || codes[i] == '=') && codes[i+1] == '=' {

			// } else {

			// }
			switch codes[i] {
			case '!', '=':
				// fmt.Println("[DEBUG]", string(codes[i]))
				switch codes[i+1] {
				case '=':
					buf.WriteRune(codes[i])
					// fmt.Println("asdasd", string(codes[i]))
					buf.WriteRune(codes[i+1])
					i++
					cow++
				default:
					buf.WriteRune(codes[i])
				}
			default:
				buf.WriteRune(codes[i])
			}
			// for isOperators(string(codes[i])) {
			// 	buf.WriteRune(codes[i])
			// 	i++
			// }

			tks = append(tks, Token{
				Value: buf.String(),
				Type:  OP,
				Line:  line,
				Cow:   cow,
			})
		}

		cow++
	}
	return tks
}

func PrintTokens(tks []Token) {
	var buf []string
	for _, tk := range tks {
		buf = append(buf, string(tk.Type+":"+TokenType(tk.Value)))
	}
	fmt.Print(strings.Join(buf, " "))
}

func isKeyWords(c string) bool {
	_, is := KeyWords[c]
	return is
}

func isOneOperators(c string) bool {
	if len(c) > 1 {
		return false
	}
	return isOperators(c)
}

func isOperators(c string) bool {
	_, is := Operators[c]
	return is
}

func isSpace(c rune) bool {
	return c == ' ' || c == '\b' || c == '\t' || c == '\n'
}

func isNumber(c rune) bool {
	return c >= '0' && c <= '9'
}

func isAsciiLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}
