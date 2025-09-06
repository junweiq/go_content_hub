package util

func GetUserSidKey(username string) string {
	return "usid:" + username
}

func GetUserSidCreateAtKey(sid string) string {
	return "usid_ca:" + sid
}
