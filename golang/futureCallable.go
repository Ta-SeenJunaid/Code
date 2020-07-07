package main

import (
	"reflect"
	"strings"
)

type lines struct {
	out string
}

func (l *lines) str(lines []string) {
	l.out = strings.Join(lines, "\n")
}

func (l *lines) delete() {
	l.out = ""
}

type Future struct {
	functionType  reflect.Type
	functionValue reflect.Value
	args          []interface{}
	wait          chan struct{}
	result        interface{}
	err           error
}
