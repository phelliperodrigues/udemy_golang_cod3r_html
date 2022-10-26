package udemy_golang_cod3r_html

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulo obtem o titulo de uma pagina HTML
func Titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			html, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			r, err := regexp.Compile("<title>(.*?)</title>")
			if err != nil {
				fmt.Println(err)
			}
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
