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

func (f *Future) Wait() {
	<-f.wait
}

func (f *Future) Result() (interface{}, error) {
	<-f.wait

	return f.result, f.err
}

func (f *Future) run() {
	defer close(f.wait)

	numParams := f.functionType.NumIn()
	values := make([]reflect.Value, numParams)
	for i := 0; i < numParams; i++ {
		values[i] = reflect.ValueOf(f.args[i])
	}

	ret := f.functionValue.Call(values)

	if len(ret) == 0 {
		return
	}

	f.result = ret[0].Interface()

	if f.functionType.NumOut() > 1 && !ret[1].IsNil() {
		// nolint
		f.err = ret[1].Interface().(error)
	}
}
