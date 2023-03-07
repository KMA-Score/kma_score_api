package utils

// Contains https://play.golang.org/p/Qg_uv_inCek
// Contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
