package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowercase = "abcdefghijklmnopqrstuvwxyz"
const numbers = "0123456789"
const symbols = "!@#$%^&*()-_=+"

func generatePassword(length int, pool string) (string, error) {
	var password strings.Builder
	max := big.NewInt(int64(len(pool)))

	for i := 0; i < length; i++ {
		number, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		password.WriteByte(pool[number.Int64()])
	}

	return password.String(), nil
}

func main() {
	lengthPtr := flag.Int("len", 12, "The length of characters.")
	quantityPtr := flag.Int("quantity", 5, "The amount of passwords.")
	symbolsPtr := flag.Bool("symbols", false, "Include special characters.")

	flag.Parse()

	pool := uppercase + lowercase + numbers
	if *symbolsPtr {
		pool += symbols
	}

	if *lengthPtr <= 0 || *quantityPtr <= 0 {
		fmt.Fprintln(os.Stderr, "Length / Quantity must be positive!")
		os.Exit(1)
	}

	for i := 0; i < *quantityPtr; i++ {
		password, err := generatePassword(*lengthPtr, pool)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Something went wrong while trying to generate the password.")
			fmt.Fprintln(os.Stderr, err.Error())

			os.Exit(1)
		}

		fmt.Println(password)
	}
}
