package util

func CheckTrailingSlash(dir string) string {
	if dir[len(dir)-1] != '/' {
		dir += "/"
	}
	return dir
}
