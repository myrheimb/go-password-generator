package generator

import (
	cryptoRand "crypto/rand"
	"math/big"
	mathRand "math/rand"
)

// Generate a random index within the charset
func randomCharacter(charset string) byte {
	randomIndex, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(len(charset))))
	return charset[randomIndex.Int64()]
}

// Shuffle the password using the Fisher-Yates algorithm
func shufflePassword(password []byte) {
	mathRand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})
}

func GeneratePassword(passwordLength int, lower bool, upper bool, numbers bool, special bool) string {
	// Define the character sets for each type
	const lowerChars = "abcdefghijklmnopqrstuvwxyz"
	const upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const numberChars = "0123456789"
	const specialChars = "!@#$%^&*-_+<>?"

	var charset string
	var password []byte

	// Add at least one character from each selected type
	if lower {
		password = append(password, randomCharacter(lowerChars))
		charset += lowerChars
	}
	if upper {
		password = append(password, randomCharacter(upperChars))
		charset += upperChars
	}
	if numbers {
		password = append(password, randomCharacter(numberChars))
		charset += numberChars
	}
	if special {
		password = append(password, randomCharacter(specialChars))

		// Limit special characters to 10% of the password length for ease of typing
		remainingLength := passwordLength - len(password)
		numSpecial := 0
		numSpecial = remainingLength / 10
		for i := 0; i < numSpecial; i++ {
			password = append(password, randomCharacter(specialChars))
		}
	}

	// Fill the remaining length of the password with characters from the charset
	for len(password) < passwordLength {
		password = append(password, randomCharacter(charset))
	}

	// Shuffle the password to ensure randomness
	shufflePassword(password)

	return string(password)
}
