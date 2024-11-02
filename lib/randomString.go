package lib

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	time := time.Now().UnixNano()

	return fmt.Sprintf("data : %s\ntime : %d\n", string(b), time)
}
