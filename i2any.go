package i2any

import (
    "errors"
    "strconv"
)

type Result struct {
    Data interface{}
}

var undefined = errors.New("i2any: 400, undefined")
var failedFormat = errors.New("i2any: 400, failed format")

func (r *Result) Get(k string) *Result {
    if v, ok := r.Data.(map[string]interface{}); ok {
        _v1 := v[k]
        return &Result{Data: _v1}
    }
    panic(undefined)
}

func (r *Result) String() string {
    if v, ok := r.Data.(string); ok {
        return v
    }
    return ""
}

func (r *Result) FString() string {
    switch r.Data.(type) {
    case string:
        return r.Data.(string)
    case float32:
        return strconv.FormatFloat(float64(r.Data.(float32)), 'f', 2, 64)
    case float64:
        return strconv.FormatFloat(r.Data.(float64), 'f', 2, 64)
    case int:
        return strconv.Itoa(r.Data.(int))
    case int8:
        return strconv.FormatInt(int64(r.Data.(int8)), 10)
    case int16:
        return strconv.FormatInt(int64(r.Data.(int16)), 10)
    case int32:
        return strconv.FormatInt(int64(r.Data.(int32)), 10)
    case int64:
        return strconv.FormatInt(r.Data.(int64), 10)
    case uint:
        return strconv.FormatUint(uint64(r.Data.(uint)), 10)
    case uint8:
        return strconv.FormatUint(uint64(r.Data.(uint8)), 10)
    case uint16:
        return strconv.FormatUint(uint64(r.Data.(uint16)), 10)
    case uint32:
        return strconv.FormatUint(uint64(r.Data.(uint32)), 10)
    case uint64:
        return strconv.FormatUint(r.Data.(uint64), 10)
    default:
        panic(failedFormat)
    }
}

func (r *Result) Int() int64 {
    if v, ok := r.Data.(int64); ok {
        return v
    }
    return 0
}

func (r *Result) FInt() int64 {
    switch r.Data.(type) {
    case string:
        v, err := strconv.ParseInt(r.Data.(string), 10, 64)
        if err != nil {
            panic(err)
        }
        return v
    case float32:
        return int64(r.Data.(float32))
    case float64:
        return int64(r.Data.(float64))
    case int:
        return int64(r.Data.(int))
    case int8:
        return int64(r.Data.(int8))
    case int16:
        return int64(r.Data.(int16))
    case int32:
        return int64(r.Data.(int32))
    case int64:
        return r.Data.(int64)
    case uint:
        return int64(r.Data.(uint))
    case uint8:
        return int64(r.Data.(uint8))
    case uint16:
        return int64(r.Data.(uint16))
    case uint32:
        return int64(r.Data.(uint32))
    case uint64:
        return int64(r.Data.(uint64))
    default:
        panic(failedFormat)
    }
}
