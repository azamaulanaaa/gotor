package lib

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

var (
    ErrNotStruct = errors.New("interface is not a struct")
)

type toString interface {
    String() string
}

func QueryString(v interface{}, key string) (string, error) {
    var queryString string
    refType := reflect.TypeOf(v)
    if refType.Kind() != reflect.Struct {
        return "", ErrNotStruct
    }

    numFilds := refType.NumField()

    for i:=0; i< numFilds; i++ {
        fieldType := refType.Field(i)
        fieldValue := reflect.ValueOf(v).Field(i)

        if fieldValue.IsZero() {
            continue
        }

        if queryString != "" {
            queryString = fmt.Sprintf("%s&", queryString)
        }

        var valueString string
        switch value := fieldValue.Interface().(type) {
        case toString:
            valueString = value.String()
        case int:
            valueString = strconv.FormatInt(int64(value), 10)
        case int8:
            valueString = strconv.FormatInt(int64(value), 10)
        case int16:
            valueString = strconv.FormatInt(int64(value), 10)
        case int32:
            valueString = strconv.FormatInt(int64(value), 10)
        case int64:
            valueString = strconv.FormatInt(int64(value), 10)
        case uint:
            valueString = strconv.FormatUint(uint64(value), 10)
        case uint8:
            valueString = strconv.FormatUint(uint64(value), 10)
        case uint16:
            valueString = strconv.FormatUint(uint64(value), 10)
        case uint32:
            valueString = strconv.FormatUint(uint64(value), 10)
        case uint64:
            valueString = strconv.FormatUint(uint64(value), 10)
        case bool:
            valueString = strconv.FormatBool(value)
        default:
            valueString = fieldValue.String()
        }

        valueString = url.QueryEscape(valueString)

        queryString = fmt.Sprintf(
            "%s%s=%s",
            queryString,
            fieldType.Tag.Get(key),
            valueString,
        )
    }

    return queryString, nil
}
