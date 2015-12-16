package utils

// Mod is the function which always return the positive modulus
func Mod(a int, b int) int {
	rem := a % b
	if rem < 0 {
		rem += b
	}
	return rem
}
