package net_utils

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type CommonResult[T any] struct {
	Code    int64
	Message string
	Data    T
}

func RequestToStruct[T any](ctx context.Context, hc *http.Client, request *http.Request) (*CommonResult[T], error) {
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
	return &rst, nil
}
