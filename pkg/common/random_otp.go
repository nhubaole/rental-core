package common

import (
	"time"
	"math/rand"

)

func GenerateDigitOTP() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + rng.Intn(900000)
	return otp
}	