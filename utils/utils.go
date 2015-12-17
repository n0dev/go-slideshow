package utils

// Mod is the function which always return the positive modulus
func Mod(a int, b int) int {
	rem := a % b
	if rem < 0 {
		rem += b
	}
	return rem
}

// StringInSlice returns true if the string a is in slice, else false
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
