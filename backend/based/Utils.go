package based

import "strings"
import (
	"github.com/Doresimon/SM-Collection/SM3"
	"strconv"
	"bytes"
)

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

func counthash(dataHash []byte,prevHash []byte,ts uint64,height int) []byte{
	temp := dataHash
	temp = append(temp, prevHash...)
	temp = append(temp, []byte(strconv.Itoa(height))...)
	return SM3.SM3_256(temp)
}

func splitBytesbyn(a []byte) [][]byte{
	return bytes.SplitN(a, []byte("\n"), -1)
}

func splitStringbyn(a string) []string {
	return strings.SplitN(a, "\n", -1)
}


