package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
		requestUrl := "https://manhua.dmzj.com/bianfuxiav3"
		response, err := http.Get(requestUrl)
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		content := string(body)
		fmt.Println(content)

}
