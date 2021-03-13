package util

func AddMapBToA(a map[string]int, b map[string]int) map[string]int {
	for k, _ := range b {
		if _, present := a[k]; present {
			a[k] = a[k] + b[k]
		} else {
			a[k] = b[k]
		}
	}
	return a
}
