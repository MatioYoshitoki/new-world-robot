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
		{"fish_id": %d}
	`
	OnlyPageParamTemplate = `
		{"page": %d, "page_size": %d}
	`
	OnlyProductIdParamTemplate = `
		{"product_id": %d}
	`
	MarketSellParamTemplate = `
		{"fish_id": %d, "sell_duration": 0, "price": %d}
	`
)
