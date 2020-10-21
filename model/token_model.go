package model

type Token struct{}

func (t *Token) NewToken() *Token {
	return &Token{}
}
