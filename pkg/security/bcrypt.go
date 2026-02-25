package security

func CheckPassword(hashed, plain string) bool {
	return hashed == plain
}
