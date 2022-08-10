/*
@Time : 2022/8/5 下午8:15
@Author : MuYiMing
@File : lexer
@Software: GoLand
*/
package sqlparser

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type lexer struct {
	input string

	start    int
	position int
	width    int

	token chan lexToken
}

type lexFunc func(token *lexer) lexFunc

func newLexer(input string) *lexer {
	l := &lexer{
		input: input,
		token: make(chan lexToken),
	}
	return l
}

func (l *lexer) run() {
	for state := lexStatement; state != nil; {
		state(l)
	}
}

func (l *lexer) nextToken() lexToken {
	return <-l.token
}

func (l *lexer) produce(t tokenType) {
	l.token <- lexToken{
		tokenType: t,
		vlaue:     l.input[l.start:l.position],
	}
	l.start = l.position
}

func (l *lexer) next() rune {
	if l.position > len(l.input) {
		l.width = 0
		return end
	}

	r, w := utf8.DecodeLastRuneInString(l.input[l.position:])

	l.width = w
	l.position += w

	return r
}

func (l *lexer) revert() {
	l.position -= l.width
}

func (l *lexer) peek() rune {
	r := l.next()
	l.revert()

	return r
}

func (l *lexer) drain() {
	for range l.token {
	}
}

func lexStatement(l *lexer) lexFunc {
	r := l.next()

	switch true {
	case unicode.IsDigit(r):
		return lexInteger
	case isAlpthaNumeric(r):
		return lexIdentifier

	case unicode.IsSpace(r):
		l.produce(tokenSpace)
		return lexStatement
	case r == '(':
		l.produce(tokenLeftParenthesis)
		return lexStatement
	case r == ')':
		l.produce(tokenRightParenthesis)
		return lexStatement
	case r == '"':
		return lexString
	case r == ',':
		l.produce(tokenDelimeter)
		return lexStatement
	case r == '=':
		switch l.peek() {
		case '=':
			l.next()
			l.produce(tokenEquals)
		default:
			l.produce(tokenAssign)
		}
		return lexStatement

	case r == end:
		l.produce(tokenEnd)
		return nil
	}

	return l.errorf("unexpected rune")
}

func lexInteger(l *lexer) lexFunc {
	r := l.peek()
	if unicode.IsDigit(r) {
		l.next()
		return lexInteger
	}
	l.produce(tokenTypeInteger)
	return lexStatement
}

func lexIdentifier(l *lexer) lexFunc {
	r := l.next()
	if isAlpthaNumeric(r) {
		return lexIdentifier
	}
	l.revert()

	word := l.input[l.start:l.position]
	if t, ok := keyWords[strings.ToUpper(word)]; ok {
		l.produce(t)
	} else {
		l.produce(tokenIdentifier)
	}

	return lexStatement
}

func isAlpthaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

func lexString(l *lexer) lexFunc {
	r := l.next()

	switch r {
	case '"':
		l.produce(tokenString)

		return lexStatement

	case end:
		return l.errorf("expected \"")
	}

	return lexString
}

func (l *lexer) errorf(format string, arg ...interface{}) lexFunc {
	l.token <- lexToken{
		tokenType: tokenError,
		vlaue:     fmt.Sprintf(format, arg),
	}
	return nil
}
