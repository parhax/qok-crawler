package logic

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Crawl(sUrl string) (string, error) {
	resp, err := http.Get(sUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	strBody := string(body)
	re := regexp.MustCompile("<title*>(.*?)</title>")
	match := re.FindStringSubmatch(strBody)
	if len(match) <= 0 {
		str := fmt.Sprintf("Could not find any title for %s ", sUrl)
		err := errors.New(str)
		return "", err
	} else {
		return match[0], nil
	}
}
