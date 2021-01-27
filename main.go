package main

import (
    "io/ioutil"
    "log"
    "net/http"
	"net/http/cookiejar"
	"fmt"
	"regexp"
	"strings"
	"encoding/json"
	"bytes"

)

type MyTransport struct {
    Transport http.RoundTripper
}

func (t *MyTransport) transport() http.RoundTripper {
    if nil != t.Transport {
        return t.Transport
    }
    return http.DefaultTransport
}

func (t *MyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)")
    return t.transport().RoundTrip(req)
}

func NewClient() *http.Client {
    t := &MyTransport{}
    jar, err := cookiejar.New(nil)
    if nil != err {
        log.Fatal(err)
    }
    client := http.DefaultClient
    client.Transport = t
    client.Jar = jar
    return client
}

type LoginInfo struct{
		Login string `json:"login"`
		Password string `json:"password"`
		Remember int `json:"remember"`
		A_xfRedirect string `json:"_xfRedirect"`
		A_xfToken string `json:"_xfToken"`
	}

func main() {
   
	c := NewClient()

    // sUrl 是登录验证页面地址
    sUrl := "https://www.minebbs.com/login/login"
    req, err := http.NewRequest("GET", sUrl, nil)
	// 执行登录操作
	ress, errr := http.Get("http://www.minebbs.com/")
	if errr != nil {
		log.Fatal(errr)
	}
	ioutil.ReadAll(ress.Body)
    res, err := c.Do(req)
    if nil != err {
        log.Fatal(err)
	}
	
	out, err := ioutil.ReadAll(res.Body)
	
	str := string(out)
	forma := "\"_xfToken\" value=\"([a-zA-Z0-9,]+)\""
	r,_ := regexp.Compile(forma)
	
	strr := r.FindStringSubmatch(str)
	xftk := strr[1]
	
	
	logInfo := LoginInfo {
		Login: "sakuranaranbom%E6%9A%AE%E9%9B%AA",
		Password: "",
		Remember: 1,
		A_xfRedirect: "https%3A%2F%2Fwww.minebbs.com%2",
		A_xfToken: xftk,
	}
	info, err := json.Marshal(logInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bytes.NewReader(info))
	//fmt.Println(bytes.NewBuffer([]byte(info)))
	//fmt.Println(string(info))
	//req, err = http.NewRequest("POST", sUrl, strings.NewReader(canshu))
	//req, err = http.NewRequest("POST", sUrl, bytes.NewReader(info))
	req, err = http.NewRequest("GET", sUrl, nil)
	req.Header.Set("Pragma","no-cache")
    req.Header.Set("Upgrade-Insecure-Requests","1")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
    req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header.Set("Connection","keep-alive")
    req.Header.Set("Content-Type","application/json;charset=UTF-8")
	req.Header.Set("Cookie","yourcookie")
	respp, err := http.Post("https://www.minebbs.com/login/login","application/json",bytes.NewBuffer([]byte(info)))
	//resp, err := http.DefaultClient.Do(req)
	client := http.Client{}
	resp, err := client.Do(req)
        if err != nil {
                log.Fatal(err)
		}
	
	newout, err := ioutil.ReadAll(resp.Body)
	nnewout, err := ioutil.ReadAll(respp.Body)
	//fmt.Println(string(newout))
	
	fmt.Println(strings.NewReader(string(info)))
	ioutil.WriteFile("site.txt", newout, 0644)
	ioutil.WriteFile("nsite.html",nnewout, 0644)
}