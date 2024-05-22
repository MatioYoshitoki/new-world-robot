package net_utils

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"io"
	"net/http"
)

type CommonResult[T any] struct {
	Code    int64
	Message string
	Data    T
}

func RequestToStruct[T any](hc *http.Client, request *http.Request) (*CommonResult[T], error) {
	response, err := hc.Do(request)
	if err != nil {
		return nil, err
	}
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	rst := CommonResult[T]{}
	if err := json.Unmarshal(respBody, &rst); err != nil {
		return nil, err
	}
	if rst.Code == 503 {
		//panic(err)
		return nil, errors.New(-503, "网络错误", "")
	}
	return &rst, nil
}
