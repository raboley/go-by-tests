package iteration

// Repeat takes a character and returns the character multiplies by the count parameter as a string.
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}