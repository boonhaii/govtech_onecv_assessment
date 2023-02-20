package Util

func Intersection(s1, s2 []string) (res []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}

	for _, e := range s2 {
		if hash[e] {
			res = append(res, e)
		}
	}

	return res
}

func Difference(s1, s2 []string) (res []string) {
	hash := make(map[string]bool)
	for _, e := range s2 {
		hash[e] = true
	}

	for _, e := range s1 {
		_, ok := hash[e]
		if !ok {
			res = append(res, e)
		}
	}
	
	return res
}