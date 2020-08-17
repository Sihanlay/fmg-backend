package mailUtils

import "regexp"

func CheckMailFormat(email string) bool {

	mailCompile := regexp.MustCompile("^(.*)@(.*)\\.(.*)$")

	r := mailCompile.FindSubmatch([]byte(email))
	if len(r) != 4 {
		return false
	}
	return true
}
