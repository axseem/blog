package test

import (
	"strings"
	"testing"
)

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	return &Assert{t: t}
}

func (a *Assert) True(condition bool, msg ...string) {
	if !condition {
		a.t.Error(strings.Join(msg, ": "))
	}
}

func (a *Assert) Nil(err error) {
	if err != nil {
		a.t.Error(err)
	}
}

func (a *Assert) NilDefer(fn func() error) {
	if err := fn(); err != nil {
		a.t.Error(err)
	}
}
