package endPointMethods

import "../../models/security/origin"

func CheckToken(token, addr string) (ok bool) {
	var count int
	origin.Model.Where(origin.Allow{
		ID:       addr,
		Password: token,
	}).Count(&count)

	if count > 0 {
		ok = true
	}
	return
}
