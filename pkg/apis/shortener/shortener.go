package shortener

import (
	"crypto/md5"
	"github.com/catinello/base62"
)

var UrlHashMap map[string]string

func Short(link string) string{
	//check if the url already exist 
	if _, ok := UrlHashMap[link]; ok {
		return UrlHashMap[link]
	}
	//generate hash
	hash := GetMD5Hash(link)
	base62 := Getbase62(hash)
	//store url -> hash in map
	UrlHashMap[link] = base62
	return base62
}


func GetMD5Hash(link string)[16]byte{
	data := []byte(link)
	return md5.Sum(data)
}

func Getbase62(hash [16]byte) string{
	result := ""
	for _,value := range hash{
		char := base62.Encode(int(value))
		result = result + char
	}
	return result
}

//initialize urlhash map when the server starts
func init(){
	UrlHashMap = make(map[string]string)
}