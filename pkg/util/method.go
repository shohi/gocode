package util

import (
	"reflect"
	"strings"
)

// Signature represents a func or a method.
type Signature struct {
	receiver string
	name     string
	input    []string
	output   []string
}

// String format - `func (receiver) name(input...) (output...)`
func (s Signature) String() string {
	var buf strings.Builder
	buf.WriteString("func ")

	if s.receiver != "" {
		buf.WriteString("(")
		buf.WriteString(s.receiver)
		buf.WriteString(") ")
	}

	// Write input arguments
	buf.WriteString(s.name)
	buf.WriteString("(")
	buf.WriteString(strings.Join(s.input, ", "))
	buf.WriteString(")")

	// Write output arguments
	osz := len(s.output)
	switch {
	case osz == 1:
		buf.WriteString(" ")
		buf.WriteString(s.output[0])
	case osz > 1:
		buf.WriteString(" (")
		buf.WriteString(strings.Join(s.input, ", "))
		buf.WriteString(")")
	}

	return buf.String()
}

// ExtractPublicMethods from struct
// https://stackoverflow.com/questions/21397653/how-to-dump-methods-of-structs-in-golang
// https://stackoverflow.com/questions/54129042/how-to-get-a-functions-signature-as-string-in-go
// https://github.com/golang/go/issues/20995
func ExtractPublicMethods(v interface{}) []Signature {
	vType := reflect.TypeOf(v)

	// TODO: check other type
	if vType.Kind() == reflect.Ptr {
		if vType.Elem().Kind() == reflect.Ptr {
			panic("multiple pointers are not supported")
		}
	}

	var sigs []Signature
	for i := 0; i < vType.NumMethod(); i++ {
		m := vType.Method(i)
		sigs = append(sigs, methodToSignature(m))
	}

	return sigs
}

func methodToSignature(f interface{}) Signature {
	if !isMethod(f) {
		panic("<not a method>")
	}

	m := f.(reflect.Method)
	t := m.Func.Type()

	var input []string
	for i := 1; i < t.NumIn(); i++ {
		input = append(input, t.In(i).String())
	}

	var output []string
	for k := 0; k < t.NumOut(); k++ {
		output = append(output, t.Out(k).String())
	}

	return Signature{
		receiver: t.In(0).String(),
		name:     m.Name,
		input:    input,
		output:   output,
	}
}

func funcToSignature(f interface{}) Signature {
	if !isFunc(f) {
		panic("<not a function>")
	}

	t := reflect.TypeOf(f)

	var input []string
	for i := 0; i < t.NumIn(); i++ {
		input = append(input, t.In(i).String())
	}

	var output []string
	for k := 0; k < t.NumOut(); k++ {
		output = append(output, t.Out(k).String())
	}

	return Signature{
		receiver: "",
		name:     GetFuncBasename(f),
		input:    input,
		output:   output,
	}

}

func isFunc(f interface{}) bool {
	t := reflect.TypeOf(f)
	return (t.Kind() == reflect.Func)
}

// only Method value is supported,  pointer to Method not.
func isMethod(f interface{}) bool {
	_, ok := f.(reflect.Method)

	return ok
}
