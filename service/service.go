package service

import (
	"context"
	"strings"
	"errors"
)

type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type StringServiceImpl struct {
}

func (StringServiceImpl) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringServiceImpl) Count(_ context.Context, s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string ")
