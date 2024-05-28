package generators

import (
	"math/rand"
	"strings"
)

const chars = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(size int) string {
    strBuilder := &strings.Builder{}
    for range size {
        idx := rand.Intn(len(chars))
        strBuilder.WriteByte(chars[idx])
    }

    return strBuilder.String()
}
