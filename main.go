// `pwgo` is a fast, modern and optimized CLI Tool.
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

/*
	[DISCLAIMER]
	these variables should not be modified, they are safe-
	to use in some applications.
*/

const (
	minimumLength = 5
	maximumLength = 128

	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "0123456789"
	symbols   = "!@#$%&*"
)

// `generatePassword` generates a secure password with the given length.
func generatePassword(passwordLength int, characterPool string) (string, error) {
	if minimumLength > passwordLength {
		return "", fmt.Errorf("password length is below minimum. (%d < %d)", passwordLength, minimumLength)
	}

	if maximumLength < passwordLength {
		return "", fmt.Errorf("password length is above maximum. (%d > %d)", passwordLength, maximumLength)
	}

	if len(characterPool) == 0 {
		return "", fmt.Errorf("character pool cannot be empty.")
	}

	passwordBuilder := strings.Builder{}
	passwordBuilder.Grow(passwordLength)

	for i := 0; i < passwordLength; i++ {
		max := len(characterPool)
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil {
			return "", fmt.Errorf("failed to generate the random number: %w", err)
		}

		err = passwordBuilder.WriteByte(characterPool[randomIndex.Int64()])
		if err != nil {
			return "", fmt.Errorf("failed to write byte: %w", err)
		}
	}

	return passwordBuilder.String(), nil
}

func main() {
	lengthPtr := flag.Int("length", 12, "The length of characters.")
	quantityPtr := flag.Int("quantity", 1, "The amount of passwords.")
	symbolsPtr := flag.Bool("symbols", false, "Include special characters.")

	flag.Parse()

	charactersPool := (uppercase + lowercase + numbers)

	if *symbolsPtr {
		charactersPool += symbols
	}

	if *lengthPtr < minimumLength {
		fmt.Fprintf(os.Stderr, "password length is below minimum. (%d < %d)", *lengthPtr, minimumLength)
		os.Exit(1)
	}

	if *quantityPtr <= 0 {
		fmt.Fprintf(os.Stderr, "quantity is below minimum. (%d <= %d)", *quantityPtr, 0)
		os.Exit(1)
	}

	for i := 0; i < *quantityPtr; i++ {
		password, err := generatePassword(*lengthPtr, charactersPool)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[%d] something went wrong, continuing...\n", (i + 1))
			continue
		}

		fmt.Printf("[%d] %s\n", (i + 1), password)
	}
}
