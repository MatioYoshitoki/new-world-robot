package utils

import (
	"math/rand"
	"time"
)

func RandomDuration(start, end int, duration time.Duration, rd *rand.Rand) time.Duration {
	return time.Duration(start+rd.Intn(end-start)) * duration
}
