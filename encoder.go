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
	encoded, err := getEncodedValue(v)

	e.b.WriteString(encoded)
	return err
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
		return encodeList(x)
	case []interface{}:
		return encodeSlice(x)
	case map[string]interface{}:
		return encodeDictionary(x)
	default:
		fmt.Printf("The type is %v", reflect.TypeOf(v))
		return "", &TypeNotSupportedError{msg: "Type not supported", Type: reflect.TypeOf(v).String()}
	}
}

func encodeDictionary(d map[string]interface{}) (string, *TypeNotSupportedError) {
	var buffer bytes.Buffer
	buffer.WriteString("d")
	for k, v := range d {
		encodedKey := encodeString(k)
		encodedValue, err := getEncodedValue(v)
		if err != nil {
			return "", err
		}
		buffer.WriteString(encodedKey)
		buffer.WriteString(encodedValue)
	}
	buffer.WriteString("e")

	return buffer.String(), nil
}

func encodeSlice(s []interface{}) (string, *TypeNotSupportedError) {
	var buffer bytes.Buffer
	buffer.WriteString("l")
	for _, e := range s {
		encoded, err := getEncodedValue(e)
		if err != nil {
			return "", err
		}
		buffer.WriteString(encoded)
	}
	buffer.WriteString("e")
	return buffer.String(), nil
}

func encodeList(l *list.List) (string, *TypeNotSupportedError) {
	var buffer bytes.Buffer
	buffer.WriteString("l")
	for e := l.Front(); e != nil; e = e.Next() {
		encoded, err := getEncodedValue(e.Value)
		if err != nil {
			return "", err
		}
		buffer.WriteString(encoded)
	}
	buffer.WriteString("e")
	return buffer.String(), nil
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
