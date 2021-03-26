package utils

import (
	"os"
	"time"

	"github.com/speps/go-hashids"
)



func KeyGen(val int)(string ,error) {
	hd := hashids.NewData()
	hd.Salt = os.Getenv("SECRET")
	hd.MinLength = 30
	hasher, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	key, err := hasher.Encode([]int{val,int(time.Now().Unix())})
	if err != nil {
		return "", err
	}
	return key,nil

}