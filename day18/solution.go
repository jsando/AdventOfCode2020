package day18

import (
	"fmt"
	"strings"

	"github.com/jsando/aoc2020/helpers"
)

// Run day 18.
func Run(inputPath string) {
	part1 := 0
	part2 := 0
	scanner := helpers.NewScanner(inputPath)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		t := newTokenizer(line)
		value := t.eval()
		fmt.Printf("eval1: %s = %d\n", line, value)
		part1 += value

		t = newTokenizer(line)
		value = t.eval2()
		fmt.Printf("eval2: %s = %d\n", line, value)
		part2 += value
	}
	fmt.Printf("part 1: %d\n", part1) // 209335026987
	fmt.Printf("part 2: %d\n", part2) // 33331817392479
}

type tokenizer struct {
	input     string
	idx       int
	tokenType byte
	number    int
}

const (
	tokenEOF    byte = 0
	tokenPlus   byte = '+'
	tokenSplat  byte = '*'
	tokenLParen byte = '('
	tokenRParen byte = ')'
	tokenNumber byte = '0'
)

func newTokenizer(input string) *tokenizer {
	return &tokenizer{input: input}
}

func (t *tokenizer) backup() {
	if t.tokenType != tokenEOF {
		t.idx--
	}
}

func (t *tokenizer) next() byte {
	for ; t.idx < len(t.input) && t.input[t.idx] == ' '; t.idx++ {
	}
	if t.idx >= len(t.input) {
		t.tokenType = tokenEOF
		return t.tokenType
	}
	ch := t.input[t.idx]
	t.idx++
	switch ch {
	case '+':
		t.tokenType = tokenPlus
	case '*':
		t.tokenType = tokenSplat
	case '(':
		t.tokenType = tokenLParen
	case ')':
		t.tokenType = tokenRParen
	default:
		t.number = int(ch - '0')
		t.tokenType = tokenNumber
	}
	return t.tokenType
}

func (t *tokenizer) eval() int {
	value := t.subExpr()
	for {
		op := t.next()
		if op == tokenEOF || op == tokenRParen {
			break
		}
		rhs := t.subExpr()
		switch op {
		case tokenPlus:
			value += rhs
		case tokenSplat:
			value *= rhs
		default:
			panic(fmt.Sprintf("unexpected token: '%c'", t.tokenType))
		}
	}
	return value
}

func (t *tokenizer) subExpr() int {
	tok := t.next()
	if tok == tokenLParen {
		return t.eval()
	}
	if tok == tokenNumber {
		return t.number
	}
	panic(fmt.Sprintf("unexpected token: %c", tok))
}

func (t *tokenizer) eval2() int {
	value := t.subExpr2()
	for {
		op := t.next()
		if op == tokenEOF || op == tokenRParen {
			break
		}
		rhs := t.subExpr2()
		if op == tokenSplat {
			value *= rhs
		} else {
			panic(fmt.Sprintf("unexpected token: '%c'", t.tokenType))
		}
	}
	return value
}

func (t *tokenizer) subExpr2() int {
	value := 0
	tok := t.next()
	if tok == tokenLParen {
		value = t.eval2()
	} else if tok == tokenNumber {
		value = t.number
	} else {
		panic(fmt.Sprintf("unexpected token: '%c'", tok))
	}
	// greedy consume plus
	peek := t.next()
	if peek == tokenPlus {
		value2 := t.subExpr2()
		value += value2
	} else {
		t.backup()
	}
	return value
}
