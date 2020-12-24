package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func main() {
	var str = RandString(16)
	fmt.Printf("str: %+v", str)
	fmt.Printf("\n\n")
}
