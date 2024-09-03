package main

import (
	"flag"
	"fmt"
	"go-password-generator/generator"
	"os"
	"strconv"
	"strings"
)

// Ensure that the character types are valid
func validateCharacterTypes(types string) error {
	for _, char := range types {
		if !strings.ContainsRune("luns", char) {
			return fmt.Errorf("invalid character type: %c. Only 'l', 'u', 'n', and 's' are allowed", char)
		}
	}
	return nil
}

// Ensure that the password length is a valid integer
func validateLength(lengthStr string) (int, error) {
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return 0, fmt.Errorf("invalid length: %s. Please provide a numeric value", lengthStr)
	}
	if length <= 0 {
		return 0, fmt.Errorf("password length must be greater than 0")
	}
	return length, nil
}

// Determine which character types to include
func parseCharacterTypes(types string) (lower, upper, numbers, special bool) {
	lower = strings.Contains(types, "l")
	upper = strings.Contains(types, "u")
	numbers = strings.Contains(types, "n")
	special = strings.Contains(types, "s")
	return
}

func main() {
	// Define the command line flags
	passwordLength := flag.String("length", "12", "Length of the password")
	charTypes := flag.String("types", "luns", "Character types to include (l=lower, u=upper, n=numbers, s=special)")

	// Parse the flags from the command line
	flag.Parse()

	// Validate the character types
	if err := validateCharacterTypes(*charTypes); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Validate the length and convert it to an integer
	length, err := validateLength(*passwordLength)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Determine which character types to include
	lower, upper, numbers, special := parseCharacterTypes(*charTypes)

	// Generate the password using the generator package
	password := generator.GeneratePassword(length, lower, upper, numbers, special)

	// Print the generated password
	fmt.Println("Generated Password:", password)
}
