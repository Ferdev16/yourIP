package ip

import (
	"io/ioutil"
	"net/http"
)

func GetIP() (string, error) {
	res, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}
