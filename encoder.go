package encoder

import (
	"bytes"
	"fmt"
)

func Marshal(v interface{}) (string, error) {
	e := &encodeProcess{}
	err := e.marshal(v)
	return e.getString(), err
}

type encodeProcess struct {
	b bytes.Buffer
}

func (e *encodeProcess) marshal(v interface{}) (err error) {
	s, isString := v.(string)
	if isString {
		encoded := stringEncodeFunc()(s)
		e.b.WriteString(encoded)
	}
	return nil
}

func (e *encodeProcess) getString() string {
	return e.b.String()
}

func stringEncodeFunc() func(string) string {
	return func(s string) string {
		if len(s) == 0 {
			return ""
		}
		return fmt.Sprint(len(s), ":", s)
	}
}
