package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder

	var prev rune

	type states int

	const (
		begin states = iota
		escape
		symbol
	)

	state := begin

	for _, cur := range s {
		switch {
		case '0' <= cur && cur <= '9':
			switch state {
			case begin:
				return "", ErrInvalidString
			case escape:
				prev = cur
				state = symbol
			case symbol:
				for i := rune(0); i < cur-'0'; i++ {
					sb.WriteRune(prev)
				}
				state = begin
			}
		case cur == '\\':
			switch state {
			case begin:
				state = escape
			case escape:
				prev = cur
				state = symbol
			case symbol:
				sb.WriteRune(prev)
				state = escape
			}
		default:
			switch state {
			case begin:
				prev = cur
				state = symbol
			case escape:
				return "", ErrInvalidString
			case symbol:
				sb.WriteRune(prev)
				prev = cur
			}
		}
	}

	switch state {
	case begin:
	case escape:
		return "", ErrInvalidString
	case symbol:
		sb.WriteRune(prev)
	}

	return sb.String(), nil
}
