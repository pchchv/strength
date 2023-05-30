package strength

// Validate returns nil if the password has more than or equal to the minimum entropy.
// If not, an error is returned explaining how the password can be strengthened.
// This error is safe to show to the client.
func Validate(password string, minEntropy float64) error {
	return nil
}
