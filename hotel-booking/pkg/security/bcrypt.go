package security

func HashPassword(password string) string        { return password }
func ComparePassword(hash, password string) bool { return hash == password }
