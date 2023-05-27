package ip

import (
	"io/ioutil"
	"net/http"
)

func GetIP() string {
	res, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	ip, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return ""
	}

	return string(ip)
}
