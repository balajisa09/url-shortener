package shortener

import (
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/catinello/base62"
)

var UrlHashMap map[string]string

type ShortenValue struct{
	URL string 
	Timestamp time.Time
}

func Short(link string) (*string,error){

	//check if the url already exist 
	readData, err := ioutil.ReadFile("db/values.json")
	if(err != nil){
		return nil,err
	}
	var readResp map[string]ShortenValue
	json.Unmarshal(readData,&readResp)
	
	for key,value := range readResp{
		if(value.URL == link){
			return &key,nil
		}
	}
	//generate hash
	hash := GetMD5Hash(link)
	base62 := Getbase62(hash)
	//store hash -> url in map in db
	if(readResp == nil){
		readResp = make(map[string]ShortenValue)
	}
	shortValue := ShortenValue{
		URL: link,
		Timestamp: time.Now(),
	}
	readResp[base62[:6]] = shortValue
	data,err := json.Marshal(readResp)
	if(err!= nil){
		return nil,err
	}
	err = ioutil.WriteFile("db/values.json",data,0644)
	if(err!= nil){
		return nil,err
	}
	return &base62,nil
}


func GetMD5Hash(link string)[16]byte{
	data := []byte(link)
	return md5.Sum(data)
}

func Getbase62(hash [16]byte) string{
	result := ""
	//itertae each byte and encode to base62
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