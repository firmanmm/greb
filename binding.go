package greb

import (
	"net/http"
	"strconv"
)

type BindDataType uint

const (
	DATA_BIND_TYPE_INT    BindDataType = 1
	DATA_BIND_TYPE_FLOAT  BindDataType = 2
	DATA_BIND_TYPE_STRING BindDataType = 3
	DATA_BIND_TYPE_BOOL   BindDataType = 4
)

type BindType uint

const (
	BIND_TYPE_FORM   BindType = 1
	BIND_TYPE_QUERY  BindType = 2
	BIND_TYPE_JSON   BindType = 3
	BIND_TYPE_HEADER BindType = 4
	BIND_TYPE_COOKIE BindType = 5
)

type IBindable interface {
	BindRequest(req *http.Request) error
}

func _ResolveData(req *http.Request, key string, bindType BindType) string {
	switch bindType {
	case BIND_TYPE_FORM:
		return req.PostFormValue(key)
	case BIND_TYPE_QUERY:
		return req.URL.Query().Get(key)
	case BIND_TYPE_HEADER:
		return req.Header.Get(key)
	case BIND_TYPE_COOKIE:
		cookie, err := req.Cookie(key)
		if err != nil {
			return ""
		}
		return cookie.Value
	default:
		return ""
	}
}

func BindInt(req *http.Request, key string, bindType BindType) (int, error) {
	data := _ResolveData(req, key, bindType)
	return strconv.Atoi(data)
}

func BindFloat(req *http.Request, key string, bindType BindType) (float64, error) {
	data := _ResolveData(req, key, bindType)
	return strconv.ParseFloat(data, 64)
}

func BindString(req *http.Request, key string, bindType BindType) (string, error) {
	data := _ResolveData(req, key, bindType)
	return data, nil
}

func BindBool(req *http.Request, key string, bindType BindType) (bool, error) {
	data := _ResolveData(req, key, bindType)
	return strconv.ParseBool(data)
}
