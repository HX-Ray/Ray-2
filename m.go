package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strings"

	"github.com/howeyc/gopass"
)

func main() {
	requestUrl := "http://pass.muxi-tech.xyz/auth/api/signin"

	var name string
	fmt.Print("木犀用户名：")
	_, _ = fmt.Scanln(&name)
	fmt.Print("输入密码：")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}

	/*data := url.Values{}
	data.Set("username", name)
	data.Set("password", string(password))*/

	data := "{\"password\":\""+string(password)+"\",\"username\":\""+name+"\"}"

	payload := strings.NewReader(/*data.Encode()*/data)

	request, err := http.NewRequest("POST", requestUrl, payload)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Accept-Encoding" ,"gzip, deflate")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Content-Length", "50")
	request.Header.Add("content-type", "application/json")
	request.Header.Add("Host", "pass.muxi-tech.xyz")
	request.Header.Add("origin", "http://pass.muxi-tech.xyz")
	request.Header.Add("Referer", "http://pass.muxi-tech.xyz/?landing=work.muxi-tech.xyz/landing")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)

	fmt.Println(string(body))
}
