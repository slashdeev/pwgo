package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
)

const charList = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*()-_=+[]{}<>?"

func generatePassword(length int) string {
	var stringBuilder strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charList))
		stringBuilder.WriteByte(charList[randomIndex])
	}

	return stringBuilder.String()
}

func main() {
	lengthPtr := flag.Int("len", 12, "The length of characters.")
	flag.Parse()
	length := *lengthPtr

	if length <= 0 {
		fmt.Println("Invalid length: Must be a positive number.")
		return
	}

	password := generatePassword(length)
	fmt.Println(password)
}
