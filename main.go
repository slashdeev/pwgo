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
	lengthPtr := flag.Int("length", 12, "The length of characters.")
	quantityPtr := flag.Int("quantity", 1, "The amount of passwords.")
	flag.Parse()

	if *lengthPtr <= 0 {
		fmt.Println("Invalid length: Must be a positive number.")
		return
	}

	if *quantityPtr <= 0 {
		fmt.Println("Invalid quantity: Must be a positive number.")
		return
	}

	for i := 0; i < *quantityPtr; i++ {
		fmt.Println(generatePassword(*lengthPtr))
	}
}
