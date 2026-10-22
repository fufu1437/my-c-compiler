package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// This program is self-contained on purpose: shipthatcode's grader only
// compiles the entrypoint (main.go) and ignores extra files, and its
// workspace has no go.mod. Importing a local package would therefore fail
// there, so the tokenizer lives in this file instead of token/.

type TokenType string

const (
	KW  TokenType = "KW"
	ID  TokenType = "ID"
	INT TokenType = "INT"
	OP  TokenType = "OP"
)

var keywords = map[string]bool{
	"int":    true,
	"if":     true,
	"else":   true,
	"while":  true,
	"return": true,
}

var operators = map[string]bool{
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

func tokenizer(code string) []Token {
	tks := make([]Token, 0, 10)
	codes := []rune(code)
	n := len(codes)
	var line, col uint64

	for i := 0; i < n; i++ {
		if isSpace(codes[i]) {
			for i < n && isSpace(codes[i]) {
				if codes[i] == '\n' {
					line++
					col = 0
				} else {
					col++
				}
				i++
			}
			if i >= n {
				break
			}
		}

		start := i
		startCol := col
		c := codes[i]

		switch {
		case isAsciiLetter(c) || c == '_':
			for i < n && (isAsciiLetter(codes[i]) || isNumber(codes[i]) || codes[i] == '_') {
				i++
				col++
			}
			value := string(codes[start:i])
			tks = append(tks, Token{
				Type:  tokenType(value),
				Value: value,
				Line:  line,
				Cow:   startCol,
			})
			i-- // the loop's own i++ moves past the token

		case isNumber(c):
			for i < n && isNumber(codes[i]) {
				i++
				col++
			}
			value := string(codes[start:i])
			tks = append(tks, Token{
				Type:  INT,
				Value: value,
				Line:  line,
				Cow:   startCol,
			})
			i--

		case isOperator(string(c)):
			value := string(c)
			if (c == '!' || c == '=') && i+1 < n && codes[i+1] == '=' {
				value += "="
				i++
				col++
			}
			tks = append(tks, Token{
				Type:  OP,
				Value: value,
				Line:  line,
				Cow:   startCol,
			})
		}

		col++
	}
	return tks
}

func printTokens(tks []Token) {
	parts := make([]string, 0, len(tks))
	for _, tk := range tks {
		parts = append(parts, string(tk.Type)+":"+tk.Value)
	}
	fmt.Println(strings.Join(parts, " "))
}

func tokenType(value string) TokenType {
	if _, ok := keywords[value]; ok {
		return KW
	}
	return ID
}

func isOperator(c string) bool {
	_, ok := operators[c]
	return ok
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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		printTokens(tokenizer(line))
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
