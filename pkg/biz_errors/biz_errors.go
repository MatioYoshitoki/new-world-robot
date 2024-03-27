package biz_errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	AuthError = errors.New(401, "", "")
)
