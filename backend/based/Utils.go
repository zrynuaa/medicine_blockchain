package based

import "strings"

//need policy == hid* OR(cname* AND region1)
func match(attr []string, policy string) bool{
	policy = strings.Replace(policy, " ", "", -1)
	s := strings.Split(policy, "OR")
	if isexist(attr, s[0]){
		return true
	}
	s2 := strings.Replace(s[1], "(", "", -1)
	s2 = strings.Replace(s2, ")", "", -1)
	s3 :=  strings.Split(s2, "AND")
	if isexist(attr, s3[0]) && isexist(attr, s3[1]) {
		return true
	}
	return false
}

func isexist(attr []string, a string) bool{
	for _, i := range attr {
		if i == a {
			return true
		}
	}
	return false
}
