package auth

func VerifyAccess(name string) bool {
	if name == "Admin" {
		return true
	} else {
		return false
	}
}
