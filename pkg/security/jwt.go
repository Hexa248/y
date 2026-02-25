package security

func GenerateToken(email string) string {
	return "dev-token-" + email
}
