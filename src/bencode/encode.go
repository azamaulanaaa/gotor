package bencode

import (
	"errors"
	"fmt"
	"sort"
)

type encoder func(v interface{}) (string, error)

var (
    ErrorNotSupported = errors.New("type not supported")
)

func Encode(v interface{}) (string, error) {
    for _, theEcoder := range []encoder{
        encodeString,
        encodeInteger,
        encodeList,
        encodeDictionary,
    } {
        out, err := theEcoder(v)
        if err == nil {
            return out, nil
        }
    }
    return "", ErrorNotSupported
}

func encodeString(v interface{}) (string, error) {
    var out String

    switch value := v.(type) {
    case String:
        out = value
    case []byte:
        out = String(value)
    case string:
        out = String(value)
    default:
        return "", ErrorNotSupported
    }

    return fmt.Sprintf("%d:%s", len(out), out), nil
}

func encodeInteger(v interface{}) (string, error) {
    var out Integer

    switch value := v.(type) {
    case Integer:
        out = value
    case int:
        out = Integer(value)
    case int8:
        out = Integer(value)
    case int16:
        out = Integer(value)
    case int32:
        out = Integer(value)
    case int64:
        out = Integer(value)
    case uint:
        out = Integer(value)
    case uint8:
        out = Integer(value)
    case uint16:
        out = Integer(value)
    case uint32:
        out = Integer(value)
    case uint64:
        out = Integer(value)
    default:
        return "", ErrorNotSupported
    }

    return fmt.Sprintf("i%de", out), nil
}

func encodeList(v interface{}) (string, error) {
    var out List

    switch value := v.(type) {
    case List:
        out = value
    case []interface{}:
        out = List(value)
    default:
        return "", ErrorNotSupported
    }

    var outStr string
    for _, v := range out {
        value, err := Encode(v)
        if err != nil {
            return "", err
        }

        outStr = fmt.Sprintf("%s%s", out, value)
    }

    outStr = fmt.Sprintf("l%se", outStr)
    return outStr, nil
}

func encodeDictionary(v interface{}) (string, error) {
    var out Dictionary

    switch value := v.(type) {
    case Dictionary:
        out = value
    case map[String]interface{}:
        out = Dictionary(value)
    case map[string]interface{}:
        for k, v := range value{
            out[String(k)] = v
        }
    default:
        return "", ErrorNotSupported
    }

    sortedKey := make([]string, 0, len(out))
    for k := range out {
        sortedKey = append(sortedKey, string(k))
    }
    sort.Strings(sortedKey)

    var outStr string
    for _, k := range sortedKey {
        key, err := Encode(k)
        if err != nil {
            return "", err
        }

        value, err := Encode(out[String(k)])
        if err != nil {
            return "", err
        }

        outStr = fmt.Sprintf("%s%s%s", outStr, key, value)
    }
    outStr = fmt.Sprintf("d%se", outStr)

    return outStr, nil
}
