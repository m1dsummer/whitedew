package file

import (
	"encoding/base64"
	"io/ioutil"
	"log"
)

func FileToBase64(file string) string {
	buffer,err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(buffer)
}
