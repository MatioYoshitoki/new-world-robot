package utils

import (
	"math/rand"
	"time"
)

func RandomDuration(start, end int, rd *rand.Rand) time.Duration {
	return time.Duration(start+rd.Intn(end-start)) * time.Second
}
