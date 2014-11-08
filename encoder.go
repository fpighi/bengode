package encoder

import (
	"bytes"
	"container/list"
	"fmt"
	"reflect"
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
	encoded, _ := getEncodedValue(v)

	e.b.WriteString(encoded)
	return nil
}

func (e *encodeProcess) getString() string {
	return e.b.String()
}

func getEncodedValue(v interface{}) (string, *TypeNotSupportedError) {
	switch x := v.(type) {
	case int:
		return encodeInt64(int64(x)), nil
	case int8:
		return encodeInt64(int64(x)), nil
	case int16:
		return encodeInt64(int64(x)), nil
	case int32:
		return encodeInt64(int64(x)), nil
	case int64:
		return encodeInt64(int64(x)), nil
	case string:
		return encodeString(string(x)), nil
	case *list.List:
		return encodeList(x), nil
	default:
		fmt.Printf("The type is %v", reflect.TypeOf(v))
		return "", &TypeNotSupportedError{msg: "Type not supported", Type: reflect.TypeOf(v).String()}
	}
}

func encodeList(l *list.List) string {
	var buffer bytes.Buffer
	buffer.WriteString("l")
	for e := l.Front(); e != nil; e = e.Next() {
		encoded, _ := getEncodedValue(e.Value)
		buffer.WriteString(encoded)
	}
	buffer.WriteString("e")
	return buffer.String()
}

func encodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return fmt.Sprint(len(s), ":", s)
}

func encodeInt64(n int64) string {
	return fmt.Sprintf("i%de", n)
}

type TypeNotSupportedError struct {
	msg  string
	Type string
}

func (e *TypeNotSupportedError) Error() string { return e.msg }
