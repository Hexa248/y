package security

import "fmt"

func GenerateToken(userID int, role string) string {
	return fmt.Sprintf("token-%d-%s", userID, role)
}
