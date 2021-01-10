package day19

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Run day 19.
func Run(inputPath string) {
	bytes, err := ioutil.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	sections := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
	matcher := matcher(parseRules(sections[0]))
	matcher.printRules()

	linesToMatch := strings.Split(sections[1], "\n")
	fmt.Printf("Part 1: %d\n", matcher.countMatchingLines(linesToMatch)) // 171

	matcher[8] = &rule{ruleID: 8, sequences: [][]int{{42}, {42, 8}}}
	matcher[11] = &rule{ruleID: 11, sequences: [][]int{{42, 31}, {42, 11, 31}}}
	fmt.Printf("Part 2: %d\n", matcher.countMatchingLines(linesToMatch)) // 369
}

type matcher map[int]*rule

func (m matcher) printRules() {
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := 0; i < len(keys); i++ {
		rule := m[keys[i]]
		if rule != nil {
			fmt.Printf("Rule %s\n", rule)
		}
	}
}

func (m matcher) countMatchingLines(lines []string) int {
	count := 0
	for _, line := range lines {
		if m.matches(line) {
			count++
		}
	}
	return count
}

func (m matcher) matches(input string) bool {
	stream := &stringReader{text: input}
	match := m.matchRule(stream, m[0], []int{})
	return match && stream.consumed()
}

func (m matcher) matchRule(s *stringReader, r *rule, followedBy []int) bool {
	mark := s.offset
	if r.letter != 0 {
		if s.next() == r.letter {
			if m.matchSequence(s, followedBy) {
				return true
			}
		}
	} else {
		for _, seq := range r.sequences {
			fullSequence := append(seq, followedBy...)
			if m.matchSequence(s, fullSequence) {
				return true
			}
			s.offset = mark
		}
	}
	s.offset = mark
	return false
}

func (m matcher) matchSequence(s *stringReader, seq []int) bool {
	if len(seq) > 0 {
		return m.matchRule(s, m[seq[0]], seq[1:])
	}
	return true
}

type stringReader struct {
	text   string
	offset int
}

func (s *stringReader) next() byte {
	if s.offset >= len(s.text) {
		return 0
	}
	ch := s.text[s.offset]
	s.offset++
	return ch
}

func (s *stringReader) consumed() bool {
	return s.offset >= len(s.text)
}

type rule struct {
	ruleID    int
	letter    byte // = 0 if sequence node
	sequences [][]int
}

func (r *rule) String() string {
	if r.letter == 0 {
		return fmt.Sprintf("%d: seq %v", r.ruleID, r.sequences)
	} else {
		return fmt.Sprintf("%d: '%c'", r.ruleID, r.letter)
	}
}

func (r *rule) newSequence() {
	r.sequences = append(r.sequences, []int{})
}

func (r *rule) addToSequence(ruleID int) {
	i := len(r.sequences) - 1
	r.sequences[i] = append(r.sequences[i], ruleID)
}

func parseRules(text string) map[int]*rule {
	rules := map[int]*rule{}
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		rule := parseRule(line)
		rules[rule.ruleID] = rule
	}
	return rules
}

func parseRule(text string) *rule {
	t := newTokenizer(text)
	if t.next() != TokenNumber {
		t.error("exptected rule id")
	}
	rule := &rule{ruleID: t.value, sequences: make([][]int, 0)}
	token := t.next()
	if token == TokenLetter {
		rule.letter = byte(t.value)
	} else if token == TokenNumber {
		t.backup()
		parseSequence(t, rule)
	} else {
		t.error("expected letter or number")
	}
	return rule
}

func parseSequence(t *tokenizer, r *rule) {
	if t.next() != TokenNumber {
		t.error("expected number")
	}
	r.newSequence()
	r.addToSequence(t.value)
	for {
		tok := t.next()
		if tok == TokenEOF {
			break
		} else if tok == TokenNumber {
			// append to current sequence
			r.addToSequence(t.value)
		} else if tok == TokenPipe {
			// start a new sequence
			r.newSequence()
		} else {
			t.error("expected number or pipe")
		}
	}
}

type tokenizer struct {
	expr        string
	tokens      []string
	index       int
	token       tokenType
	tokenString string
	value       int // ascii for letter or int number
}

type tokenType int

const (
	TokenEOF    tokenType = 0
	TokenNumber tokenType = 1
	TokenLetter tokenType = 2
	TokenPipe   tokenType = 3
)

func newTokenizer(expr string) *tokenizer {
	tokens := strings.Split(expr, " ")
	return &tokenizer{
		expr:   expr,
		tokens: tokens,
	}
}

func (t *tokenizer) next() tokenType {
	if t.index >= len(t.tokens) {
		return TokenEOF
	}
	s := t.tokens[t.index]
	t.tokenString = s
	t.index++
	ch := rune(s[0])
	t.token = TokenEOF
	t.value = 0
	var err error
	if ch == '|' {
		t.token = TokenPipe
	} else if ch == '"' {
		t.token = TokenLetter
		t.value = int(s[1])
	} else if unicode.IsDigit(ch) {
		t.token = TokenNumber
		t.value, err = strconv.Atoi(strings.TrimSuffix(s, ":"))
		if err != nil {
			panic(err)
		}
	}
	return t.token
}

func (t *tokenizer) backup() {
	if t.token != TokenEOF {
		t.index--
	}
}

func (t *tokenizer) error(msg string, args ...interface{}) {
	s := fmt.Sprintf(msg, args...)
	panic(fmt.Sprintf("Error in '%s' at '%s': %s", t.expr, t.tokenString, s))
}
