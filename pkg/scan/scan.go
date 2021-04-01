package scan

import (
	"log"
	"net/http"
	"strings"
)

func RespStatus(target string)int{
	resp,err := http.Get(target)
	if err!= nil {
		log.Fatal(err)
	}
	return resp.StatusCode
}


func ScanBacnkup(target string) bool{
	if strings.Contains(target, "http") {
		target = target
	}else {
		target = "http://" + target
		target = target
	}

	CheckCode := RespStatus(target)

	if CheckCode == 200 || CheckCode == 301 || CheckCode == 302 || CheckCode == 304 || CheckCode == 307 || CheckCode == 403{
		return true
	}
	return false
}
