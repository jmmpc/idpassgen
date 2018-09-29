package idpassgen

import (
	"math/rand"
)

// chars[:23]   - special characters
// chars[23:33] - digits
// chars[33:59] - uppercase letters
// chars[59:]   - lowercase letters
// chars[23:39] - hex
// chars[23:]   - characters acceptable for an ID
const chars = "!?.,{}[]<>()^*-+=#@:;~_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// NewID returns a pseudo random string of a given length, which may
// contain uppercase letters, lowercase letters and digits.
func NewID(length int, rnd *rand.Rand) string {
	if length < 1 {
		length = 1
	}

	id := make([]byte, length)
	for i := range id {
		id[i] = chars[23+rnd.Int63()%62] // character in chars[23:84]
	}

	return string(id)
}

// NewPassword returns a pseudo random string of a given length, which may
// contain uppercase letters, lowercase letters, digits and special characters.
// Min length == 4.
func NewPassword(length int, rnd *rand.Rand) string {
	if length < 4 {
		length = 4
	}

	pass := make([]byte, length)

	// make sure the password contains a digit, uppercase letter, lowercase letter and special character
	pass[0] = chars[rnd.Int63()%23]    // special character in chars[:23]
	pass[1] = chars[23+rnd.Int63()%10] // digit in chars[23:33]
	pass[2] = chars[33+rnd.Int63()%26] // uppercase letter in chars[33:59]
	pass[3] = chars[59+rnd.Int63()%26] // lowercase letter in chars[59:]

	for i := 4; i < length; i++ {
		pass[i] = chars[rnd.Int63()%int64(len(chars))]
	}

	rnd.Shuffle(length, func(i, j int) {
		pass[i], pass[j] = pass[j], pass[i]
	})

	return string(pass)
}

// NewHex returns a pseudo random string of a given length, which may
// contain digits in range 0-9 and uppercase letters in range A-F.
func NewHex(length int, rnd *rand.Rand) string {
	if length < 1 {
		length = 1
	}

	id := make([]byte, length)
	for i := range id {
		id[i] = chars[23+rnd.Int63()%16] // character in chars[23:39]
	}

	return string(id)
}
