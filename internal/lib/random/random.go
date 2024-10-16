package random

import (
	"math/rand"
)

func NewRandomString(aliasLength int) string {
	symb := []rune(
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
			"abcdefghijklmnopqrstuvwxyz" +
			"0123456789")
	str := make([]rune, aliasLength)

	for i := range aliasLength {
		str[i] = symb[rand.Intn(len(symb))]
	}
	return string(str)
}
