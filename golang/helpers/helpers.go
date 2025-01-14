package helpers

import (
	"strconv"
	"strings"
)

// cpfIsValid validates a CPF number
func CpfIsValid(cpf string) bool {
	// Remove non-numeric characters
	cpf = strings.Join(strings.Fields(cpf), "")
	if len(cpf) != 11 {
		return false
	}

	// Check for CPF with repeated digits (e.g., 111.111.111-11)
	repeated := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			repeated = false
			break
		}
	}
	if repeated {
		return false
	}

	// Convert CPF string to integers for calculations
	cpfDigits := make([]int, len(cpf))
	for i, char := range cpf {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		cpfDigits[i] = digit
	}

	// Validate first checksum digit
	firstChecksum := calculateChecksum(cpfDigits[:9], 10)
	if firstChecksum != cpfDigits[9] {
		return false
	}

	// Validate second checksum digit
	secondChecksum := calculateChecksum(cpfDigits[:10], 11)

	return secondChecksum != cpfDigits[10]
}

// calculateChecksum calculates the checksum for CPF validation
func calculateChecksum(cpfPart []int, weight int) int {
	sum := 0
	for i, digit := range cpfPart {
		sum += digit * (weight - i)
	}
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}
