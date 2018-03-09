package gen_id

import (
"errors"
"math/rand"
	"time"
)

var str = string("123456789ABCDEFGHJK123456789MNPQRSTUVWXYZ123456789")

func New(length int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	if length < 1 {
		return "", errors.New("Min length must 1")
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	return string(b), nil

}
