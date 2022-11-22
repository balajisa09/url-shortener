package shortener

import (
	"testing"
)

var GetMD5HashInputs map[string][16]byte
var Getbase62Inputs map[[16]byte]string
var ShortInputs map[string]string

func init(){
	GetMD5HashInputs = map[string][16]byte{
		"https://bruce.com" : [16]uint8{244,116,66,136,129,205,70,209,55,28,61,61,199,6,253,109},
	}

	Getbase62Inputs = map[[16]byte]string{
		[16]uint8{244,116,66,136,129,205,70,209,55,28,61,61,199,6,253,109} : "3w1s14",
	}
	
	ShortInputs = map[string]string{
		"https://bruce.com": "3w1s142C253J183NtSzz3D6451l",
	}
}

func TestGetMD5Hash(t *testing.T){
	for url,hash := range GetMD5HashInputs{
		got := GetMD5Hash(url)
		expected := hash
		if got != expected{
			t.Errorf("got %q, expected %q", got, expected)
		}
	}
}

func TestGetbase62(t *testing.T){
	for hash,str := range Getbase62Inputs{
		got := Getbase62(hash)
		expected := str
		if got[:6] != expected{
			t.Errorf("got %q, expected %q", got, expected)
		}
	}
}

func TestShort(t *testing.T){
	for url,hash := range ShortInputs{
		got := Short(url)
		expected := hash
		if got != expected{
			t.Errorf("got %q, expected %q", got, expected)
		}
	}
}