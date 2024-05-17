package consts

import (
	"new-world-robot/pkg/convert"
)

const (
	HeaderUid         = "uid"
	HeaderAccessToken = "world-access-token"
)

var (
	EmptyRequestParam = convert.StringToBytes("{}")
)

const (
	AuthRequestParamTemplate = `
		{"account": "%s", "password": "%s", "auth_type": 0}
	`
	OnlyFishIdParamTemplate = `
		{"fish_id": "%s"}
	`
	OnlyPageParamTemplate = `
		{"page": %d, "page_size": %d}
	`
	OnlyProductIdParamTemplate = `
		{"product_id": "%s"}
	`
	MarketSellParamTemplate = `
		{"fish_id": "%s", "sell_duration": 0, "price": %d}
	`
)
