package nlenc

// Bytes returns a null-terminated byte slice with the contents of s.
func Bytes(s string) []byte {
	return append([]byte(s), 0x00)
}

// String returns a string with the contents of b from a null-terminated
// byte slice.
func String(b []byte) string {
	// If the string has more than one NULL terminator byte, we want to remove
	// all of them before returning the string to the caller.
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0x00 {
			return string(b[:i+1])
		}
	}
	return ""
}
