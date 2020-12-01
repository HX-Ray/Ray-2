package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var keyword string
	fmt.Scanln(&keyword)

	//num := 3

	for i := 1; i <= 5; i++ {
		requestUrl := "https://www.dmzj.com/dynamic/o_search/index/"+ keyword + "/" + strconv.Itoa(i)

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

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			panic(err)
		}

		dom.Find(".update_con.autoHeight").Each(func(i int, selection *goquery.Selection) {
			// fmt.Println(selection.Text())
			selection.Find("p").Each(func(i int, title *goquery.Selection) {
				// fmt.Println(title.Text())
				// fmt.Printf("%3d   ", num/3)
				fmt.Println(title.Text())
				// num++
			})
		})
	}
}
